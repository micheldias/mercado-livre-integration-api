package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/newrelic/go-agent/v3/newrelic"
	util "mercado-livre-integration/internal/infrastructure/http"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type MercadoLivre interface {
	CreateToken(ctx context.Context, clientID, clientSecret, redirectURL, authCode string) (AuthTokenResponse, error)
	GetUser(ctx context.Context, userID string) (User, error)
	GetSites(ctx context.Context) (Sites, error)
	GetCategories(ctx context.Context, appId int, siteID string) (Categories, error)
	CreateProduct(ctx context.Context, product ProductRequest) (ProductResponse, error)
}

// NewMercadoLivre creates a new Mercado Livre client
func NewMercadoLivre(url string, executeTimes time.Duration) MercadoLivre {
	client := &mercadoLivre{
		url:          url,
		executeTimes: executeTimes,
		httpClient: &http.Client{
			Transport: newrelic.NewRoundTripper(&util.RoundTripperLogger{Inner: http.DefaultTransport}),
		},
		cache: make(map[string]CacheAuth),
	}

	go client.refreshTokenTask()
	return client
}

type mercadoLivre struct {
	url          string
	httpClient   *http.Client
	cache        map[string]CacheAuth
	executeTimes time.Duration
}

func (m mercadoLivre) CreateProduct(ctx context.Context, product ProductRequest) (ProductResponse, error) {

	body, err := json.Marshal(product)
	if err != nil {
		return ProductResponse{}, fmt.Errorf("serializing request body: %w", err)
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/items", m.url), bytes.NewReader(body))

	var createdProduct ProductResponse
	if err != nil {
		return createdProduct, fmt.Errorf("failed to create http request: %s ", err.Error())
	}
	request.Header.Add("content-type", "handler/json")
	return makeRequestAndConvertResponseBody[ProductResponse](m, request)
}

func (m mercadoLivre) GetCategories(ctx context.Context, appId int, siteID string) (Categories, error) {
	//TODO: get cache by appID.

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/sites/%s/categories", m.url, siteID), nil)
	var categories Categories
	if err != nil {
		return categories, fmt.Errorf("failed to create http request: %s ", err.Error())
	}
	request.Header.Add("content-type", "handler/json")
	return makeRequestAndConvertResponseBody[Categories](m, request)
}

func (m mercadoLivre) GetSites(ctx context.Context) (Sites, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/sites", m.url), nil)
	var sites Sites
	if err != nil {
		return sites, fmt.Errorf("failed to create http request: %s ", err.Error())
	}
	request.Header.Add("content-type", "handler/json")
	user, err := makeRequestAndConvertResponseBody[Sites](m, request)
	return user, err
}

func (m mercadoLivre) GetUser(ctx context.Context, userID string) (User, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/users/%s", m.url, userID), nil)
	if err != nil {
		return User{}, fmt.Errorf("failed to create http request: %s ", err.Error())
	}
	request.Header.Add("content-type", "handler/json")
	user, err := makeRequestAndConvertResponseBody[User](m, request)
	return user, err
}

func (m mercadoLivre) CreateToken(ctx context.Context, clientID, clientSecret, redirectURL, authCode string) (AuthTokenResponse, error) {
	body := m.toTokenBody(clientID, clientSecret, redirectURL, authCode)
	tokenResponse, err := m.requestToken(ctx, body)
	m.cache["refresh_token"] = CacheAuth{
		ClientID:          clientID,
		ClientSecret:      clientSecret,
		AuthTokenResponse: tokenResponse,
	}

	return tokenResponse, err
}

func (m mercadoLivre) refreshToken(clientID, clientSecret, refreshToken string) (AuthTokenResponse, error) {
	body := m.toRefreshTokenBody(clientID, clientSecret, refreshToken)
	tokenResponse, err := m.requestToken(context.Background(), body)
	return tokenResponse, err
}

func (m mercadoLivre) requestToken(ctx context.Context, body *strings.Reader) (AuthTokenResponse, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/oauth/token", m.url), body)
	if err != nil {
		return AuthTokenResponse{}, fmt.Errorf("failed to create http request: %s ", err.Error())
	}
	request.Header.Add("content-type", "handler/x-www-form-urlencoded")
	token, err := makeRequestAndConvertResponseBody[AuthTokenResponse](m, request)
	return token, err
}

func (m mercadoLivre) toTokenBody(clientID, clientSecret, redirectURL, authCode string) *strings.Reader {
	data := m.buildTokenBodyFields(clientID, clientSecret)
	data.Set("grant_type", "authorization_code")
	data.Set("code", authCode)
	data.Set("redirect_uri", redirectURL)
	return strings.NewReader(data.Encode())
}
func (m mercadoLivre) toRefreshTokenBody(clientID, clientSecret, refreshToken string) *strings.Reader {
	data := m.buildTokenBodyFields(clientID, clientSecret)
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	return strings.NewReader(data.Encode())
}

func (m mercadoLivre) buildTokenBodyFields(clientID, clientSecret string) url.Values {
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	return data
}

func makeRequestAndConvertResponseBody[T any](m mercadoLivre, request *http.Request) (T, error) {
	var base T

	response, err := m.httpClient.Do(request)
	if err != nil {
		return base, fmt.Errorf("failed to execute http request. Error: %s", err.Error())
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		errorResponse := &Error{}
		if err := json.NewDecoder(response.Body).Decode(&errorResponse); err != nil {
			return base, fmt.Errorf("failed to parse body: %s", err.Error())
		}
		return base, fmt.Errorf("status code: %d", response.StatusCode)
	}

	if err := json.NewDecoder(response.Body).Decode(&base); err != nil {
		return base, fmt.Errorf("failed to parse response: %s", err.Error())
	}
	return base, nil
}

func (m mercadoLivre) refreshTokenTask() {
	ticker := time.NewTicker(m.executeTimes)
	go func() {
		for range ticker.C {
			//TODO: make refresh token by applicationID
			refreshToken, ok := m.cache["refresh_token"]
			if ok {
				token, err := m.refreshToken(refreshToken.ClientID, refreshToken.ClientSecret, refreshToken.AuthTokenResponse.RefreshToken)
				if err != nil {
					fmt.Println("error ao fazer refresh de token")
				}
				m.cache["refresh_token"] = CacheAuth{
					ClientID:          refreshToken.ClientID,
					ClientSecret:      refreshToken.ClientSecret,
					AuthTokenResponse: token,
				}
			} else {
				fmt.Println("refresh token nao localizado")
			}

		}
	}()
	//ticker.Stop()
}

var _ MercadoLivre = (*mercadoLivre)(nil)
