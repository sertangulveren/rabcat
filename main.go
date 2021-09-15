package main

import (
	"log"
	"os"

	"github.com/fatih/color"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/crypto/ssh/terminal"
)

var authorColor = color.New(color.Underline).Add(color.Bold).SprintFunc()

func receiveMessages(ch *amqp.Channel, name string) {
	msgs := newConsumer(ch, name)
	go func() {
		log.Println("Waiting messages")
		for d := range msgs {
			msg := parseMessage(d.Body)
			log.Printf("%s %s \n", authorColor(msg.Sender+":"), msg.Content)
		}
	}()
}

func sendMessage(msg *Message, ch *amqp.Channel) {
	err := ch.Publish(
		"chat",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg.generatePayload(),
		})
	if err != nil {
		color.New(color.FgRed).Add(color.Bold).Printf("Failed: %s\n", err)
	}
}

func listenUser(ch *amqp.Channel) {
	author := getAuthor()

	log.Println("Listening user")

	for {
		// what a worst method
		termRes, err := terminal.ReadPassword(0)

		if err != nil {
			continue
		}

		if string(termRes) == "bye" {
			color.New(color.FgCyan).Add(color.Bold).Println("Good bye!")
			sendMessage(&Message{Sender: author, Content: "Has gone!"}, ch)
			os.Exit(0)
		}

		if len(termRes) < 1 {
			continue
		}
		sendMessage(&Message{Sender: author, Content: string(termRes)}, ch)
	}
}

func main() {
	conn := newAMQP()
	defer conn.Close()

	ch := newCh(conn)
	defer ch.Close()

	newExchange(ch)

	q := newQueue(ch)

	receiveMessages(ch, q.Name)

	sendMessage(&Message{Sender: getAuthor(), Content: "Joined!"}, ch)

	listenUser(ch)
}
