package handler

import (
	"encoding/json"
	contexthelper "mercado-livre-integration/internal/infrastructure/contextHelper"
	logs "mercado-livre-integration/internal/infrastructure/log"
	"mercado-livre-integration/internal/service"
	"net/http"
)

type tokenRequest struct {
	AuthCode string `json:"authCode"`
}
type urlResponse struct {
	Url string `json:"url"`
}
type Authentication interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetUrlAuthentication(w http.ResponseWriter, r *http.Request)
}

func NewToken(tokenService service.AuthenticationService) Authentication {
	return authHandler{
		AuthenticationService: tokenService,
	}
}

type authHandler struct {
	AuthenticationService service.AuthenticationService
}

func (t authHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx := r.Context()
	logger := logs.New("mercado-livre-api")
	ctx = contexthelper.SetLogger(ctx, logger)

	var payload tokenRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := t.AuthenticationService.Create(ctx, payload.AuthCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	response, _ := json.Marshal(token)
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (t authHandler) GetUrlAuthentication(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	logger := logs.New("mercado-livre-api")
	ctx := contexthelper.SetLogger(r.Context(), logger)
	url := t.AuthenticationService.GetUrlAuthentication(ctx)

	response, _ := json.Marshal(urlResponse{Url: url})
	w.Write(response)
}
