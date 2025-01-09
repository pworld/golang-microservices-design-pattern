package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// User struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginUser generates JWT tokens
func LoginUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	// Debug Log: Print environment variable
	if os.Getenv("RUNNING_IN_DOCKER") == "" {
		// Load .env file only in local environment
		err := godotenv.Load("local.env")
		if err != nil {
			log.Fatal("Error loading local.env file")
		}
	} else {
		fmt.Println("Running inside Docker, skipping .env file loading")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Println("JWT_SECRET not set inside user-service")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "JWT_SECRET not set"})
	} else {
		log.Printf("JWT_SECRET inside user-service: %s", jwtSecret)
	}

	// Hardcoded user credentials (Replace with DB lookup)
	if user.Username != "john" || user.Password != "1234" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign token with JWT_SECRET
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

// GetProfile is a protected route
func GetProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "User profile data"})
}
