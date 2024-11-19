package token
import (
	"fmt"
	// "log"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	// "github.com/jackc/pgx/v5"
	"strings"
	// "golang.org/x/crypto/bcrypt"
	// "context"
)
var mySigningKey = []byte("107db0dc9d4fd8b8cc0387f816d5e1ecb6c8183ce7fee4825cfe4e8dc99a2c464ea772daf6cbe8f4507b144bd1898861aac5f80923b2c7345fcbcf17078b2c45ad2e7a902b303349846578d1d4793155d8a9757bc91a09578a6e971d8882a1178f9ee9b3c0b32919aac08c8f610de126e117d0e05522a18a23251d2b4f1ac3366bbd68c4050526c8d2c5fd2dccd683ea982642f4ee53ba569482a46cb2320f603aff011acc6dc5f6a0b2222fb0d52346e4454325a712e60e5942b2f441a7dbcfb9f9e440054f271c86ecc060280fbd1262522fe6e920eb859fb1afefbf93451af705c5d70f6285a327a2ce80cd4ac2c7f754c73a04ed1aaf93cb132e4e80b7f5")

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
