package main

import (
	"log"
	"net"

	"github.com/Go-payments/internal/config"
	"github.com/Go-payments/internal/db"

	// "github.com/Go-payments/internal/api/grpc"
	paymentgrpc "github.com/Go-payments/internal/api/grpc"
	middlewares "github.com/Go-payments/internal/api/middleware" // Assuming your middleware is in this package
	pb "github.com/Go-payments/internal/proto"
	"github.com/Go-payments/internal/rabbitmq"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc" // <-- Add this import
)

func main() {
	// Load configurations
	cfg := config.Load()

	// Initialize Database
	dbConn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close() // Close the database connection when main exits

	// Initialize RabbitMQ
	rabbitConn, err := rabbitmq.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitConn.Close() // Close the RabbitMQ connection when main exits

	// Initialize Echo server (optional if you want HTTP alongside gRPC)
	e := echo.New()

	// Register middlewares if you want to use them for HTTP requests
	e.Use(middlewares.AuthMiddleware)

	// Start gRPC Server
	listener, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", cfg.GRPCPort, err)
	}
	grpcServer := grpc.NewServer() // Correctly calling the grpc.NewServer() function

	// Register gRPC services
	pb.RegisterPaymentServiceServer(grpcServer, paymentgrpc.NewPaymentHandler(dbConn, rabbitConn))

	// Run gRPC server in a goroutine
	go func() {
		log.Printf("gRPC server running on port %s", cfg.GRPCPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// If you're also running HTTP server using Echo, you can start it here
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "Service is running")
	})

	// Start HTTP server (optional)
	go func() {
		if err := e.Start(cfg.HTTPPort); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Block the main goroutine to keep both servers running
	select {}
}
