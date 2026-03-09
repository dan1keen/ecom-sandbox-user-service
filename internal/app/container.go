package app

import (
	"user-service/internal/http/handlers"
	"user-service/internal/repositories"
	"user-service/internal/services"

	"gorm.io/gorm"
)

// Container хранит все зависимости приложения
type Container struct {
	db *gorm.DB

	userService services.UserService
	userHandler *handlers.UserHandler
}

// NewContainer создаёт все зависимости
func NewContainer(db *gorm.DB) (c *Container) {
	c = &Container{
		db: db,
	}

	//services
	c.userService = services.NewUserService(
		repositories.NewUserRepository(db),
	)

	//handlers
	c.userHandler = handlers.NewUserHandler(c.userService)

	return c
}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}

func (c *Container) UserService() services.UserService {
	return c.userService
}
