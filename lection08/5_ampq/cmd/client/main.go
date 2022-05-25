package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "failed to connect!")

	ch, err := conn.Channel()
	failOnError(err, "failed to create channel!")

	q, err := ch.QueueDeclare(
		"myQueue",
		false,
		false,
		false,
		false,
		nil)
	failOnError(err, "failed to create queue!")

	msgCh, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil)

	for msg := range msgCh {
		log.Println(string(msg.Body))
	}
	failOnError(err, "failed to receive a message!")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
