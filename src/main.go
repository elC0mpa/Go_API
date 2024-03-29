package main

import (
	"api-rest/src/handlers"
	"api-rest/src/server"
	"context"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port: PORT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/user", handlers.CreateUserHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", handlers.GetUserByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/project", handlers.CreateProjectHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/bug", handlers.CreateBugHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/bug", handlers.ListBugsHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/bug/{id}", handlers.GetBugByIdHandler(s)).Methods(http.MethodGet)
}
