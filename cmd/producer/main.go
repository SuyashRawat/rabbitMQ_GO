// cmd/producer/main.go
package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"github.com/suyashrawat/go-rabbit/internal/rabbitmq"
)

func main() {
	rabbitMQURL := "amqp://guest:guest@localhost:5672/" // Replace with your RabbitMQ server URL
	queueName := "testQueue"

	// Establish connection and create channel
	conn, ch, err := rabbitmq.Connect(rabbitMQURL)
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitmq.Close(conn, ch)

	// Declare a queue
	_, err = ch.QueueDeclare(
		queueName, // Queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Send a message
	message := "Hello, prate!"
	err = ch.Publish(
		"",        // default exchange
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}

	fmt.Println("Message sent:", message)
}
