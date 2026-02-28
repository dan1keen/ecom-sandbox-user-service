package services

import (
	"context"
	"errors"
	"user-service/internal/dto"
	"user-service/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, credentials dto.LoginRequest) (*string, error)
}

type authService struct {
	jwtService JWTService
	userRepo   repositories.UserRepository
}

func NewAuthService(jwtService JWTService, userRepo repositories.UserRepository) AuthService {
	return &authService{
		jwtService: jwtService,
		userRepo:   userRepo,
	}
}

func (as *authService) Login(ctx context.Context, credentials dto.LoginRequest) (*string, error) {
	user, err := as.userRepo.GetByPhone(ctx, credentials.PhoneNumber)
	if err != nil {
		return nil, errors.New("invalid phone or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return nil, errors.New("invalid phone or password")
	}

	token, err := as.jwtService.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
