package osuapiservice

import (
	"backend/services/userservice"

	client "github.com/deissh/osu-go-client"
)

func NewClient(auth_token string) (*client.OsuAPI, error) {
	// get self from auth token
	self, err := userservice.GetUserFromToken(auth_token)
	if err != nil {
		return nil, err
	}

	access_token := self.Token.AccessToken
	refresh_token := self.Token.RefreshToken

	return client.WithAccessToken(access_token, refresh_token), nil
}
