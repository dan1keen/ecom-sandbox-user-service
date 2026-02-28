package services

import (
	"fmt"
	"time"
	"user-service/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(user models.User) (string, error)
}

type jwtService struct {
	secretKey string
	tokenTTL  time.Duration
}

func NewJWTService(secretKey string) JWTService {
	return &jwtService{
		secretKey: secretKey,
		tokenTTL:  time.Hour * 24,
	}
}

func (js *jwtService) GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"exp":     time.Now().Add(js.tokenTTL).Unix(),
		"iat":     time.Now().Unix(),
		"iss":     "auth_service",
		"user_id": user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println(js.secretKey)
	return token.SignedString([]byte(js.secretKey))
}
