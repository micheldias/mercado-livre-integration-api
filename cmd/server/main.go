package main

import (
	"github.com/gorilla/mux"
	"log"
	handler "mercado-livre-integration/internal/handler"
	"mercado-livre-integration/internal/infrastructure/client"
	"mercado-livre-integration/internal/service"
	"net/http"
	"time"
)

//TODO: colocar aqui um init

func main() {
	client := client.NewMercadoLivre("", "", "", "", time.Hour*2)
	service := service.NewCategory(client)
	categoryHandler := handler.NewCategory(service)
	r := mux.NewRouter()

	r.HandleFunc("sites/{siteID}/categories", categoryHandler.GetCategories).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", r))
}
