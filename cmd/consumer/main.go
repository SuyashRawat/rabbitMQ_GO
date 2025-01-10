// cmd/consumer/main.go
package main

import (
	"fmt"
	"log"

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

	// Declare the queue (ensure it's the same as the producer)
	_, err = ch.QueueDeclare(
		queueName,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Set up a consumer
	msgs, err := ch.Consume(
		queueName, // queue name
		"",        // consumer name
		true,      // auto-acknowledge
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	// Consume messages
	for msg := range msgs {
		fmt.Printf("Received a message: %s\n", msg.Body)
	}
}
