package util

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	// Database
	DatabaseURI string `default:":memory:"`
	// Oauth for bancho
	BanchoOauthClientID     string
	BanchoOauthClientSecret string
	BanchoOauthRedirectURL  string
	// Frontend
	FrontendURL string
}

var Config config

func init() {
	Config.BanchoOauthClientID = os.Getenv("BANCHO_OAUTH_CLIENT_ID")
	Config.BanchoOauthClientSecret = os.Getenv("BANCHO_OAUTH_CLIENT_SECRET")
	Config.BanchoOauthRedirectURL = os.Getenv("BANCHO_OAUTH_REDIRECT_URL")

	// check if the oauth values are set
	if Config.BanchoOauthClientID == "" || Config.BanchoOauthClientSecret == "" || Config.BanchoOauthRedirectURL == "" {
		panic("bancho oauth values are not set")
	}

	Config.FrontendURL = os.Getenv("FRONTEND_URL")
}
