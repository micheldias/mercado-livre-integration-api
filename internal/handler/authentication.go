package handler

import (
	"encoding/json"
	"mercado-livre-integration/internal/infrastructure/server"
	"mercado-livre-integration/internal/service"
	"net/http"
)

type tokenRequest struct {
	AuthCode string `json:"authCode"`
}
type UrlResponse struct {
	Url string `json:"url"`
}
type Authentication interface {
	Create(request *http.Request) (server.HttpResponse, error)
	GetUrlAuthentication(request *http.Request) (server.HttpResponse, error)
}

func NewToken(tokenService service.AuthenticationService) Authentication {
	return authHandler{
		AuthenticationService: tokenService,
	}
}

type authHandler struct {
	AuthenticationService service.AuthenticationService
}

func (t authHandler) Create(r *http.Request) (server.HttpResponse, error) {
	ctx := r.Context()

	var payload tokenRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return server.HttpResponse{}, err
	}

	token, err := t.AuthenticationService.Create(ctx, payload.AuthCode)
	if err != nil {
		return server.HttpResponse{}, err
	}
	return server.HttpResponse{
		StatusCode: http.StatusCreated,
		Body:       token,
	}, nil
}

func (t authHandler) GetUrlAuthentication(request *http.Request) (server.HttpResponse, error) {
	ctx := request.Context()
	url := t.AuthenticationService.GetUrlAuthentication(ctx)

	return server.HttpResponse{
		StatusCode: 200,
		Body:       UrlResponse{Url: url},
	}, nil
}
