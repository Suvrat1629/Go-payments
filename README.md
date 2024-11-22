# Go Payments System

This is a simple payment system written in Go, which allows users to make payments, check payment status, and receive status updates via gRPC and HTTP APIs. It also integrates with RabbitMQ to handle asynchronous payment status updates.

## Table of Contents

1. [File Structure](#file-structure)
2. [Overview](#overview)
3. [Configuration](#configuration)
   - [1. Database URL](#1-database-url)
   - [2. RabbitMQ URL](#2-rabbitmq-url)
   - [3. Config File Example](#3-config-file-example)
4. [Dependencies](#dependencies)
5. [Running the Application](#running-the-application)
   - [1. Clone the repository](#1-clone-the-repository)
   - [2. Install dependencies](#2-install-dependencies)
   - [3. Set up the Database](#3-set-up-the-database)
   - [4. Start RabbitMQ](#4-start-rabbitmq)
   - [5. Run the Application](#5-run-the-application)
6. [HTTP Endpoints](#http-endpoints)
   - [1. Make Payment](#1-make-payment-http)
   - [2. Get Payment Status](#2-get-payment-status-http)
7. [gRPC Service](#grpc-service)

---

## File Structure

```
/internal
  ├── /api
  │    └── /grpc
  │         ├── payment_handler.go
  │         ├── proto
  │         └── main.go
  ├── /db
  ├── /rabbitmq
  └── /config
/proto
```

---

## Overview

The system supports:
- **Making payments**: Using gRPC (`/make-payment` endpoint in the Echo server).
- **Checking payment status**: Using gRPC (`/get-payment-status` endpoint in the Echo server).
- **Asynchronous updates**: Using RabbitMQ to listen for updates on payment statuses.

### Key Components:
- **gRPC Server**: Handles payment requests and status updates.
- **Database (PostgreSQL)**: Stores payments and their statuses.
- **RabbitMQ**: Sends and receives payment status updates.

---

## Configuration

### 1. Database URL

The database connection URL is located in the `/internal/config/config.go` file. You can update the following line with the appropriate database connection URL:

```go
DBURL: "postgres://username:password@hostname:port/database",
```

Modify the connection string according to your database configuration:
- **PostgreSQL** connection string example:  
  `postgres://user:password@localhost:5432/payments`
  
  Where:
  - `user`: the PostgreSQL username
  - `password`: the PostgreSQL password
  - `localhost`: the hostname of the database server
  - `5432`: the PostgreSQL port (change if necessary)
  - `payments`: the database name

Example for **MySQL**:
```go
DBURL: "mysql://username:password@hostname:port/database",
```

### 2. RabbitMQ URL

Similarly, RabbitMQ's connection URL is set in the same `config.go` file. You can change the RabbitMQ URL by modifying the following line:

```go
RabbitMQURL: "amqp://guest:guest@localhost:5672/",
```

Update this URL to reflect your RabbitMQ configuration:
- The default RabbitMQ connection URL is `amqp://guest:guest@localhost:5672/`, where:
  - `guest`: the username
  - `guest`: the password
  - `localhost`: the hostname of the RabbitMQ server
  - `5672`: the default RabbitMQ port

If you're running RabbitMQ on a custom host, change `localhost` to the actual hostname or IP address.

### 3. Config File Example

Here is how the `LoadConfig` function in `/internal/config/config.go` should look like after modifying the URLs:

```go
package config

import "log"

// Config struct holds the configuration values
type Config struct {
	DBURL       string
	RabbitMQURL string
}

// LoadConfig loads the configuration for the application
func LoadConfig() *Config {
	return &Config{
		// Change the database URL here
		DBURL:      "postgres://user:password@localhost:5432/payments", // Replace with your database URL
		// Change RabbitMQ URL here
		RabbitMQURL: "amqp://guest:guest@localhost:5672/", // Replace with your RabbitMQ URL
	}
}
```

---

## Dependencies

The following dependencies are required to run the application:

- Go 1.18 or higher
- PostgreSQL database
- RabbitMQ server
- gRPC libraries (for defining and using gRPC services)
- Protobuf compiler (`protoc`)

You can install dependencies using Go modules:

```bash
go mod tidy
```

---

## Running the Application

### 1. Clone the repository

```bash
git clone https://github.com/your-username/go-payments.git
cd go-payments
```

### 2. Install dependencies

Run `go mod tidy` to install all the dependencies.

### 3. Set up the Database

Ensure your PostgreSQL database is set up and running. You may need to create the `payments` table manually, or you can write a migration script based on the `SavePayment` function.

### 4. Start RabbitMQ

Make sure RabbitMQ is installed and running on your machine or use a cloud-based RabbitMQ service. The default URL is `amqp://guest:guest@localhost:5672/`, but you can update it in the `config.go` file as mentioned earlier.

### 5. Run the Application

After modifying the configuration, you can start the application with:

```bash
go run internal/api/grpc/main.go
```

This will start both the **gRPC server** on port `50051` and the **HTTP server** on port `8080`.

---

## HTTP Endpoints

The following HTTP endpoints are exposed by the application:

- `POST /make-payment`: Make a payment (calls the `MakePayment` gRPC service).
- `POST /get-payment-status`: Get the status of a payment (calls the `GetPaymentStatus` gRPC service).

### 1. Make Payment (HTTP)

```bash
curl -X POST http://localhost:8080/make-payment -d '{"sender_id": "user1", "receiver_id": "user2", "amount": 100.00, "currency": "USD"}' -H "Content-Type: application/json"
```

### 2. Get Payment Status (HTTP)

```bash
curl -X POST http://localhost:8080/get-payment-status -d '{"transaction_id": "some-transaction-id"}' -H "Content-Type: application/json"
```

---

## gRPC Service

The application also exposes the following gRPC services:

- `MakePayment(PaymentRequest)`: Processes a payment and returns a transaction ID.
- `GetPaymentStatus(PaymentStatusRequest)`: Retrieves the payment status for a given transaction ID.
- `UpdatePaymentStatus(PaymentUpdateRequest)`: Updates the status of a payment.

---

