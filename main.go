package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pycore/config"
	"pycore/routes"
)

func main() {
	// Kết nối database
	config.InitFirebase()
	config.InitOAuth()

	// Khởi tạo router
	router := routes.UserHandleRoutes()

	// Cấu hình CORS
	corsHandler := config.SetupCORS(router)
	fmt.Println("CORS is running")

	// Khởi chạy server
	ip := os.Getenv("SERVER_IP")
	if ip == "" {
		ip = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server is running on http://" + ip + ":" + port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
