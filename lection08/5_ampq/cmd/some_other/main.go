package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/streadway/amqp"
)

type RabbitClient struct {
	sendConn *amqp.Connection
	recConn  *amqp.Connection
	sendChan *amqp.Channel
	recChan  *amqp.Channel
}

type ConfConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

var config = ConfConfig{Username: "guest", Password: "guest", Port: "5672", Host: "localhost"}

func main() {
	var rcl RabbitClient

	conn, err := rcl.connect(true, true)
	failOnError(err, "failed to establish connection!")

	chann, err := rcl.channel(true, true)

}

// Подписываемся исходя из имени очереди
func (rcl *RabbitClient) Consume(n string, f func(interface{}) error) {
	for {
		for {
			_, err := rcl.channel(true, true)
			if err == nil {
				break
			}
		}
		log.Printf("--- connected to consume '%s' ---\r\n", n)
		q, err := rcl.recChan.QueueDeclare(
			n,
			true,
			false,
			false,
			false,
			amqp.Table{"x-queue-mode": "lazy"},
		)
		if err != nil {
			log.Println("--- failed to declare a queue, trying to reconnect ---")
			continue
		}
		connClose := rcl.recConn.NotifyClose(make(chan *amqp.Error))
		connBlocked := rcl.recConn.NotifyBlocked(make(chan amqp.Blocking))
		chClose := rcl.recChan.NotifyClose(make(chan *amqp.Error))
		m, err := rcl.recChan.Consume(
			q.Name,
			uuid.NewV4().String(),
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Println("--- failed to consume from queue, trying again ---")
			continue
		}
		shouldBreak := false
		for {
			if shouldBreak {
				break
			}
			select {
			case _ = <-connBlocked:
				log.Println("--- connection blocked ---")
				shouldBreak = true
				break
			case err = <-connClose:
				log.Println("--- connection closed ---")
				shouldBreak = true
				break
			case err = <-chClose:
				log.Println("--- channel closed ---")
				shouldBreak = true
				break
			case d := <-m:
				err := f(d.Body)
				if err != nil {
					_ = d.Ack(false)
					break
				}
				_ = d.Ack(true)
			}
		}
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (rcl *RabbitClient) channel(isRec, recreate bool) (*amqp.Channel, error) {
	if recreate {
		if isRec {
			rcl.recChan = nil
		} else {
			rcl.sendChan = nil
		}
	}
	if isRec && rcl.recConn == nil {
		rcl.recChan = nil
	}
	if !isRec && rcl.sendConn == nil {
		rcl.recChan = nil
	}
	if isRec && rcl.recChan != nil {
		return rcl.recChan, nil
	} else if !isRec && rcl.sendChan != nil {
		return rcl.sendChan, nil
	}
	for {
		_, err := rcl.connect(isRec, recreate)
		if err == nil {
			break
		}
	}
	var err error
	if isRec {
		rcl.recChan, err = rcl.recConn.Channel()
	} else {
		rcl.sendChan, err = rcl.sendConn.Channel()
	}
	if err != nil {
		log.Println("--- could not create channel ---")
		time.Sleep(1 * time.Second)
		return nil, err
	}
	if isRec {
		return rcl.recChan, err
	} else {
		return rcl.sendChan, err
	}
}

func (rcl *RabbitClient) connect(isRec, reconnect bool) (*amqp.Connection, error) {
	if reconnect {
		if isRec {
			rcl.recConn = nil
		} else {
			rcl.sendConn = nil
		}
	}
	if isRec && rcl.recConn != nil {
		return rcl.recConn, nil
	} else if !isRec && rcl.sendConn != nil {
		return rcl.sendConn, nil
	}
	var c string
	if config.Username == "" {
		c = fmt.Sprintf("amqp://%s:%s/", config.Host, config.Port)
	} else {
		c = fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Username, config.Password, config.Host, config.Port)
	}
	conn, err := amqp.Dial(c)
	if err != nil {
		log.Printf("\r\n--- could not create a conection ---\r\n")
		time.Sleep(1 * time.Second)
		return nil, err
	}
	if isRec {
		rcl.recConn = conn
		return rcl.recConn, nil
	} else {
		rcl.sendConn = conn
		return rcl.sendConn, nil
	}
}

// Публикуем байтовый массив в очередь
func (rcl *RabbitClient) Publish(n string, b []byte) {
	r := false
	for {
		for {
			_, err := rcl.channel(false, r)
			if err == nil {
				break
			}
		}
		q, err := rcl.sendChan.QueueDeclare(
			n,
			true,
			false,
			false,
			false,
			amqp.Table{"x-queue-mode": "lazy"},
		)
		if err != nil {
			log.Println("--- failed to declare a queue, trying to resend ---")
			r = true
			continue
		}
		err = rcl.sendChan.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				MessageId:    uuid.NewV4().String(),
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         b,
			})
		if err != nil {
			log.Println("--- failed to publish to queue, trying to resend ---")
			r = true
			continue
		}
		break
	}
}
