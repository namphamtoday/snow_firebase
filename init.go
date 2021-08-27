package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var FBClient *auth.Client
var FireStoreClient *firestore.Client

func ConnectFireBase() {
	filePath := os.Getenv("FIREBASE_CREDENTIAL_FILE_PATH")
	if filePath == "" {
		log.Fatalf("Credential file not found")
	}

	opt := option.WithCredentialsFile(filePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatalf("Cannot initialize firebase client %v\n", err)
	}

	// Auth client
	authClient, err := app.Auth(context.Background())

	if err != nil {
		log.Fatalf("Cannot getting Auth client %v\n", err)
		panic(err)
	}

	FBClient = authClient

	// Firestore client
	fireStoreClient, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Cannot initialize firestore client %v\n", err)
		panic(err)
	}
	FireStoreClient = fireStoreClient

	log.Println("Firebase client has been created")
}
