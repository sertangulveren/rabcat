package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func newQueue(ch *amqp.Channel) *amqp.Queue {
	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	errorHandler(err, "Failed to declare a queue")

	log.Println("Queue declared")

	err = ch.QueueBind(
		q.Name,
		"",
		"chat",
		false,
		nil,
	)

	errorHandler(err, "Failed to bind queue")

	log.Println("Queue bound with exchange")

	return &q
}
