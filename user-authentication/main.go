package main

import (
	// "fmt"
	"log"
	"net/http"
	// "time"
	// "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
"github.com/labstack/echo/v4/middleware"
	"github.com/jackc/pgx/v5"
	// "strings"
	"golang.org/x/crypto/bcrypt"
	token	"user-auth/jwt-tokenization"
	database "user-auth/db"
	"context"
)


var db *pgx.Conn

// Register a new user
func RegisterUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
	return err
}

// Authenticate user from the database
func AuthenticateUser(username, password string) (bool, error) {
	var storedHash string
	err := db.QueryRow(context.Background(), "SELECT password FROM users WHERE username = $1", username).Scan(&storedHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil // User not found
		}
		return false, err // Database error
	}

	// Compare hashed passwords
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		return false, nil // Invalid password
	}

	return true, nil
}

// Signup handler to register a new user
func SignupHandler(c echo.Context) error {
	var user token.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Check if the username already exists
	var existingUser token.User
	err := db.QueryRow(context.Background(), "SELECT username FROM users WHERE username = $1", user.Username).Scan(&existingUser.Username)
	if err == nil {
		// Username already exists
		return c.JSON(http.StatusConflict, map[string]string{"error": "Username already exists"})
	} else if err != pgx.ErrNoRows {
		// Database error
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	// Hash the user's password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error hashing password"})
	}

	// Insert the new user into the database
	_, err = db.Exec(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, hashedPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create user"})
	}

	// Optionally, generate a JWT token for the user
	token, err := token.GenerateJWT(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
	}

	// Return a success message with the generated token
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User created successfully",
		"token":   token,
	})
}

// Login handler to authenticate the user and return a JWT
func LoginHandler(c echo.Context) error {
	var user token.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Authenticate user from the database
	valid, err := AuthenticateUser(user.Username, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	if !valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}

	// User authenticated, generate JWT
	token, err := token.GenerateJWT(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create token"})
	}

	// Return JWT to the client
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func main() {
// Initialize database
	dbConn, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close(context.Background())
	db=dbConn

	// Ensure that the 'users' table exists
	if err := database.CreateTableIfNotExists(dbConn); err != nil {
		log.Fatal(err)
	}

	// Create a new Echo instance
	e := echo.New()

e.Use(middleware.Logger())
	// Define routes
	e.POST("/login", LoginHandler)
	e.POST("/signup", SignupHandler)
	
e.GET("/protected", token.TokenValidationMiddleware(token.ProtectedHandler))


	// Start the server
	log.Println("Server running on http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
