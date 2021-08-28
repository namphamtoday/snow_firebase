package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/namphamtoday/snow_firebase/authen"
	"github.com/namphamtoday/snow_firebase/connect"
	"github.com/sirupsen/logrus"
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

	// AddUser()
	// UpdateUser()
	// DeleteUser()

	// authen.GetUserByEmail("phamhoangnam2608@gmail.com")
	// CreateCustomToken()
	// VerifyToken()

	RevokeRefreshToken()
}

func CreateCustomToken() {
	claims := map[string]interface{}{
		"employee_code": "01038",
	}

	token, err := authen.CreateCustomToken("fjMiBL69BTVGskTbR1LyRo9PaN33", claims)
	if err != nil {
		logrus.Errorf("Error create custom token %v\n", err)
	}
	logrus.Println(token)
}

func VerifyToken() {
	token, err := authen.VerifyToken("eyJhbGciOiJSUzI1NiIsImtpZCI6IjkwMDk1YmM2ZGM2ZDY3NzkxZDdkYTFlZWIxYTU1OWEzZDViMmM0ODYiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiUGjhuqFtIEhvw6BuZyBOYW0iLCJwaWN0dXJlIjoicGhvdG8iLCJlbXBsb3llZV9jb2RlIjoiMDEwMzgiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vc25vd2FwcC02ZjliYyIsImF1ZCI6InNub3dhcHAtNmY5YmMiLCJhdXRoX3RpbWUiOjE2MzAxNTgzMDEsInVzZXJfaWQiOiJmak1pQkw2OUJUVkdza1RiUjFMeVJvOVBhTjMzIiwic3ViIjoiZmpNaUJMNjlCVFZHc2tUYlIxTHlSbzlQYU4zMyIsImlhdCI6MTYzMDE1ODMwMSwiZXhwIjoxNjMwMTYxOTAxLCJlbWFpbCI6InBoYW1ob2FuZ25hbTI2MDhAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJwaG9uZV9udW1iZXIiOiIrODQ5Nzg5NzE2MjQiLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7InBob25lIjpbIis4NDk3ODk3MTYyNCJdLCJlbWFpbCI6WyJwaGFtaG9hbmduYW0yNjA4QGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6ImN1c3RvbSJ9fQ.E8bZ5jyKtAS73_K37lNaLlN3-PKGu1QQMEkz1QncUmsljmA1-9TphQEXGKZtr1rOTZ2xHocJ2XHyqoqHoRmo3haH1Mt7aiqht1CwEGvHrSTi-dfIAVV530c_g43FT9_vyTxuZRlCq881gLnvOi7o1M2R9RxIP6QgiakHlJ7s4M9C_5LgpsqpwnHMlVEcI3UVKBtaCyWxF3P992HguG-WP36i25BS_LGs4rVg-J60CjHfrxzF5Em5O3vZJcy3Kvr09rDJd-XKJ2-Rz537nZJwvuN68cvw1ekz6nLNdIiI1Ywc_LXQNAVztGB2udoAlKMJrc-rb5Qhiye1CHeUWIxO0Q")
	if err != nil {
		logrus.Errorf("Error verify token %v\n", err)
	}
	strToken, _ := json.Marshal(token.Claims)
	logrus.Println(string(strToken))
}

func RevokeRefreshToken() {
	authen.RevokeRefreshToken("fjMiBL69BTVGskTbR1LyRo9PaN33")
}
func AddUser() {
	var user authen.FireBaseUser
	user.Email = "phamhoangnam2608@gmail.com"
	user.PhoneNumber = "+84978971624"
	user.Disabled = false
	user.DisplayName = "Phạm Hoàng Nam"
	user.PhotoURL = "photo"
	user.Password = "123456789"
	authen.CreateUser(user)
}

func UpdateUser() {
	var user authen.FireBaseUser
	user.Email = "phamhoangnam2608@gmail.com"
	user.PhoneNumber = "+84978971625"
	user.Disabled = false
	user.DisplayName = "Phạm Hoàng Nam"
	user.PhotoURL = "photo"
	user.Password = "123456789"
	user.Id = "US13l9cIt1MY9arWGrXdEfxhSVR2"
	authen.UpdateUser(user)
}

func DeleteUser() {
	authen.DeleteUser([]string{"US13l9cIt1MY9arWGrXdEfxhSVR2", "ULtZYpXfJBSc2Mx2k565BA7sbqo1"})
}
