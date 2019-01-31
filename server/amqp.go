package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://michong:michong@192.168.1.18:5672/dev")
	failOnError(err, "Failed to connect RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
	"imTask",
	true,
	false,
	false,
	false,
	nil,
	)
	failOnError(err, "Failed to declare a queue")

	forever := make(chan bool)
	message, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range message {
			log.Printf("Recevied a message: %s", d.Body)
		}
	}()
	<-forever
}
