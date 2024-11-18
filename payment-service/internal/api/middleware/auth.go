package middlewares

import (
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	// Error definitions
	ErrUnauthorized = errors.New("unauthorized")

	// Secret key (should be loaded from environment variable for production use)
	jwtSecretKey = []byte("your-secret-key")
)

// AuthMiddleware is an Echo-compatible middleware that validates incoming requests for authentication.
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, ErrUnauthorized.Error())
		}

		// Validate the token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if !isValidToken(token) {
			return echo.NewHTTPError(http.StatusUnauthorized, ErrUnauthorized.Error())
		}

		// Extract user ID from the token
		userID := extractUserID(token)
		if userID == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, ErrUnauthorized.Error())
		}

		// Set user information in the context
		c.Set("user_id", userID)

		// Call the next handler in the chain
		return next(c)
	}
}

// isValidToken validates the provided JWT token.
func isValidToken(tokenString string) bool {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	// Return whether the token is valid
	return err == nil && token.Valid
}

// extractUserID extracts the user ID from a valid JWT token.
func extractUserID(tokenString string) string {
	// Parse the token with claims
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil || !token.Valid {
		return ""
	}

	// Extract and return the user ID
	if claims, ok := token.Claims.(*jwt.MapClaims); ok {
		if userID, found := (*claims)["user_id"]; found {
			return fmt.Sprintf("%v", userID)
		}
	}

	return ""
}
