package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Port string
	JWTSecret string 
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config{
	return b.config
}

func NewServer (ctx context.Context, config *Config) (*Broker, error) {
	if config.DatabaseUrl == "" {
		return nil, errors.New("DatabaseUrl required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("JWTSecret required")
	}
	if config.Port == "" {
		return nil, errors.New("Port required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func (s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	err := http.ListenAndServe(b.config.Port, b.router)
	if err != nil {
		log.Fatal("Listen and Serve", err)
	}
}
