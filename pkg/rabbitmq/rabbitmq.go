package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch, nil
}

func Consumer(ch *amqp.Channel, out chan<- amqp.Delivery, queue string) error {
	messages, err := ch.Consume(
		queue,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil
	}

	for message := range messages {
		out <- message
	}

	return nil
}

func Publish(ch *amqp.Channel, body string, exchange string) error {
	err := ch.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	return err
}
