package main

import (
	"fmt"

	"github.com/jailtonjunior94/fullcycle-goevents/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	messages := make(chan amqp.Delivery)
	go rabbitmq.Consumer(ch, messages, "minhafila")

	for message := range messages {
		fmt.Println(string(message.Body))
		message.Ack(false)
	}
}
