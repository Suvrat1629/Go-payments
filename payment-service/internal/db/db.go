package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type DB struct {
	*sql.DB
}

// Connect initializes a new DB connection
func Connect(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Check if the connection is valid
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	return db, nil
}

func (d *DB) SavePayment(senderId, receiverId string, amount float64, transactionID string) error {
	// Prepare SQL statement to save payment
	query := `
		INSERT INTO payments (transaction_id, sender_id, receiver_id, amount, status)
		VALUES ($1, $2, $3, $4, $5)`

	// Execute the insert query with values
	_, err := d.Exec(query, transactionID, senderId, receiverId, amount, "PENDING")
	if err != nil {
		return fmt.Errorf("failed to save payment: %v", err)
	}
	return nil
}

// GetPaymentStatus retrieves the payment status for a given payment ID
func (db *DB) GetPaymentStatus(paymentID string) (string, error) {
	// Query to get the payment status
	var status string
	err := db.QueryRow("SELECT status FROM payments WHERE transaction_id = $1", paymentID).Scan(&status)
	if err != nil {
		return "", fmt.Errorf("failed to fetch payment status: %v", err)
	}

	return status, nil
}

// UpdatePaymentStatus updates the status of a payment in the database
func (db *DB) UpdatePaymentStatus(transactionID string, status string) error {
	// Update the payment status based on the provided transaction ID
	query := `UPDATE payments SET status = $1 WHERE transaction_id = $2`

	// Execute the update query
	_, err := db.Exec(query, status, transactionID)
	if err != nil {
		log.Printf("Error updating payment status for transaction %s: %v", transactionID, err)
		return fmt.Errorf("failed to update payment status")
	}

	// If successful, return nil (no error)
	log.Printf("Payment status updated to %s for transaction %s", status, transactionID)
	return nil
}

// GetAllPayments fetches all transactions from the database
func (db *DB) GetAllPayments() ([]map[string]interface{}, error) {
	// SQL query to fetch all payments
	query := "SELECT transaction_id, sender_id, receiver_id, amount, status, created_at FROM payments"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch payments: %v", err)
	}
	defer rows.Close()

	var payments []map[string]interface{}

	for rows.Next() {
		// Create a map to hold the row data
		var transactionID, senderID, receiverID, status, createdAt string
		var amount float64

		// Scan the current row into variables
		err := rows.Scan(&transactionID, &senderID, &receiverID, &amount, &status, &createdAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		// Create a map for the row and append to the slice
		payment := map[string]interface{}{
			"transaction_id": transactionID,
			"sender_id":      senderID,
			"receiver_id":    receiverID,
			"amount":         amount,
			"status":         status,
			"created_at":     createdAt,
		}
		payments = append(payments, payment)
	}

	// Check for any error encountered while iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during row iteration: %v", err)
	}

	return payments, nil
}
