// db/db.go
package db

import (
	"fmt"
	"database/sql"
	"log"

	"github.com/Go-payments/internal/config"
	_ "github.com/go-sql-driver/mysql" // 
)

type DB struct {
	*sql.DB
}

func Connect(cfg config.Config) (*DB, error) {
	// Open a connection to the MySQL database
	conn, err := sql.Open("mysql", cfg.MySQLDSN)
	if err != nil {
		return nil, err // Return the error for proper handling
	}

	// Verify the connection with a ping
	if err := conn.Ping(); err != nil {
		conn.Close() // Close the connection on error
		return nil, err
	}

	log.Println("Connected to MySQL database.")
	return &DB{conn}, nil
}

// SavePayment saves the payment to the database.
func (db *DB) SavePayment(senderId, receiverId string, amount float64) error {
	// Example query to insert payment into database
	query := `INSERT INTO payments (sender_id, receiver_id, amount) VALUES (?, ?, ?)`
	_, err := db.Exec(query, senderId, receiverId, amount)
	if err != nil {
		return fmt.Errorf("failed to save payment: %v", err)
	}
	return nil
}

// GetPaymentStatus retrieves the payment status from the database.
func (db *DB) GetPaymentStatus(paymentId string) (string, error) {
	var status string
	// Example query to fetch payment status from database
	query := `SELECT status FROM payments WHERE payment_id = ?`
	err := db.QueryRow(query, paymentId).Scan(&status)
	if err != nil {
		return "", fmt.Errorf("failed to get payment status: %v", err)
	}
	return status, nil
}