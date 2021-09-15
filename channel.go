package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func newCh(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	errorHandler(err, "Failed to open a channel")

	log.Println("Channel opened")
	return ch
}
