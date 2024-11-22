package grpc_server

import (
	"context"
	"database/sql" // Importing the SQL package to handle database queries
	"fmt"
	"log"
	"github.com/Go-payments/internal/db"
	"github.com/Go-payments/internal/rabbitmq"
	"github.com/google/uuid" // For generating unique transaction IDs
	"google.golang.org/protobuf/proto" // Import the proto package for unmarshalling
	pb "github.com/Go-payments/internal/proto/grpc" // Import the generated proto package
)

// PaymentHandler structure to handle payment logic
type PaymentHandler struct {
	pb.UnimplementedPaymentServiceServer // Embeds the unimplemented methods to allow for graceful upgrades
	DB         *db.DB
	RabbitConn *rabbitmq.Connection
}

// NewPaymentHandler creates and returns a new PaymentHandler instance
func NewPaymentHandler(dbConn *db.DB, rabbitConn *rabbitmq.Connection) *PaymentHandler {
	return &PaymentHandler{
		DB:         dbConn,
		RabbitConn: rabbitConn,
	}
}

// MakePayment processes a payment request and generates a transaction ID
func (h *PaymentHandler) MakePayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Printf("Processing payment from %s to %s with amount %f", req.SenderId, req.ReceiverId, req.Amount)

	// Validate input (e.g., check if amount is greater than zero)
	if req.Amount <= 0 {
		return nil, fmt.Errorf("invalid payment amount")
	}

	// Generate a unique transaction ID using UUID
	transactionID := uuid.New().String()

	// Save payment in the database along with the transaction ID
	err := h.DB.SavePayment(req.SenderId, req.ReceiverId, float64(req.Amount), transactionID)
	if err != nil {
		log.Printf("Error saving payment: %v", err)
		return nil, fmt.Errorf("database error")
	}

	// Publish payment event to RabbitMQ
	paymentUpdate := &pb.PaymentUpdateRequest{
		TransactionId: transactionID, // Use the actual transaction ID
		Status:        "PENDING",     // Initially set to PENDING
	}

	// Marshal the payment update into a byte slice (for logging or processing if needed)
	body, err := proto.Marshal(paymentUpdate)
	if err != nil {
		log.Printf("Error marshaling payment event: %v", err)
		return nil, fmt.Errorf("internal error")
	}

	// Log the marshaled data for debugging (optional)
	log.Printf("Marshalled payment update: %x", body)

	// Publish the original paymentUpdate object to RabbitMQ (not the marshaled body)
	err = h.RabbitConn.Publish("payment_updates", paymentUpdate)
	if err != nil {
		log.Printf("Error publishing payment event: %v", err)
		return nil, fmt.Errorf("internal error")
	}

	// Return response with the transaction ID and payment status
	return &pb.PaymentResponse{
		TransactionId: transactionID, // Include the transaction ID in the response
		Status:        "PENDING",     // Initial status set to PENDING
		Message:       "Payment processing, please check status later.",
	}, nil
}

// GetPaymentStatus retrieves the payment status for a given payment ID
func (h *PaymentHandler) GetPaymentStatus(ctx context.Context, req *pb.PaymentStatusRequest) (*pb.PaymentStatusResponse, error) {
	log.Printf("Fetching payment status for transaction ID: %s", req.TransactionId)

	// Fetch payment status from the database
	status, err := h.DB.GetPaymentStatus(req.TransactionId)
	if err != nil {
		log.Printf("Error fetching payment status: %v", err)
		return nil, fmt.Errorf("failed to fetch payment status")
	}

	// Return the payment status in the response
	return &pb.PaymentStatusResponse{
		TransactionId: req.TransactionId,
		Status:        status,
		Message:       "Payment status retrieved successfully",
	}, nil
}

// UpdatePaymentStatus updates the status of a payment in the database and returns a response
func (h *PaymentHandler) UpdatePaymentStatus(ctx context.Context, req *pb.PaymentUpdateRequest) (*pb.PaymentUpdateResponse, error) {
	log.Printf("Updating payment status for transaction ID: %s to status: %s", req.TransactionId, req.Status)

	// Check if the payment exists in the database before updating
	var existingStatus string
	err := h.DB.QueryRow("SELECT status FROM payments WHERE transaction_id = $1", req.TransactionId).Scan(&existingStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No payment found with transaction ID: %s", req.TransactionId)
		}
		log.Printf("Error fetching payment for transaction ID %s: %v", req.TransactionId, err)
		return nil, fmt.Errorf("failed to fetch payment status for transaction ID %s", req.TransactionId)
	}

	// Update the payment status in the database
	err = h.DB.UpdatePaymentStatus(req.TransactionId, req.Status)
	if err != nil {
		log.Printf("Error updating payment status: %v", err)
		return nil, fmt.Errorf("failed to update payment status")
	}

	// Return the response indicating successful status update
	return &pb.PaymentUpdateResponse{
		TransactionId: req.TransactionId,
		Status:        req.Status,
		Message:       "Payment status updated successfully",
	}, nil
}

// Listen for updates from RabbitMQ and update the payment status in the database
func (h *PaymentHandler) ListenForPaymentStatusUpdates() {
	// Consume messages from the "payment_updates" queue
	msgs, err := h.RabbitConn.Consume("payment_updates")
	if err != nil {
		log.Fatalf("Failed to start consuming messages: %v", err)
	}

	// Process the messages
	for msg := range msgs {
		// Log the raw message for debugging
		log.Printf("Received message: %s", msg.Body)

		// Unmarshal the Protobuf message into a PaymentUpdateRequest
		var statusUpdate pb.PaymentUpdateRequest
		err := proto.Unmarshal(msg.Body, &statusUpdate)
		if err != nil {
			log.Printf("Error unmarshalling payment status update message: %v", err)
			continue
		}

		// Log the unmarshalled data
		log.Printf("Received payment status update: Transaction ID: %s, Status: %s", statusUpdate.TransactionId, statusUpdate.Status)

		// Update the payment status in the database
		err = h.DB.UpdatePaymentStatus(statusUpdate.TransactionId, "SUCCESSFUL")//statusUpdate.Status
		if err != nil {
			log.Printf("Error updating payment status for transaction %s: %v", statusUpdate.TransactionId, err)
		} else {
			log.Printf("Successfully updated payment status for transaction %s: %s", statusUpdate.TransactionId, statusUpdate.Status)
		}
	}
}
