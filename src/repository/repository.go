package repository

import (
	"api-rest/src/models"
	"context"
	"time"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id uint32) (*models.User, error)

	InsertProject(ctx context.Context, project *models.Project) error
	GetProjectById(ctx context.Context, id uint32) (*models.Project, error)

	InsertBug(ctx context.Context, bug *models.Bug) error
	GetBugById(ctx context.Context, id uint32) (*models.Bug, error)
	ListBugs(ctx context.Context, userId, projectId uint32, startDate, endDate *time.Time) ([]*models.Bug, error)

	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
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

func InsertBug(ctx context.Context, bug *models.Bug) error {
	return implementation.InsertBug(ctx, bug)
}

func GetBugById(ctx context.Context, id uint32) (*models.Bug, error) {
	return implementation.GetBugById(ctx, id)
}

func ListBugs(ctx context.Context, userId, projectId uint32, startDate, endDate *time.Time) ([]*models.Bug, error) {
	return implementation.ListBugs(ctx, userId, projectId, startDate, endDate)
}

func Close() error {
	return implementation.Close()
}
