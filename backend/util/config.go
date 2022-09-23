package util

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	// Database
	DatabaseURI string
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
	Config.DatabaseURI = os.Getenv("DATABASE_URI")

	// make sure the database uri is set. if not, default it to "development.db"
	if Config.DatabaseURI == "" {
		Config.DatabaseURI = "development.db"
	}

	// check if the oauth values are set
	if Config.BanchoOauthClientID == "" || Config.BanchoOauthClientSecret == "" || Config.BanchoOauthRedirectURL == "" {
		panic("bancho oauth values are not set")
	}

	Config.FrontendURL = os.Getenv("FRONTEND_URL")
}
