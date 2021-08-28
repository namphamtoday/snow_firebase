package authen

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/namphamtoday/snow_firebase/connect"
	"github.com/sirupsen/logrus"
)

func GetUserByUID(uid string) (*auth.UserRecord, error) {
	u, err := connect.FBClient.GetUser(context.Background(), uid)
	if err != nil {
		logrus.Errorf("error getting user %s: %v\n", uid, err)
		return nil, err
	}
	return u, err
}

func GetUserByEmail(email string) (*auth.UserRecord, error) {
	u, err := connect.FBClient.GetUserByEmail(context.Background(), email)
	if err != nil {
		logrus.Errorf("error getting user by email %s: %v\n", email, err)
		return nil, err
	}
	return u, err
}

func GetUserByPhone(phone string) (*auth.UserRecord, error) {
	u, err := connect.FBClient.GetUserByPhoneNumber(context.Background(), phone)
	if err != nil {
		logrus.Errorf("error getting user by phone %s: %v\n", phone, err)
		return nil, err
	}
	return u, err
}

func GetListUser() {

}

func CreateUser(user FireBaseUser) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(user.EmailVerified).
		PhoneNumber(user.PhoneNumber).
		Password(user.Password).
		DisplayName(user.DisplayName).
		PhotoURL(user.PhotoURL).
		Disabled(user.Disabled)
	u, err := connect.FBClient.CreateUser(context.Background(), params)
	if err != nil {
		logrus.Errorf("error creating user: %v\n", err)
		return nil, err
	}
	return u, nil
}

func UpdateUser(user FireBaseUser) (*auth.UserRecord, error) {
	params := (&auth.UserToUpdate{}).
		Email(user.Email).
		EmailVerified(user.EmailVerified).
		PhoneNumber(user.PhoneNumber).
		Password(user.Password).
		DisplayName(user.DisplayName).
		PhotoURL(user.PhotoURL).
		Disabled(user.Disabled)
	u, err := connect.FBClient.UpdateUser(context.Background(), user.Id, params)
	if err != nil {
		logrus.Errorf("error updating user: %v\n", err)
		return nil, err
	}
	return u, nil
}

func DeleteUser(arrUID []string) (*auth.DeleteUsersResult, error) {
	deleteUsersResult, err := connect.FBClient.DeleteUsers(context.Background(), arrUID)
	if err != nil {
		logrus.Errorf("error deleting users: %v\n", err)
		return nil, err
	}
	for _, err := range deleteUsersResult.Errors {
		logrus.Errorf("%v", err)
	}
	return deleteUsersResult, nil
}
