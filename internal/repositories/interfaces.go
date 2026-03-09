package repositories

import (
	"context"
	"user-service/internal/infrastructure/db/models"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*models.User, error)
	GetByPhone(ctx context.Context, phone string) (*models.User, error)
}
