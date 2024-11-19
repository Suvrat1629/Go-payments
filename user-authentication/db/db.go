package db
import (
	"fmt"
	"log"
	// "net/http"
	// "time"
	// "github.com/golang-jwt/jwt/v4"
	// "github.com/labstack/echo/v4"
	"github.com/jackc/pgx/v5"
	// "strings"
	// "golang.org/x/crypto/bcrypt"
	// token	"user-auth/jwt-tokenization"
	"context"
)

// Connect to PostgreSQL Database
func InitDB()(*pgx.Conn,error)  {
	var err error
	// Replace with your PostgreSQL connection details
	connString := "postgres://postgres:password@localhost:5432/mydb?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	// Ensure the connection is valid
	err = conn.Ping(context.Background())
	if err != nil {
		conn.Close(context.Background()) // Close if ping fails
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	log.Println("Connected to PostgreSQL successfully")
	return conn, nil
}

// Create the users table if it does not exist
func CreateTableIfNotExists(db *pgx.Conn) error {
	// SQL to create the 'users' table if it does not exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	);
	`

	// Execute the query to create the table
	_, err := db.Exec(context.Background(), createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}
	log.Println("Checked for users table and created if it did not exist")
	return nil
}
