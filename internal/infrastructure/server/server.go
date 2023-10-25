package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"io"
	logs "mercado-livre-integration/internal/infrastructure/log"
	"net/http"
	"os"
)

func LoadEnvVars() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

type WebServerBuilder interface {
	AddRouter(path, method string, handler func(w http.ResponseWriter, r *http.Request)) WebServerBuilder
	Use(func(http.Handler) http.Handler) WebServerBuilder
	StartServer()
}

func NewWebServerBuilder() WebServerBuilder {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthCheckHandler).Methods(http.MethodGet)
	return &server{
		Router: r,
	}
}

type server struct {
	Router *mux.Router
}

func (s server) AddRouter(path, method string, handler func(w http.ResponseWriter, r *http.Request)) WebServerBuilder {
	s.Router.HandleFunc(path, handler).Methods(method)
	s.Router.Use()
	return s
}

func (s server) Use(middleware func(http.Handler) http.Handler) WebServerBuilder {
	s.Router.Use(middleware)
	return s
}
func (s server) StartServer() {
	port := viper.GetString("SERVER_PORT")
	srv := &http.Server{
		Handler:      s.Router,
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: viper.GetDuration("SERVER_WRITE_TIMEOUT"),
		ReadTimeout:  viper.GetDuration("SERVER_READ_TIMEOUT"),
	}
	logger := logs.New("mercado-livre-api")
	logger.Info(fmt.Sprintf("server started on port :%s", port))
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("error on start server", "error", err)
		os.Exit(1)
	}

}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}
