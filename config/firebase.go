package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var Client *firestore.Client

func InitFirebase() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Initial Firebase Error: %v\n", err)
	}

	Client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Can connect Firestore: %v\n", err)
	}

	log.Println("Firebase initialized successfully")
}
