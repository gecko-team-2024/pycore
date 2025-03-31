package config

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var Client *firestore.Client

func InitFirebase() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v\n", err)
		}
	}

	ctx := context.Background()

	// Lấy đường dẫn file credentials từ biến môi trường
	credFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credFile == "" {
		log.Fatalf("GOOGLE_APPLICATION_CREDENTIALS is not set in the environment")
	}

	// Sử dụng file credentials từ biến môi trường
	opt := option.WithCredentialsFile(credFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Initial Firebase Error: %v\n", err)
	}

	Client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Cannot connect to Firestore: %v\n", err)
	}

	log.Println("Firebase initialized successfully")
}
