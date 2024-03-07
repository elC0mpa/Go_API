package main

import (
	"api-rest/src/handlers"
	"api-rest/src/server"
	"context"
	"fmt"
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
	DATABASE_URL := os.Getenv("DATABASE_URL")
	fmt.Println(DATABASE_URL)

	s, err := server.NewServer(context.Background(), &server.Config{
		Port: PORT,
		JWTSecret: JWT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/user", handlers.CreateUserHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/project", handlers.CreateProjectHandler(s)).Methods(http.MethodPost)
}
