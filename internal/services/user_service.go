package services

import (
	"context"
	"fmt"
	"user-service/internal/models"
	"user-service/internal/repositories"
)

type UserService interface {
	GetUserById(ctx context.Context, id int) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) GetUserById(ctx context.Context, id int) (user *models.User, err error) {
	user, err = us.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("userService.GetUserById: %w", err)
	}

	return
}
