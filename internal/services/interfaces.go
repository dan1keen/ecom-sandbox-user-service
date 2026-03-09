package services

import (
	"context"
	"user-service/internal/infrastructure/db/models"
)

type UserService interface {
	GetUserById(ctx context.Context, id int) (*models.User, error)
}
