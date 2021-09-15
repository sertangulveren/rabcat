package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func newExchange(ch *amqp.Channel) {
	err := ch.ExchangeDeclare(
		"chat",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	errorHandler(err, "Failed to declare a exchange")

	log.Println("Exchange declared")
}
