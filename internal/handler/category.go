package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mercado-livre-integration/internal/service"
	"net/http"
)

type CategoryHandler interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
}

func NewCategory(categoryService service.CategoryService) CategoryHandler {
	return categoryHandler{
		CategoryService: categoryService,
	}
}

type categoryHandler struct {
	CategoryService service.CategoryService
}

func (c categoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	siteID := vars["siteID"]
	categories, err := c.CategoryService.GetCategories(siteID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	response, _ := json.Marshal(categories)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "handler/json")
	w.Write(response)
}
