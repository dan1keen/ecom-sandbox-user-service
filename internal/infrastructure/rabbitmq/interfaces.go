package rabbitmq

type Publisher interface {
	Publish(queue string, message interface{}) error
}

type Consumer interface {
	Consume(queue string, handler func([]byte) error) error
}
