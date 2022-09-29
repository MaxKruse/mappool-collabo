package osuapiservice

import (
	"backend/services/userservice"
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

func NewClient(auth_token string) (*http.Client, error) {
	// get self from auth token
	self, err := userservice.GetUserFromToken(auth_token)
	if err != nil {
		return nil, err
	}

	oauthClient := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken:  self.Token.AccessToken,
		TokenType:    self.Token.TokenType,
		RefreshToken: self.Token.RefreshToken,
		Expiry:       self.Token.Expiry,
	}))

	return oauthClient, nil
}
