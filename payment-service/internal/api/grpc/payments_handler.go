package grpc

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Go-payments/internal/proto"
	"github.com/Go-payments/internal/db"
	"github.com/Go-payments/internal/rabbitmq"
)

var (
	ErrInvalidAmount = fmt.Errorf("invalid payment amount")
	ErrDatabase      = fmt.Errorf("database error")
	ErrInternal      = fmt.Errorf("internal server error")
)

// PaymentHandler implements the gRPC PaymentServiceServer interface.
type PaymentHandler struct {
	pb.UnimplementedPaymentServiceServer
	DB         *db.DB
	RabbitConn *rabbitmq.Connection
}

// NewPaymentHandler initializes and returns a PaymentHandler.
func NewPaymentHandler(dbConn *db.DB, rabbitConn *rabbitmq.Connection) pb.PaymentServiceServer {
	return &PaymentHandler{
		DB:         dbConn,
		RabbitConn: rabbitConn,
	}
}
// ProcessPayment handles payment processing logic.
func (h *PaymentHandler) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Printf("Processing payment from %s to %s with amount %f", req.SenderId, req.ReceiverId, req.Amount)

	// Validate input
	if req.Amount <= 0 {
		return nil, ErrInvalidAmount
	}

	// Save payment in the database
	err := h.DB.SavePayment(req.SenderId, req.ReceiverId, req.Amount)
	if err != nil {
		log.Printf("Error saving payment: %v", err)
		return nil, ErrDatabase
	}

	// Publish payment event to RabbitMQ
	err = h.RabbitConn.Publish("payments", req)
	if err != nil {
		log.Printf("Error publishing payment event: %v", err)
		return nil, ErrInternal
	}

	// Return success response
	return &pb.PaymentResponse{
		Status:  "SUCCESS",
		Message: "Payment processed successfully",
	}, nil
}

// GetPaymentStatus retrieves the status of a payment.
func (h *PaymentHandler) GetPaymentStatus(ctx context.Context, req *pb.PaymentStatusRequest) (*pb.PaymentStatusResponse, error) {
	log.Printf("Fetching payment status for payment ID %s", req.PaymentId)

	// Query payment status from the database
	status, err := h.DB.GetPaymentStatus(req.PaymentId)
	if err != nil {
		log.Printf("Error fetching payment status: %v", err)
		return nil, ErrDatabase
	}

	return &pb.PaymentStatusResponse{
		Status: status,
	}, nil
}
