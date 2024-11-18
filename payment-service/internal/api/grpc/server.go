package grpc

import (
	"context"
	"log"

	"database/sql"

	"github.com/Go-payments/internal/rabbitmq"
	pb "github.com/Go-payments/internal/proto"
	"github.com/google/uuid"
)

// Define the implementation struct correctly
type PaymentServiceServerImpl struct {
	pb.UnimplementedPaymentServiceServer
	db         *sql.DB
	rabbitConn *rabbitmq.Connection
}

// Constructor for the PaymentServiceServerImpl
func NewPaymentServiceServer(db *sql.DB, rabbitConn *rabbitmq.Connection) *PaymentServiceServerImpl {
	return &PaymentServiceServerImpl{db: db, rabbitConn: rabbitConn}
}

// Implementing the MakePayment method for the server
func (s *PaymentServiceServerImpl) MakePayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	// Generate a unique Transaction ID
	transactionID := uuid.New().String()

	// Save payment details to the database
	_, err := s.db.Exec("INSERT INTO payments (transaction_id, sender_id, receiver_id, amount, currency) VALUES (?, ?, ?, ?, ?)",
		transactionID, req.SenderId, req.ReceiverId, req.Amount, req.Currency)
	if err != nil {
		log.Printf("Database error: %v", err)
		return &pb.PaymentResponse{TransactionId: "", Status: "FAILED", Message: "Database error"}, nil
	}

	// Prepare the event to publish to RabbitMQ
	event := map[string]interface{}{
		"transaction_id": transactionID,
		"status":         "SUCCESS",
		"amount":         req.Amount,
		"currency":       req.Currency,
	}

	// Publish the event to RabbitMQ
	if err := s.rabbitConn.Publish("payment_events", event); err != nil {
		log.Printf("RabbitMQ publish error: %v", err)
		return &pb.PaymentResponse{TransactionId: transactionID, Status: "FAILED", Message: "Event publish error"}, nil
	}

	// Return success response
	return &pb.PaymentResponse{
		TransactionId: transactionID,
		Status:        "SUCCESS",
		Message:       "Payment processed successfully",
	}, nil
}
