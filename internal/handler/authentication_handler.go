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
type Token interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func NewToken(tokenService service.TokenService) Token {
	return tokenHandler{
		TokenService: tokenService,
	}
}

type tokenHandler struct {
	TokenService service.TokenService
}

func (t tokenHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx := r.Context()
	logger := logs.New("mercado-livre-api")
	ctx = contexthelper.SetLogger(ctx, logger)

	var payload tokenRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := t.TokenService.Create(ctx, payload.AuthCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	response, _ := json.Marshal(token)
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
