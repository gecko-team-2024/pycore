package routes

import (
	"pycore/controllers"

	"github.com/gorilla/mux"
)

func UserHandleRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/register", controllers.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/v1/login", controllers.LoginHandler).Methods("POST")
	router.HandleFunc("/auth/google", controllers.GoogleLoginHandler).Methods("GET")
	router.HandleFunc("/auth/google/callback", controllers.GoogleCallbackHandler).Methods("GET")

	return router
}
