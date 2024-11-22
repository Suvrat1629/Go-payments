package config

type Config struct {
    DBURL        string // Database URL (e.g., "postgres://user:password@localhost/dbname")
    RabbitMQURL  string // RabbitMQ URL (e.g., "amqp://guest:guest@localhost:5672/")
    // Other configuration fields...
}

func LoadConfig() Config {
    // Hardcoding values instead of using environment variables
    return Config{
        DBURL:       "postgres://aneesh:newpassword@localhost:5432/mydb?sslmode=disable", // Replace with actual DB URL
        RabbitMQURL: "amqp://guest:guest@localhost:5672/",          // Replace with actual RabbitMQ URL
    }
}
