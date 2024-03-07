package handlers

import (
	"api-rest/src/models"
	"api-rest/src/repository"
	"api-rest/src/server"
	"encoding/json"
	"fmt"
	"net/http"
)

type SignUpRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id uint32 `json:"id"`
	Email string `json:"email"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := SignUpRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			fmt.Println("Error decoding")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user := models.User{
			Email: request.Email,
			Password: request.Password,
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			fmt.Println("Error inserting user")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpResponse{
			Id: user.Id,
			Email: user.Email,
		})
	}
}
