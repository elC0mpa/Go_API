package handlers

import (
	"api-rest/src/models"
	"api-rest/src/server"
	"api-rest/src/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateProjectRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type CreateProjectResponse struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func CreateProjectHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := CreateProjectRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			fmt.Println("Error decoding", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		project := models.Project{
			Name: request.Name,
			Description: request.Description,
		}
		err = services.CreateProject(r.Context(), &project)
		if err != nil {
			fmt.Println("Error inserting project")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(CreateProjectResponse{
			Name: project.Name,
			Description: project.Description,
		})
	}
}
