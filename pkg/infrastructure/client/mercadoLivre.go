package client

import (
	"encoding/json"
	"fmt"
	util "mercado-livre-integration/pkg/infrastructure/http"
	"net/http"
	"net/url"
	"strings"
)

type MercadoLivre interface {
	CreateToken(authCode string) (*AuthTokenResponse, error)
	CreateRefreshToken(refreshToken string) (*AuthTokenResponse, error)
	GetUser(userID string) (User, error)
}

// NewMercadoLivre creates a new Mercado Livre client
func NewMercadoLivre(clientID, secret, redirectUrl, url string) MercadoLivre {
	client := &mercadoLivre{
		url:          url,
		clientID:     clientID,
		clientSecret: secret,
		redirectUrl:  redirectUrl,
		httpClient: &http.Client{
			Transport: &util.RoundTripperLogger{Inner: http.DefaultTransport},
		},
		Cache: make(map[string]string),
	}
	return client
}

type mercadoLivre struct {
	url          string
	clientID     string
	clientSecret string
	redirectUrl  string
	httpClient   *http.Client
	Cache        map[string]string
}

func (m mercadoLivre) GetUser(userID string) (User, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/users/%s", m.url, userID), nil)
	if err != nil {
		return User{}, fmt.Errorf("failed to create http request: %s ", err.Error())
	}
	request.Header.Add("content-type", "application/json")

	response, err := m.httpClient.Do(request)

	if err != nil {
		return User{}, fmt.Errorf("failed to execute http request: %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errorResponse := &Error{}
		if err = json.NewDecoder(response.Body).Decode(&errorResponse); err != nil {
			return User{}, fmt.Errorf("failed to parse body: %s", err.Error())
		}
		return User{}, fmt.Errorf("status code: %d", response.StatusCode)
	}
	user := User{}
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		return User{}, fmt.Errorf("failed to parse response: %s", err.Error())
	}
	return user, err
}

func (m mercadoLivre) CreateToken(authCode string) (*AuthTokenResponse, error) {
	body := m.toTokenBody(authCode)
	tokenResponse, err := m.requestToken(body)
	return &tokenResponse, err
}

func (m mercadoLivre) CreateRefreshToken(refreshToken string) (*AuthTokenResponse, error) {
	body := m.toRefreshTokenBody(refreshToken)
	tokenResponse, err := m.requestToken(body)
	return &tokenResponse, err
}

func (m mercadoLivre) requestToken(body *strings.Reader) (AuthTokenResponse, error) {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/oauth/token", m.url), body)
	if err != nil {
		return AuthTokenResponse{}, fmt.Errorf("failed to create http request: %s ", err.Error())
	}
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	response, err := m.httpClient.Do(request)
	if err != nil {
		return AuthTokenResponse{}, fmt.Errorf("failed to execute http request: %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errorResponse := &Error{}
		if err = json.NewDecoder(response.Body).Decode(&errorResponse); err != nil {
			return AuthTokenResponse{}, fmt.Errorf("failed to parse body: %s", err.Error())
		}
		return AuthTokenResponse{}, fmt.Errorf("status code: %d", response.StatusCode)
	}
	token := AuthTokenResponse{}
	if err = json.NewDecoder(response.Body).Decode(&token); err != nil {
		return AuthTokenResponse{}, fmt.Errorf("failed to parse response: %s", err.Error())
	}
	return token, err
}

func (m mercadoLivre) toTokenBody(authCode string) *strings.Reader {
	data := m.buildTokenBodyFields()
	data.Set("grant_type", "authorization_code")
	data.Set("code", authCode)
	data.Set("redirect_uri", m.redirectUrl)
	return strings.NewReader(data.Encode())
}
func (m mercadoLivre) toRefreshTokenBody(refreshToken string) *strings.Reader {
	data := m.buildTokenBodyFields()
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	return strings.NewReader(data.Encode())
}

func (m mercadoLivre) buildTokenBodyFields() url.Values {
	data := url.Values{}
	data.Set("client_id", m.clientID)
	data.Set("client_secret", m.clientSecret)
	return data
}

var _ MercadoLivre = (*mercadoLivre)(nil)
