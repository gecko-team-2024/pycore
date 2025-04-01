package config

import (
	"net/http"

	"github.com/rs/cors"
)

// SetupCORS trả về một HTTP handler với cấu hình CORS
func SetupCORS(router http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Thay "*" bằng danh sách các origin cụ thể nếu cần
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)
}
