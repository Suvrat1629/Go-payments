package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/jackc/pgx/v5"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"context"
)

var mySigningKey = []byte("your-very-secure-secret-key")

// Struct for representing user data (for JWT payload)
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Custom claims for JWT token
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var db *pgx.Conn

// Helper function to generate JWT token
func GenerateJWT(username string) (string, error) {
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Helper function to validate JWT token
func ValidateJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// Middleware to validate JWT token for protected routes
func TokenValidationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
		}

		if len(tokenString) > 7 && strings.ToUpper(tokenString[:7]) == "BEARER " {
			tokenString = tokenString[7:]
		}

		claims, err := ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		c.Set("username", claims.Username)
		return next(c)
	}
}

// Protected route handler
func ProtectedHandler(c echo.Context) error {
	username := c.Get("username").(string)
	return c.String(http.StatusOK, fmt.Sprintf("Hello, %s! You have accessed a protected route.", username))
}

// Connect to PostgreSQL Database
func initDB() error {
	var err error
	// Replace with your PostgreSQL connection details
	connString := "postgres://postgres:password@localhost:5432/mydb?sslmode=disable"
	db, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	// Ensure the connection is valid
	err = db.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("unable to ping database: %v", err)
	}
	log.Println("Connected to PostgreSQL successfully")
	return nil
}

// Create the users table if it does not exist
func CreateTableIfNotExists() error {
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
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Check if the username already exists
	var existingUser User
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
	token, err := GenerateJWT(user.Username)
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
	var user User
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
	token, err := GenerateJWT(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create token"})
	}

	// Return JWT to the client
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func main() {
	// Initialize database
	if err := initDB(); err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	// Ensure that the 'users' table exists
	if err := CreateTableIfNotExists(); err != nil {
		log.Fatal(err)
	}

	// Create a new Echo instance
	e := echo.New()

	// Define routes
	e.POST("/login", LoginHandler)

e.POST("/signup", SignupHandler)
	e.GET("/protected", ProtectedHandler, TokenValidationMiddleware)

	// Start the server
	log.Println("Server running on http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
