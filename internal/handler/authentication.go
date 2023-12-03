package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mercado-livre-integration/internal/infrastructure/server"
	"mercado-livre-integration/internal/service"
	"net/http"
	"strconv"
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
	appID, _ := strconv.Atoi(mux.Vars(r)["id"])
	token, err := t.AuthenticationService.Create(ctx, appID, payload.AuthCode)
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
	appID, _ := strconv.Atoi(mux.Vars(request)["id"])
	url, err := t.AuthenticationService.GetUrlAuthentication(ctx, appID)
	if err != nil {
		return server.HttpResponse{}, err
	}
	return server.HttpResponse{
		StatusCode: 200,
		Body:       UrlResponse{Url: url},
	}, nil
}
