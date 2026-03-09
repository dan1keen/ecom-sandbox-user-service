package consumers

import (
	"encoding/json"
	"log"
	"user-service/internal/services"
)

type UserConsumer struct {
	userService services.UserService
}

func NewUserConsumer(userService services.UserService) *UserConsumer {
	return &UserConsumer{
		userService: userService,
	}
}

func (uc *UserConsumer) UserCreated(data []byte) error {
	var event map[string]interface{}
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	log.Println("Received user.created event:", event)
	// TODO: userService create user

	return nil
}
