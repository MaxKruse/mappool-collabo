package osuapiservice

import (
	"backend/services/userservice"
	"context"

	gosuapiclient "github.com/maxkruse/gosu-api-client"

	"golang.org/x/oauth2"
)

func NewClient(auth_token string) (*gosuapiclient.Client, error) {
	// get self from auth token
	self, err := userservice.GetUserFromToken(auth_token)
	if err != nil {
		return nil, err
	}

	oauthClient := gosuapiclient.NewClient(oauth2.Token{
		AccessToken:  self.Token.AccessToken,
		TokenType:    self.Token.TokenType,
		RefreshToken: self.Token.RefreshToken,
		Expiry:       self.Token.Expiry,
	}, context.Background())

	return oauthClient, nil
}
