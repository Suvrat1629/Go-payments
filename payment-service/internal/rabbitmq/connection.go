package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/Go-payments/internal/config"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"
)

type Connection struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func Connect(cfg config.Config) (*Connection, error) {
	conn, err := amqp.Dial(cfg.RabbitMQDSN)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ channel: %v", err)
	}
	return &Connection{conn: conn, channel: channel}, nil
}

// Publish serializes the message body and sends it to RabbitMQ.
func (c *Connection) Publish(queueName string, body interface{}) error {
	// Serialize the body (Assuming protobuf serialization)
	var bodyBytes []byte
	var err error

	// If the body is a proto message, serialize using proto.Marshal
	switch v := body.(type) {
	case proto.Message:
		bodyBytes, err = proto.Marshal(v)
		if err != nil {
			log.Printf("Failed to serialize message to protobuf: %v", err)
			return err
		}
		// If the body is a generic interface{}, serialize it to JSON
	default:
		bodyBytes, err = json.Marshal(v)
		if err != nil {
			log.Printf("Failed to serialize message to JSON: %v", err)
			return err
		}
	}

	// Publish the message to the specified queue
	err = c.channel.Publish(
		"",        // exchange
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json", // or "application/protobuf" if using protobuf
			Body:        bodyBytes,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message to RabbitMQ: %v", err)
		return err
	}

	return nil
}

func (c *Connection) Close() {
	c.channel.Close()
	c.conn.Close()
}
