package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secret_key []byte

type contextKey string

const UserKey contextKey = "username"

func init() {
	// Load file .env
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("WARNING: Error loading .env file. Proceeding without it.")
		}
	}

	// Lấy giá trị SECRET_KEY từ biến môi trường
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		panic("SECRET_KEY environment variable is not set")
	}
	secret_key = []byte(key)
	fmt.Println("SECRET_KEY loaded successfully")
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(secret_key)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret_key, nil
	})

	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret_key, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["username"] == nil {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		username := claims["username"].(string)
		ctx := context.WithValue(r.Context(), UserKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func GetUsernameFromContext(r *http.Request) (string, error) {
	username, ok := r.Context().Value(UserKey).(string)
	if !ok {
		return "", errors.New("username not found in context")
	}
	return username, nil
}
