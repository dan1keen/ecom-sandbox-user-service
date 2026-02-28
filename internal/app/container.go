package app

import (
	"user-service/config"
	"user-service/internal/http/handlers"
	"user-service/internal/repositories"
	"user-service/internal/services"

	"gorm.io/gorm"
)

// Container хранит все зависимости приложения
type Container struct {
	db *gorm.DB

	userRepo    repositories.UserRepository
	userService services.UserService
	authService services.AuthService
	jwtService  services.JWTService
	userHandler *handlers.UserHandler
	authHandler *handlers.AuthHandler
}

// NewContainer создаёт все зависимости
func NewContainer(cfg *config.Config, db *gorm.DB) (c *Container) {
	c = &Container{
		db: db,
	}

	//repositories
	c.userRepo = repositories.NewUserRepository(db)

	//services
	c.userService = services.NewUserService(c.userRepo)
	c.jwtService = services.NewJWTService(cfg.JWTSecret)
	c.authService = services.NewAuthService(c.jwtService, c.userRepo)

	//handlers
	c.userHandler = handlers.NewUserHandler(c.userService)
	c.authHandler = handlers.NewAuthHandler(c.authService)

	return c
}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}

func (c *Container) AuthHandler() *handlers.AuthHandler {
	return c.authHandler
}
