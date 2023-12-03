package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	util "mercado-livre-integration/internal/infrastructure/http"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type MercadoLivre interface {
	CreateToken(ctx context.Context, authCode string) (AuthTokenResponse, error)
	GetUser(ctx context.Context, userID string) (User, error)
	GetSites(ctx context.Context) (Sites, error)
	GetCategories(ctx context.Context, siteID string) (Categories, error)
	CreateProduct(ctx context.Context, product ProductRequest) (ProductResponse, error)
}

// NewMercadoLivre creates a new Mercado Livre client
func NewMercadoLivre(url string, executeTimes time.Duration) MercadoLivre {
	client := &mercadoLivre{
		url:          url,
		executeTimes: executeTimes,
		httpClient: &http.Client{
			Transport: &util.RoundTripperLogger{Inner: http.DefaultTransport},
		},
		cache: make(map[string]AuthTokenResponse),
	}

	go client.refreshTokenTask()
	return client
}

type mercadoLivre struct {
	url          string
	clientID     string
	clientSecret string
	redirectUrl  string
	httpClient   *http.Client
	cache        map[string]AuthTokenResponse
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

func (m mercadoLivre) GetCategories(ctx context.Context, siteID string) (Categories, error) {
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

func (m mercadoLivre) CreateToken(ctx context.Context, authCode string) (AuthTokenResponse, error) {
	body := m.toTokenBody(authCode)
	tokenResponse, err := m.requestToken(ctx, body)
	m.cache["refresh_token"] = tokenResponse
	return tokenResponse, err
}

func (m mercadoLivre) refreshToken(refreshToken string) (AuthTokenResponse, error) {
	body := m.toRefreshTokenBody(refreshToken)
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
			refreshToken, ok := m.cache["refresh_token"]
			if ok {
				token, err := m.refreshToken(refreshToken.RefreshToken)
				if err != nil {
					fmt.Println("error ao fazer refresh de token")
				}
				m.cache["refresh_token"] = token
			} else {
				fmt.Println("refresh token nao localizado")
			}

		}
	}()
	//ticker.Stop()
}

var _ MercadoLivre = (*mercadoLivre)(nil)
