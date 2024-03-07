package handlers

import (
	"api-rest/src/models"
	"api-rest/src/repository"
	"api-rest/src/server"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
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

type GetBugResponse struct {
	Id uint32 `json:"id"`
	Description string `json:"description"`
	Username string `json:"username"`
	CreationDate time.Time `json:"creationDate"`
	Project *models.Project `json:"project"`
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

func GetBugByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseUint(params["id"], 10, 32)
		if err != nil {
			fmt.Println("Error getting id from route param")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bug, err := repository.GetBugById(r.Context(), uint32(id))
		if err != nil {
			fmt.Println("Error getting bug")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if bug.Id == 0 {
			fmt.Println("Bug not found")
			http.Error(w, "Bug not found", http.StatusNotFound)
			return
		}
		user, err := repository.GetUserById(r.Context(), bug.UserId)
		if user.Id == 0 {
			fmt.Println("User not found")
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		project, err := repository.GetProjectById(r.Context(), bug.ProjectId)
		if project.Id == 0 {
			fmt.Println("Project not found")
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GetBugResponse{
			Id: bug.Id,
			Description: bug.Description,
			CreationDate: bug.CreationDate,
			Username: user.Name + user.Surname,
			Project: project,
		})
	}
}
