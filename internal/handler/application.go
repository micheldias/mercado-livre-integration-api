package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mercado-livre-integration/internal/infrastructure/server"
	"mercado-livre-integration/internal/model"
	"mercado-livre-integration/internal/service"
	"net/http"
	"strconv"
)

type applicationResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	RedirectURL string `json:"redirect_url"`
}

func fromDomain(domain model.Application) applicationResponse {
	return applicationResponse{
		ID:          domain.ID,
		Name:        domain.Name,
		ClientID:    domain.ClientID,
		Secret:      domain.Secret,
		RedirectURL: domain.RedirectURL,
	}
}

type ApplicationHandler interface {
	GetByID(request *http.Request) (server.HttpResponse, error)
	GetAll(request *http.Request) (server.HttpResponse, error)
	Save(request *http.Request) (server.HttpResponse, error)
}

func NewApplication(service service.ApplicationService) ApplicationHandler {
	return appHandler{
		ApplicationService: service,
	}
}

type appHandler struct {
	ApplicationService service.ApplicationService
}

func (a appHandler) GetByID(r *http.Request) (server.HttpResponse, error) {
	ctx := r.Context()
	vars := mux.Vars(r)
	appID, _ := strconv.Atoi(vars["id"])
	app, err := a.ApplicationService.GetAppByID(ctx, appID)
	if err != nil {
		return server.HttpResponse{}, err
	}
	return server.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       fromDomain(app),
	}, nil
}

func (a appHandler) GetAll(r *http.Request) (server.HttpResponse, error) {
	ctx := r.Context()
	app, err := a.ApplicationService.GetApps(ctx)
	if err != nil {
		return server.HttpResponse{}, err
	}
	return server.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       app,
	}, nil
}

func (a appHandler) Save(r *http.Request) (server.HttpResponse, error) {
	ctx := r.Context()
	var app model.Application
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		return server.HttpResponse{}, err
	}

	app, err := a.ApplicationService.SaveApp(ctx, app)
	if err != nil {
		return server.HttpResponse{}, err
	}
	return server.HttpResponse{
		StatusCode: http.StatusCreated,
		Body:       app,
	}, nil
}
