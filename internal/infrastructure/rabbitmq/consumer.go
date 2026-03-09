package rabbitmq

import (
	"log"
)

func (r *RabbitMQ) Consume(queue string, handler func([]byte) error) error {
	_, err := r.Channel.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := r.Channel.Consume(
		queue,
		"",
		true,  // auto-ack
		false, // exclusive
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			if err = handler(d.Body); err != nil {
				log.Println("Failed to handle message:", err)
			}
		}
	}()

	return nil
}
