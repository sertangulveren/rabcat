package main

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func newAMQP() *amqp.Connection {
	amqpUrl := os.Getenv("AMQP_URL")
	if amqpUrl == "" {
		panic("AMQP_URL must be set")
	}
	conn, err := amqp.Dial(amqpUrl)

	errorHandler(err, "Failed to connect to RabbitMQ")

	log.Println("Connected to RabbitMQ")
	return conn
}
