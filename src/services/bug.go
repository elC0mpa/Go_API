package services

import (
	"api-rest/src/models"
	"api-rest/src/repository"
	"context"
	"net/http"
)

func CreateBug(ctx context.Context, bug *models.Bug) error {
	return repository.InsertBug(ctx, bug)
}

func GetBugById(ctx context.Context, id uint32) (*models.Bug, error) {
	return repository.GetBugById(ctx, id)
}

func PopulateBug(ctx context.Context, writer http.ResponseWriter, bug *models.Bug) (*models.Project, *models.User, error) {
	user, err := repository.GetUserById(ctx, bug.UserId)
	if user.Id == 0 {
		http.Error(writer, "User not found", http.StatusNotFound)
		return nil, nil, nil
	} else if err != nil {
		return nil, nil, err
	}
	project, err := repository.GetProjectById(ctx, bug.ProjectId)
	if project.Id == 0 {
		http.Error(writer, "Project not found", http.StatusNotFound)
		return nil, nil, nil
	} else if err != nil {
		return nil, nil, err
	}
	return project, user, nil
}
