package services

import (
	"api-rest/src/models"
	"api-rest/src/repository"
	"context"
)

func CreateProject(ctx context.Context, project *models.Project) error {
	return repository.InsertProject(ctx, project)
}

func GetProjectById(ctx context.Context, id uint32) (*models.Project, error) {
	return repository.GetProjectById(ctx, id)
}
