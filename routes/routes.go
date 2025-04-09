package routes

import (
	"pycore/controllers"
	"pycore/middleware"

	"github.com/gorilla/mux"
)

func UserHandleRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/auth/google/callback", controllers.GoogleCallbackHandler).Methods("GET")
	router.HandleFunc("/ws/logs", controllers.StreamLogsHandler)

	router.HandleFunc("/api/v1/register", controllers.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/v1/login", controllers.LoginHandler).Methods("POST")
	router.HandleFunc("/api/v1/user", controllers.GetUserByIDHandler).Methods("GET")
	router.HandleFunc("/auth/google", controllers.GoogleLoginHandler).Methods("GET")
	router.HandleFunc("/api/v1/download", controllers.DownloadFileHandler).Methods("GET")
	router.HandleFunc("/api/v1/update-password", controllers.UpdatePasswordHandler).Methods("POST")
	router.HandleFunc("/api/v1/reward", controllers.ClaimRewardHandler).Methods("POST")
	router.HandleFunc("/api/v2/download", middleware.AuthMiddleware(controllers.DownloadFileHandlerV2)).Methods("GET")

	return router
}
