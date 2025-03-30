package main

import (
	"fmt"
	"log"
	"net/http"
	"pycore/config"
	"pycore/routes"
)

func main() {
	//connect database
	config.InitFirebase()

	//start router
	router := routes.UserHandleRoutes()

	//start server
	port := "8080"
	fmt.Println("Server is running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
