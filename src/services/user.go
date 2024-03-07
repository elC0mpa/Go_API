package services

import (
	"api-rest/src/models"
	"api-rest/src/repository"
	"context"
)

func CreateUser(ctx context.Context, user *models.User) error {
	return repository.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id uint32) (*models.User, error) {
	return repository.GetUserById(ctx, id)
}
