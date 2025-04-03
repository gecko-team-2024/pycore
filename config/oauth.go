package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig *oauth2.Config

func LoadEnv() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No ENV file found")
		}
	}
}

func InitOAuth() {
	redirectURL := "http://localhost:8081/auth/google/callback" // Mặc định cho môi trường phát triển
	if os.Getenv("ENV") == "production" {
		redirectURL = "https://pycore.onrender.com/auth/google/callback" // URL cho môi trường production
	}

	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  redirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

//https://pycore.onrender.com/auth/google/callback
