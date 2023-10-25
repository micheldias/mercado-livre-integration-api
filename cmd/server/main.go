package main

import (
	"github.com/spf13/viper"
	"io"
	handler "mercado-livre-integration/internal/handler"
	"mercado-livre-integration/internal/infrastructure/client"
	"mercado-livre-integration/internal/infrastructure/server"
	"mercado-livre-integration/internal/service"
	"net/http"
)

func init() {
	server.LoadEnvVars()
}

func main() {
	client := client.NewMercadoLivre(
		viper.GetString("MERCADO_LIVRE_CLIENT_ID"),
		viper.GetString("MERCADO_LIVRE_SECRET"),
		viper.GetString("MERCADO_LIVRE_REDIRECT_URL"),
		viper.GetString("MERCADO_LIVRE_API_URL"),
		viper.GetDuration("MERCADO_LIVRE_EXECUTE_TIMES"),
	)

	categoryHandler := handler.NewCategory(service.NewCategory(client))
	tokenHandler := handler.NewToken(service.NewAuthenticationService(client))

	server.Builder().
		Use(server.InjectRequestID).
		Use(server.InjectLogger).
		Use(server.HandlePanic).
		//TODO: move this to inside the builder
		AddRouter("/health", http.MethodGet, HealthCheckHandler).
		AddRouter("/api/v1/sites/{siteID}/categories", http.MethodGet, categoryHandler.GetCategories).
		AddRouter("/api/v1/tokens", http.MethodPost, tokenHandler.Create).
		AddRouter("/api/v1/auth/url", http.MethodGet, tokenHandler.GetUrlAuthentication).
		StartServer()

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}
