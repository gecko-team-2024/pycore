package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pycore/config"
	"pycore/middleware"
	"pycore/routes"
)

func main() {
	middleware.LogEvent("Server is starting...")
	fmt.Println("Logger started")

	// connect database
	config.InitFirebase()
	config.InitOAuth()

	// create router
	router := routes.UserHandleRoutes()

	// middleware logger
	loggerRouter := middleware.LoggingMiddleware(router)

	// config CORS
	corsHandler := config.SetupCORS(loggerRouter)
	fmt.Println("CORS is running")

	// run server
	ip := os.Getenv("SERVER_IP")
	if ip == "" {
		ip = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	fmt.Println("Server is running on http://" + ip + ":" + port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
