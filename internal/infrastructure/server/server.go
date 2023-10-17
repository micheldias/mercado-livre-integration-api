package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func LoadEnvVars() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

type IBuilder interface {
	AddRouter(path, method string, handler func(w http.ResponseWriter, r *http.Request)) IBuilder
	Use(func(http.Handler) http.Handler) IBuilder
	StartServer()
}

func Builder() IBuilder {
	return &server{
		Router: mux.NewRouter(),
	}
}

type server struct {
	Router *mux.Router
}

func (s server) AddRouter(path, method string, handler func(w http.ResponseWriter, r *http.Request)) IBuilder {
	s.Router.HandleFunc(path, handler).Methods(method)
	s.Router.Use()
	return s
}

func (s server) Use(middleware func(http.Handler) http.Handler) IBuilder {
	s.Router.Use(middleware)
	return s
}
func (s server) StartServer() {
	srv := &http.Server{
		Handler:      s.Router,
		Addr:         fmt.Sprintf(":%s", viper.GetString("SERVER_PORT")),
		WriteTimeout: viper.GetDuration("SERVER_WRITE_TIMEOUT"),
		ReadTimeout:  viper.GetDuration("SERVER_READ_TIMEOUT"),
	}
	log.Fatal(srv.ListenAndServe())
}
