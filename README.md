# Go Payments System

This project is a payment processing system built with Go, gRPC, PostgreSQL, RabbitMQ, and Protobuf. It supports basic payment processing functionalities, such as making payments, checking payment status, and updating payment statuses. The system also integrates with RabbitMQ to handle payment updates asynchronously.

## Table of Contents

- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)
- [Setup Instructions](#setup-instructions)
- [API Endpoints](#api-endpoints)
- [gRPC Services](#grpc-services)
- [How It Works](#how-it-works)

## Project Structure

The project is organized as follows:

```
/Go-payments
├── /internal
│   ├── /api
│   │   └── /grpc
│   │       ├── payment_handler.go        # Handles payment logic and communication with RabbitMQ, DB
│   │       ├── proto                     # Protobuf definitions and gRPC server implementation
│   │       └── main.go                   # Main entry point to start the server
│   ├── /db
│   │   └── db.go                        # Database connection and queries (PostgreSQL)
│   ├── /rabbitmq
│   │   └── rabbitmq.go                  # RabbitMQ connection and message publishing/consuming logic
│   ├── /config
│   │   └── config.go                    # Configuration loading (e.g., DB URL, RabbitMQ URL)
│   └── /proto
│       └── grpc
│           └── payment.proto            # Protobuf definition of payment-related messages and services
└── /migrations                           # Database migrations (optional)
└── /scripts                              # Helper scripts (e.g., to setup environment or run tests)
└── README.md                            # Project overview and setup instructions
```

- **internal/api/grpc/payment_handler.go**: Handles gRPC requests related to payments, including making payments, updating payment statuses, and querying payment status.
- **internal/db/db.go**: Manages database connections and queries to store and retrieve payment information.
- **internal/rabbitmq/rabbitmq.go**: Provides functionality for connecting to RabbitMQ and publishing/consuming messages related to payment status updates.
- **internal/config/config.go**: Loads application configuration such as database connection strings and RabbitMQ URLs.
- **internal/proto/grpc/payment.proto**: Defines the Protobuf messages and gRPC service for payment processing.

## Technologies Used

- **Go**: Programming language used to implement the backend services.
- **gRPC**: Remote Procedure Call (RPC) framework used to handle client-server communication.
- **PostgreSQL**: Relational database to store payment information.
- **RabbitMQ**: Message broker to handle asynchronous payment updates.
- **Protobuf**: Serialization format for communication between services (gRPC).

## Setup Instructions

### Prerequisites

1. Install Go (version 1.18+)
2. Install PostgreSQL and RabbitMQ on your local machine, or use a cloud service.
3. Install required Go dependencies.

### Clone the repository

```bash
git clone https://github.com/Go-payments/Go-payments.git
cd Go-payments
```

### Install dependencies

Run the following command to install the necessary dependencies:

```bash
go mod tidy
```

### Configuration

Configure your environment by updating the values in the `config.go` file under `/internal/config/`:

- **DBURL**: Your PostgreSQL connection string (e.g., `postgres://user:password@localhost:5432/payments`)
- **RabbitMQURL**: URL of your RabbitMQ server (e.g., `amqp://guest:guest@localhost:5672/`)

Alternatively, you can set these configurations as environment variables:

```bash
export DB_URL="postgres://user:password@localhost:5432/payments"
export RABBITMQ_URL="amqp://guest:guest@localhost:5672/"
```

### Database Setup

Ensure you have a PostgreSQL database running and a table to store payments:

```sql
CREATE TABLE payments (
    transaction_id VARCHAR PRIMARY KEY,
    sender_id VARCHAR NOT NULL,
    receiver_id VARCHAR NOT NULL,
    amount FLOAT NOT NULL,
    status VARCHAR NOT NULL
);
```

### Running the Application

1. **Start the gRPC Server**:

```bash
go run internal/api/grpc/main.go
```

The gRPC server will run on port `50051` by default.

2. **Start the HTTP Server (Echo)**:

In a separate terminal window, start the Echo HTTP server to expose REST endpoints:

```bash
go run internal/api/grpc/main.go
```

The HTTP server will run on port `8080`.

### Testing the API

1. **Make a Payment**:

Send a `POST` request to `http://localhost:8080/make-payment` with the following JSON body:

```json
{
  "sender_id": "user123",
  "receiver_id": "user456",
  "amount": 100.50,
  "currency": "USD"
}
```

2. **Check Payment Status**:

Send a `POST` request to `http://localhost:8080/get-payment-status` with the following JSON body:

```json
{
  "transaction_id": "your_transaction_id"
}
```

### gRPC API

The system also exposes gRPC services for payment processing. To interact with the gRPC server, you can use the `grpc` client as shown in the examples:

- **MakePayment**: `rpc MakePayment(PaymentRequest) returns (PaymentResponse);`
- **GetPaymentStatus**: `rpc GetPaymentStatus(PaymentStatusRequest) returns (PaymentStatusResponse);`
- **UpdatePaymentStatus**: `rpc UpdatePaymentStatus(PaymentUpdateRequest) returns (PaymentUpdateResponse);`

You can generate Go client code using the `protoc` tool from the provided `.proto` files.

### Database Migrations

You can optionally add database migrations under the `/migrations` directory to handle schema changes.

## How It Works

1. **Make a Payment**:
   - When a payment request is made, the system validates the request (e.g., amount must be greater than zero).
   - It generates a unique transaction ID and stores the payment details in the PostgreSQL database with a "PENDING" status.
   - The system publishes a message to RabbitMQ with the payment status update, which can be consumed by other services or applications to perform follow-up actions.

2. **Payment Status Update**:
   - The system listens to RabbitMQ for status updates (e.g., successful or failed payments).
   - When a message is received, it updates the status of the corresponding payment in the database.

3. **Check Payment Status**:
   - The system allows querying the payment status via a REST or gRPC endpoint.
