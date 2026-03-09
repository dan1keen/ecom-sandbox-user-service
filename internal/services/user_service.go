package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"user-service/internal/infrastructure/db/models"
	"user-service/internal/repositories"
)

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

func (us *userService) UserCreated(ctx context.Context, data []byte) error {
	var event map[string]interface{}
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	log.Println("Received user.created event:", event)
	// TODO: create user func

	return nil
}
