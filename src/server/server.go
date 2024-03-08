package server

import (
	"api-rest/src/database"
	"api-rest/src/repository"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Config struct {
	Port string
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
	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil {
		fmt.Println("Error connecting to database")
		log.Fatal(err)
	}
	repository.SetRepository(repo)
	err = http.ListenAndServe(b.config.Port, b.router)
	if err != nil {
		log.Fatal("Listen and Serve", err)
	}
}
