package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application's configuration settings.
type Config struct {
	MySQLDSN   string
	RabbitMQDSN string
	GRPCPort   string
	HTTPPort   string // Added HTTPPort here
}

// Load reads the configuration values from the environment variables or default values.
func Load() Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	return Config{
		MySQLDSN:   getEnv("MYSQL_DSN", "user:password@tcp(localhost:3306)/payments"),
		RabbitMQDSN: getEnv("RABBITMQ_DSN", "amqp://guest:guest@localhost/"),
		GRPCPort:   getEnv("GRPC_PORT", ":50051"),
		HTTPPort:   getEnv("HTTP_PORT", ":8080"), // Added HTTP_PORT with default value
	}
}

// getEnv retrieves a value from the environment or returns a fallback default.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
