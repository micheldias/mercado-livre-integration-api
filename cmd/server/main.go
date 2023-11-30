package main

import (
	"github.com/spf13/viper"
	"mercado-livre-integration/internal/handler"
	"mercado-livre-integration/internal/infrastructure/client"
	"mercado-livre-integration/internal/infrastructure/database"
	"mercado-livre-integration/internal/infrastructure/server"
	"mercado-livre-integration/internal/repository"
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
	db, _ := database.NewDatabase(database.DBConfig{
		Host:     viper.GetString("DATABASE_HOST"),
		Port:     viper.GetInt("DATABASE_PORT"),
		DbName:   viper.GetString("DATABASE_NAME"),
		User:     viper.GetString("DATABASE_USER"),
		Password: viper.GetString("DATABASE_PASSWORD"),
	})
	service.NewApplicationService(repository.NewApplicationRepository(db))

	server.NewWebServerBuilder().
		Use(server.InjectRequestID).
		Use(server.InjectLogger).
		Use(server.Recovery).
		AddRouter("/api/v1/sites/{siteID}/categories", http.MethodGet, categoryHandler.GetCategories).
		AddRouter("/api/v1/tokens", http.MethodPost, tokenHandler.Create).
		AddRouter("/api/v1/auth/url", http.MethodGet, tokenHandler.GetUrlAuthentication).
		StartServer()

}
