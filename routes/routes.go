package routes

import (
	"pycore/controllers"

	"github.com/gorilla/mux"
)

func UserHandleRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/register", controllers.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/v1/login", controllers.LoginHandler).Methods("POST")

	return router
}
