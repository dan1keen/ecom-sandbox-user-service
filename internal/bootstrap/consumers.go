package bootstrap

import (
	"context"
	"log"
	"user-service/internal/app"
	"user-service/internal/domain/events"
	"user-service/internal/infrastructure/rabbitmq"
)

func StartConsumers(r *rabbitmq.RabbitMQ, c *app.Container) {
	go func() {
		log.Println("Starting consumers...")

		err := r.Consume(events.UserCreated, func(body []byte) error {
			return c.UserService().UserCreated(context.Background(), body)
		})
		if err != nil {
			log.Fatalf("failed to start consumer: %v", err)
		}
	}()
}
