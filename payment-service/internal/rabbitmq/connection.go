package rabbitmq

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"
)

// Connection wraps the AMQP connection to simplify usage.
type Connection struct {
	*amqp.Connection
	channel *amqp.Channel // Persistent channel for reuse
}

// Connect establishes a connection to RabbitMQ and creates a persistent channel.
func Connect() (*Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // Replace with your RabbitMQ URL if needed
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	// Create a persistent channel for reuse
	ch, err := conn.Channel()
	if err != nil {
		conn.Close() // Close the connection if we fail to create the channel
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	// Return the custom connection object with the channel
	return &Connection{
		Connection: conn,
		channel:    ch,
	}, nil
}

// Publish sends a message to a RabbitMQ queue in Protobuf format.
func (conn *Connection) Publish(queueName string, message proto.Message) error {
	// Reuse the persistent channel
	ch := conn.channel

	// Declare the queue (ensure the queue exists)
	q, err := ch.QueueDeclare(
		queueName, // Queue name
		false,     // Non-durable (should match the existing queue)
		false,     // Non-auto-delete
		false,     // Non-exclusive
		false,     // Non-passive
		nil,       // No additional arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	// Marshal the Protobuf message into a byte slice
	body, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	// Publish the message to the queue
	err = ch.Publish(
		"",        // Default exchange
		q.Name,    // Queue name
		false,     // If false, the message will not be persisted if RabbitMQ crashes
		false,     // If false, the message will not be delivered to consumers if they are not active
		amqp.Publishing{
			ContentType: "application/protobuf",
			Body:        body, // Message body in Protobuf format
		},
	)
	return err
}

// Consume listens for messages from a specified queue and unmarshals them into Protobuf.
func (conn *Connection) Consume(queueName string) (<-chan amqp.Delivery, error) {
	// Reuse the persistent channel
	ch := conn.channel

	// Declare the queue to ensure it exists
	_, err := ch.QueueDeclare(
		queueName, // Queue name
		false,     // Non-durable
		false,     // Non-auto-delete
		false,     // Non-exclusive
		false,     // Non-passive
		nil,       // No additional arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %v", err)
	}

	// Start consuming messages from the queue
	msgs, err := ch.Consume(
		queueName, // Queue name
		"",        // Consumer name (leave empty for automatic name generation)
		true,      // Auto-acknowledge messages
		false,     // Don't use exclusive access to the queue
		false,     // Don't make the consumer a priority consumer
		false,     // No wait for delivery
		nil,       // No additional arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to start consuming messages: %v", err)
	}

	return msgs, nil
}

// Close shuts down the RabbitMQ connection and channel.
func (conn *Connection) Close() {
	if err := conn.channel.Close(); err != nil {
		log.Printf("Failed to close channel: %v", err)
	}
	if err := conn.Connection.Close(); err != nil {
		log.Printf("Failed to close connection: %v", err)
	}
}
