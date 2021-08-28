package connect

import (
	"context"

	log "github.com/sirupsen/logrus"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func ConnectFireBase(filePath string) error {

	if filePath == "" {
		log.Error("Credential file not found")
	}

	var err error
	opt := option.WithCredentialsFile(filePath)
	FBApp, err = firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Errorf("Cannot initialize firebase client %v\n", err)
		return err
	}

	log.Println("Firebase client has been created")
	return nil
}

func InitAuthen() error {

	// Auth client
	authClient, err := FBApp.Auth(context.Background())
	if err != nil {
		log.Errorf("Error initialize Auth client %v\n", err)
		return err
	}

	log.Println("Init Auth Client Success!")
	FBClient = authClient
	return nil
}

func InitFireStore() error {
	// Firestore client
	fireStoreClient, err := FBApp.Firestore(context.Background())
	if err != nil {
		log.Errorf("Error initialize FireStore client %v\n", err)
		return err
	}
	log.Println("Init FireStore Client Success!")
	FireStoreClient = fireStoreClient
	return nil
}
