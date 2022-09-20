package util

import "github.com/maxkruse/flagorenv"

type config struct {
	// Database
	DatabaseURI string `default:":memory:"`
	// Oauth for bancho
	BanchoOauthClientID     string
	BanchoOauthClientSecret string
	BanchoOauthRedirectURL  string
}

var Config config

func init() {
	c, err := flagorenv.LoadFlagsOrEnv[config](&flagorenv.Config{
		Prefix:     "COLLAB",
		PreferFlag: true,
	})

	if err != nil {
		panic(err)
	}

	// check if the oauth values are set
	if c.BanchoOauthClientID == "" || c.BanchoOauthClientSecret == "" || c.BanchoOauthRedirectURL == "" {
		panic("bancho oauth values are not set")
	}

	Config = c
}
