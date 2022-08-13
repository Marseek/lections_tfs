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
	topic := "yatascreenshoter.tasks.mediafiles.v1--demo_camera--retry"

	//for _, word := range []string{"1:Welcome", "2:to", "3:the", "Confluent", "Kafka", "Golang", "client"} {
	//for _, word := range []string{"ninth check dlx mess", "tenth check dlx mess"} {
	//for _, word := range []string{"sixth check dlx mess"} {
	for _, word := range []string{"{\"target_file_storage\":\"default\",\"camera_type\":\"dahua\",\"camera_name\":\"demo_camera\",\"camera_address\":\"http://user:pass@localhost\",\"upload_prefix\":\"/moskovsky-prospect\",\"expire_at\":\"2022-07-06T17:33:34.539926735+03:00\",\"media_file\":{\"file_path\":\"\",\"plate_number\":\"\",\"start_time\":\"0001-01-01T00:00:00Z\",\"length\":0,\"meta\":null}}\n"} {
		for i := 0; i < 3; i++ {
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(word),
				Key:            []byte("hello"),
			}, nil)
		}
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
