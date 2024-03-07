package repository

import (
	"api-rest/src/models"
	"context"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id uint32) (*models.User, error)

	InsertProject(ctx context.Context, project *models.Project) error
	GetProjectById(ctx context.Context, id uint32) (*models.Project, error)
	Close() error
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id uint32) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func InsertProject(ctx context.Context, project *models.Project) error {
	return implementation.InsertProject(ctx, project)
}

func GetProjectById(ctx context.Context, id uint32) (*models.Project, error) {
	return implementation.GetProjectById(ctx, id)
}

func Close() error {
	return implementation.Close()
}
