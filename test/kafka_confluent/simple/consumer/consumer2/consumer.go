package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
		//"security.protocol": "SASL_PLAINTEXT",
		//"sasl.mechanism":    "PLAIN",
		//"sasl.username":     "admin",
		//"sasl.password":     "6giiA6tDf9ElmUH",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"SecondTopic"}, nil)

	for i := 0; i < 2; i++ {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s, key:%s\n", msg.TopicPartition, string(msg.Value), string(msg.Key))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
