package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/Go-payments/internal/db"
	"github.com/Go-payments/internal/rabbitmq"
	"github.com/Go-payments/internal/config"
	grpc_server "github.com/Go-payments/internal/api/grpc"
	"google.golang.org/grpc"
	pb "github.com/Go-payments/internal/proto/grpc"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	dbConn, err := db.Connect(cfg.DBURL) // Pass the DB URL (string) here
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close() // Ensure the database connection is closed when the server stops

	// Connect to RabbitMQ
	rabbitConn, err := rabbitmq.Connect() // You may pass cfg.RabbitMQURL if needed
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitConn.Close() // Close RabbitMQ connection when the server stops

	// Create PaymentHandler
	customDB := &db.DB{DB: dbConn}
	paymentHandler := grpc_server.NewPaymentHandler(customDB, rabbitConn)

	// Set up the gRPC server and listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Create a new gRPC server and register the PaymentService
	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, paymentHandler)
    go paymentHandler.ListenForPaymentStatusUpdates()

	// Start a goroutine for gRPC server to run in the background
	go func() {
		log.Println("Starting gRPC server on :50051...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	// Create and configure Echo HTTP server
	e := echo.New()

	// Define the routes for HTTP requests
	e.POST("/make-payment", func(c echo.Context) error {
		var paymentReq pb.PaymentRequest
		if err := c.Bind(&paymentReq); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Establish gRPC connection
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to connect to gRPC server"})
		}
		defer conn.Close()

		// Create a new gRPC client
		client := pb.NewPaymentServiceClient(conn)

		// Make the gRPC call to Process the payment
		resp, err := client.MakePayment(context.Background(), &paymentReq)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// Return the response
		return c.JSON(http.StatusOK, resp)
	})

	e.POST("/get-payment-status", func(c echo.Context) error {
		var statusReq pb.PaymentStatusRequest
		if err := c.Bind(&statusReq); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Establish gRPC connection
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to connect to gRPC server"})
		}
		defer conn.Close()

		// Create a new gRPC client
		client := pb.NewPaymentServiceClient(conn)

		// Make the gRPC call to get the payment status
		resp, err := client.GetPaymentStatus(context.Background(), &statusReq)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// Return the response
		return c.JSON(http.StatusOK, resp)
	})

	// Start the Echo HTTP server
	log.Println("Starting Echo HTTP server on :8080...")
	e.Logger.Fatal(e.Start(":8080"))
}
