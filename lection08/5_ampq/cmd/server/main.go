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
	//
	//q, err := ch.QueueDeclare(
	//	"myQueue",
	//	false,
	//	false,
	//	false,
	//	false,
	//	nil)
	//failOnError(err, "failed to create queue!")

	err = ch.Publish(
		"test-fanout",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello, test-fanout message to queue!"),
		})
	failOnError(err, "failed to send message!")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
