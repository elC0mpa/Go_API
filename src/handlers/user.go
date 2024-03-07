package handlers

import (
	"api-rest/src/models"
	"api-rest/src/repository"
	"api-rest/src/server"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserRequest struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
}

type CreateUserResponse struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
}

func CreateUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := CreateUserRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			fmt.Println("Error decoding", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user := models.User{
			Name: request.Name,
			Surname: request.Surname,
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			fmt.Println("Error inserting user")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(CreateUserResponse{
			Name: user.Name,
			Surname: user.Surname,
		})
	}
}
