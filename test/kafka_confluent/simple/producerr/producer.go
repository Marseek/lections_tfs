package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	//Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failedAAAAAAAAAAAAAAAAA: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v, value: %s\n", ev.TopicPartition, string(ev.Value))
				}
			case kafka.Error:
				fmt.Printf("Delivery failedAAAAAAAAAAAAAAAAA: %v - code\n", ev.Code())
			default:
				fmt.Printf("event type: %T\n", ev)
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "myTopic"
	//for _, word := range []string{"1:Welcome", "2:to", "3:the", "Confluent", "Kafka", "Golang", "client"} {
	for _, word := range []string{"1:Welcome", "2:to"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
			Key:            []byte("hello"),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
