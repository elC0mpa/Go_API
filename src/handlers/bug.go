package handlers

import (
	"api-rest/src/models"
	"api-rest/src/server"
	"api-rest/src/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
		err = services.CreateBug(r.Context(), &bug)
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
		bug, err := services.GetBugById(r.Context(), uint32(id))
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
		project, user, err := services.PopulateBug(r.Context(), bug)
		if err != nil {
			fmt.Println("Error when populating bug", err)
			http.Error(w, "Problem populating bug", http.StatusInternalServerError)
			return
		} else if project == nil || user == nil {
			fmt.Println("Error when populating bug", err)
			http.Error(w, "Project or user not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GetBugResponse{
			Id: bug.Id,
			Description: bug.Description,
			CreationDate: bug.CreationDate,
			Username: strings.Join([]string{user.Name, user.Surname}, " "),
			Project: project,
		})
	}
}

func ListBugsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
    		projectID := r.URL.Query().Get("project_id")
    		startDateStr := r.URL.Query().Get("start_date")
    		endDateStr := r.URL.Query().Get("end_date")

		if (userID == "" && projectID == "" && startDateStr == "" && endDateStr == "") {
			fmt.Println("Error getting id from route param")
			http.Error(w, "At least one filter is required", http.StatusBadRequest)
			return
		}

		var userIDUint, projectIDUint uint64
    		var startDate, endDate *time.Time
    		var err error

    		if userID != "" {
      			userIDUint, err = strconv.ParseUint(userID, 10, 32)
      			if err != nil {
        			fmt.Println("Error parsing user_id")
        			http.Error(w, err.Error(), http.StatusBadRequest)
        			return
      			}
    		}
		
    		if projectID != "" {
      			projectIDUint, err = strconv.ParseUint(projectID, 10, 32)
      			if err != nil {
        			fmt.Println("Error parsing project_id")
        			http.Error(w, err.Error(), http.StatusBadRequest)
        			return
      			}
    		}

    		if startDateStr != "" {
			startDateTmp, err := time.Parse("2006-01-02", startDateStr)
      			if err != nil {
        			fmt.Println("Error parsing start_date")
        			http.Error(w, err.Error(), http.StatusBadRequest)
        			return
      			}
			startDate = &startDateTmp
    		} else {
			startDate = nil
		}

    		if endDateStr != "" {
			endDateTmp, err := time.Parse("2006-01-02", endDateStr)
      			if err != nil {
        			fmt.Println("Error parsing end_date")
        			http.Error(w, err.Error(), http.StatusBadRequest)
        			return
      			}
			endDate = &endDateTmp
    		} else {
			endDate = nil
		}
		// Retrieve bugs from the repository
    		bugs, err := services.ListBugs(r.Context(), uint32(userIDUint), uint32(projectIDUint), startDate, endDate)
    		if err != nil {
      			fmt.Println("Error retrieving bugs")
      			http.Error(w, err.Error(), http.StatusInternalServerError)
      			return
    		}
		if len(bugs) == 0 {
      			fmt.Println("Error retrieving bugs")
      			http.Error(w, "Bugs not found", http.StatusNotFound)
      			return
		}

		bugResponses := make([]GetBugResponse, len(bugs))
    		for i, bug := range bugs {
			project, user, err := services.PopulateBug(r.Context(), bug)
			if err != nil {
      				http.Error(w, err.Error(), http.StatusInternalServerError)
      				return
			}
      			bugResponses[i] = GetBugResponse{
        			Id:          bug.Id,
        			Description: bug.Description,
				Username: strings.Join([]string{user.Name, user.Surname}, " "),
				CreationDate: bug.CreationDate,
				Project: project,
      			}
    		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bugResponses)
	}
}
