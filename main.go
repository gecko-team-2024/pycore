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
	//connect database
	config.InitFirebase()

	//start router
	router := routes.UserHandleRoutes()

	//start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server is running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
