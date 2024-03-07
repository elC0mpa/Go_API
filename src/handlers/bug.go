package handlers

import (
	"api-rest/src/models"
	"api-rest/src/repository"
	"api-rest/src/server"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateBugRequest struct {
	User uint32 `json:"user"`
	Project uint32 `json:"project"`
	Description string `json:"description"`
}

type CreateBugResponse struct {
	User uint32 `json:"user"`
	Project uint32 `json:"project"`
	Description string `json:"description"`
}

func CreateBugHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := CreateBugRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			fmt.Println("Error decoding", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		bug := models.Bug{
			Description: request.Description,
			UserId: request.User,
			ProjectId: request.Project,
		}
		err = repository.InsertBug(r.Context(), &bug)
		if err != nil {
			fmt.Println("Error inserting bug")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(CreateBugResponse{
			User: bug.UserId,
			Project: bug.ProjectId,
			Description: bug.Description,
		})
	}
}
