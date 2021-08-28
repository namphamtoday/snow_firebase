package connect

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var FBApp *firebase.App
var FBClient *auth.Client
var FireStoreClient *firestore.Client
