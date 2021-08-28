package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/namphamtoday/snow_firebase/connect"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	InitFireBase()
}

func InitFireBase() {
	filePath := os.Getenv("FIREBASE_CREDENTIAL_FILE_PATH")
	connect.ConnectFireBase(filePath)
	connect.InitAuthen()
	connect.InitFireStore()
}
