package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type MercadoLivre interface {
	CreateToken(authCode string) (*AuthTokenResponse, error)
}

// NewMercadoLivre creates a new Mercado Livre client
func NewMercadoLivre(clientID, secret, redirectUrl, url string, httpClient *http.Client) MercadoLivre {
	return &mercadoLivre{
		url:          url,
		clientID:     clientID,
		clientSecret: secret,
		redirectUrl:  redirectUrl,
		httpClient:   httpClient,
	}
}

type mercadoLivre struct {
	url          string
	clientID     string
	clientSecret string
	redirectUrl  string
	httpClient   *http.Client
}

func (m mercadoLivre) CreateToken(authCode string) (*AuthTokenResponse, error) {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/oauth/token", m.url), m.createTokenBody(authCode))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %s ", err.Error())
	}
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	response, err := m.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to execute http request: %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errorResponse := &AuthTokenErrorResponse{}
		if err = json.NewDecoder(response.Body).Decode(&errorResponse); err != nil {
			return nil, fmt.Errorf("failed to parse body: %s", err.Error())
		}
		return nil, fmt.Errorf("status code: %d", response.StatusCode)
	}
	token := AuthTokenResponse{}
	if err = json.NewDecoder(response.Body).Decode(&token); err != nil {
		return nil, fmt.Errorf("failed to parse response: %s", err.Error())
	}
	return &token, nil
}

func (m mercadoLivre) createTokenBody(authCode string) *strings.Reader {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", m.clientID)
	data.Set("client_secret", m.clientSecret)
	data.Set("code", authCode)
	data.Set("redirect_uri", m.redirectUrl)
	return strings.NewReader(data.Encode())
}
