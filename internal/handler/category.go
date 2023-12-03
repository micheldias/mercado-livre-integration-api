package handler

import (
	"errors"
	"github.com/gorilla/mux"
	"mercado-livre-integration/internal/infrastructure/server"
	"mercado-livre-integration/internal/service"
	"net/http"
)

type CategoryHandler interface {
	GetCategories(request *http.Request) (server.HttpResponse, error)
}

func NewCategory(categoryService service.CategoryService) CategoryHandler {
	return categoryHandler{
		CategoryService: categoryService,
	}
}

type categoryHandler struct {
	CategoryService service.CategoryService
}

func (c categoryHandler) GetCategories(r *http.Request) (server.HttpResponse, error) {
	ctx := r.Context()
	vars := mux.Vars(r)
	siteID := vars["siteID"]

	if appID := r.URL.Query().Get("applicationID"); appID == "" {
		return server.HttpResponse{}, errors.New("applicationID is required")
	}

	categories, err := c.CategoryService.GetCategories(ctx, siteID)
	if err != nil {
		return server.HttpResponse{}, err
	}

	return server.HttpResponse{
		StatusCode: 200,
		Body:       categories,
	}, nil

}
