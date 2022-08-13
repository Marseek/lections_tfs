package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	//ctx := context.Background()
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "default",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	//a, err := kafka.NewAdminClientFromConsumer(c)
	//if err != nil {
	//	panic(err)
	//}
	//
	//topicResult, err := a.CreateTopics(ctx, []kafka.TopicSpecification{{
	//	Topic:             "checkTTL",
	//	NumPartitions:     1,
	//	ReplicationFactor: 1,
	//	//Config:            map[string]string{"retention.ms": "500000"},
	//}})
	//if err != nil {
	//	panic(err)
	//}
	//if topicResult[0].Error.Code() != kafka.ErrNoError {
	//	fmt.Println("! ", topicResult[0].Error)
	//	panic(topicResult[0].Error.String())
	//}
	//
	//// Read back config to validate
	//configResources := []kafka.ConfigResource{{Type: kafka.ResourceTopic, Name: "checkTTL"}}
	//describeRes, err := a.DescribeConfigs(context.Background(), configResources)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(describeRes[0].Config)

	////Change config
	////configResources = []kafka.ConfigResource{{Type: kafka.ResourceTopic, Name: "checkTTL", Config: []kafka.ConfigEntry{}}}
	//configResources[0].Config = append(configResources[0].Config, kafka.ConfigEntry{Name: "retention.ms", Value: "400000"})
	//alterRes, err := a.AlterConfigs(ctx, configResources)
	//if err != nil {
	//	panic(err)
	//}
	//if alterRes[0].Error.Code() != kafka.ErrNoError {
	//	panic(alterRes[0].Error)
	//}

	//err = c.SubscribeTopics([]string{"myTht", "myThtd"}, nil)
	//err = c.SubscribeTopics([]string{"kafka_test_channel--retry"}, nil)
	//err = c.SubscribeTopics([]string{"yatascreenshoter.tasks.listing.v1--demo_camera--retry"}, nil)
	//err = c.SubscribeTopics([]string{"yatascreenshoter.tasks.mediafiles.v1--demo_camera--retry"}, nil)
	//err = c.SubscribeTopics([]string{"yatascreenshoter.tasks.listing.v1--demo_camera--retry", "yatascreenshoter.tasks.mediafiles.v1--demo_camera--retry"}, nil)
	err = c.SubscribeTopics([]string{"yatascreenshoter.tasks.listing.v1--demo_camera", "yatascreenshoter.tasks.mediafiles.v1--demo_camera"}, nil)

	if err != nil {
		panic(err)
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s, key:|%s|, timestamp: %s, now: %s\n", msg.TopicPartition, string(msg.Value), string(msg.Key), msg.Timestamp.String(), time.Now().String())
		} else {
			//The client will automatically try to recover from all errors.
			fmt.Printf("\n\nConsumer error: %v (%v)\n\n", err, msg)
		}
	}

	c.Close()
}
