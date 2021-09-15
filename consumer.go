package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func newConsumer(ch *amqp.Channel, qName string) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		qName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	errorHandler(err, "Failed to register a consumer")

	log.Println("Consumer registered")
	return msgs
}
