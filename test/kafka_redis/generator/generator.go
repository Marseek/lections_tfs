/*
	генератор чисел для факторизации
	программа генерирует случайные числа и записывает их в очередь kafka
*/
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const topic = "Generated"

func main() {

	// инициализация генератора случайных чисел
	rand.Seed(time.Now().Unix())
	const max = 2345678901234

	// клиент очереди
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {

		// генерируем случайное число
		n := rand.Intn(max)

		err = conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		if err != nil {
			log.Fatal(err)
		}

		// пишем чсило в очередь
		_, err = conn.WriteMessages(
			kafka.Message{Value: []byte(strconv.Itoa(n))},
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Written to topic:", n)

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(4000)))

	}

}
