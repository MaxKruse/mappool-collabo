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
	// Storage
	StorageURI string
}

var Config config

func init() {
	Config.BanchoOauthClientID = os.Getenv("BANCHO_OAUTH_CLIENT_ID")
	Config.BanchoOauthClientSecret = os.Getenv("BANCHO_OAUTH_CLIENT_SECRET")
	Config.BanchoOauthRedirectURL = os.Getenv("BANCHO_OAUTH_REDIRECT_URL")
	Config.DatabaseURI = os.Getenv("DATABASE_URI")

	// make sure the database uri is set. if not, default it to in-memory
	if Config.DatabaseURI == "" {
		Config.DatabaseURI = ":memory:"
	}

	// check if the oauth values are set
	if Config.BanchoOauthClientID == "" || Config.BanchoOauthClientSecret == "" || Config.BanchoOauthRedirectURL == "" {
		panic("bancho oauth values are not set")
	}

	Config.FrontendURL = os.Getenv("FRONTEND_URL")
	if Config.FrontendURL == "" {
		panic("frontend url is not set")
	}

	Config.StorageURI = os.Getenv("STORAGE_URI")
	if Config.StorageURI == "" {
		panic("storage uri is not set")
	}
}
