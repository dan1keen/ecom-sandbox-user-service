package rabbitmq

import (
	"log"
	"sync"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	mqOnce     sync.Once
	mqInstance *RabbitMQ
	mqErr      error
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQ(url string) *RabbitMQ {
	mqOnce.Do(func() {
		var conn *amqp.Connection
		var err error

		for i := 0; i < 5; i++ {
			conn, err = amqp.Dial(url)
			if err == nil {
				break
			}
			log.Println("Failed to connect RabbitMQ, retrying...", err)
			time.Sleep(2 * time.Second)
		}

		if err != nil {
			mqErr = err
			return
		}

		ch, err := conn.Channel()
		if err != nil {
			mqErr = err
			return
		}

		mqInstance = &RabbitMQ{Conn: conn, Channel: ch}
	})

	if mqErr != nil {
		log.Fatalf("failed to connect rabbitmq: %v", mqErr)
	}

	return mqInstance
}

func (r *RabbitMQ) Close() {
	if r.Channel != nil {
		if err := r.Channel.Close(); err != nil {
			logger.Errorf("error while closing channel: %v", err)
		}
	}

	if r.Conn != nil {
		if err := r.Conn.Close(); err != nil {
			logger.Errorf("error while closing connection: %v", err)
		}
	}
}
