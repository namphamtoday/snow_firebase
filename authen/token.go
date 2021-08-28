package authen

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/namphamtoday/snow_firebase/connect"
	"github.com/sirupsen/logrus"
)

func CreateCustomToken(uid string, claims map[string]interface{}) (string, error) {
	client, err := connect.FBApp.Auth(context.Background())
	if err != nil {
		logrus.Errorf("[Create custom token] Error init client: %v\n", err)
		return "", err
	}

	token, err := client.CustomTokenWithClaims(context.Background(), uid, claims)
	if err != nil {
		logrus.Errorf("[Create custom token] Error create custom token: %v\n", err)
		return "", err
	}

	return token, nil
}

func VerifyToken(idToken string) (*auth.Token, error) {
	client, err := connect.FBApp.Auth(context.Background())
	if err != nil {
		logrus.Errorf("[Verify token] Error init client: %v\n", err)
		return nil, err
	}

	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		logrus.Errorf("Error verify token: %v\n", err)
		return nil, err
	}

	return token, err
}

func RevokeRefreshToken(uid string) error {
	client, err := connect.FBApp.Auth(context.Background())
	if err != nil {
		logrus.Errorf("error getting Auth client: %v\n", err)
		return err
	}
	err = client.RevokeRefreshTokens(context.Background(), uid)
	if err != nil {
		logrus.Errorf("error revoking tokens for user: %v, %v\n", uid, err)
		return err
	}
	// accessing the user's TokenValidAfter
	u, err := client.GetUser(context.Background(), uid)
	if err != nil {
		logrus.Errorf("error getting user %s: %v\n", uid, err)
		return err
	}
	timestamp := u.TokensValidAfterMillis / 1000
	logrus.Printf("the refresh tokens were revoked at: %d (UTC seconds) ", timestamp)
	return nil
}
