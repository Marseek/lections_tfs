package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup12",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"myTopic", "topic_example"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s, key:|%s|\n", msg.TopicPartition, string(msg.Value), string(msg.Key))
		} else {
			//The client will automatically try to recover from all errors.
			fmt.Printf("\n\nConsumer error: %v (%v)\n\n", err, msg)
		}
	}

	c.Close()
}
