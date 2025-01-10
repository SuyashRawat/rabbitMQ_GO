// internal/rabbitmq/rabbitmq.go
package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// Connect function to connect to RabbitMQ
func Connect(rabbitMQURL string) (*amqp.Connection, *amqp.Channel, error) {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		return nil, nil, err
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to create a channel: %s", err)
		return nil, nil, err
	}

	return conn, ch, nil
}

// Close function to close the connection and channel
func Close(conn *amqp.Connection, ch *amqp.Channel) {
	err := ch.Close()
	if err != nil {
		log.Fatalf("Failed to close channel: %s", err)
	}
	err = conn.Close()
	if err != nil {
		log.Fatalf("Failed to close connection: %s", err)
	}
}
