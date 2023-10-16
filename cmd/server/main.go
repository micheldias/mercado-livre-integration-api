package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"io"
	"log"
	handler "mercado-livre-integration/internal/handler"
	"mercado-livre-integration/internal/infrastructure/client"
	"mercado-livre-integration/internal/service"
	"net/http"
)

func init() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	client := client.NewMercadoLivre(
		viper.GetString("MERCADO_LIVRE_CLIENT_ID"),
		viper.GetString("MERCADO_LIVRE_SECRET"),
		viper.GetString("MERCADO_LIVRE_REDIRECT_URL"),
		viper.GetString("MERCADO_LIVRE_API_URL"),
		viper.GetDuration("MERCADO_LIVRE_EXECUTE_TIMES"),
	)
	categoryService := service.NewCategory(client)
	categoryHandler := handler.NewCategory(categoryService)

	tokenService := service.NewAuthenticationService(client)
	tokenHandler := handler.NewToken(tokenService)
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler).Methods(http.MethodGet)
	s := r.PathPrefix("/api/v1/").Subrouter()
	s.HandleFunc("sites/{siteID}/categories", categoryHandler.GetCategories).Methods(http.MethodGet)
	s.HandleFunc("/tokens", tokenHandler.Create).Methods(http.MethodPost)
	s.HandleFunc("/auth/url", tokenHandler.GetUrlAuthentication).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", viper.GetString("SERVER_PORT")),
		WriteTimeout: viper.GetDuration("SERVER_WRITE_TIMEOUT"),
		ReadTimeout:  viper.GetDuration("SERVER_READ_TIMEOUT"),
	}

	log.Fatal(srv.ListenAndServe())
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}
