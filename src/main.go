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
	JWT := os.Getenv("JWT_SECRET")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port: PORT,
		JWTSecret: JWT,
		DatabaseUrl: "test",
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
}
