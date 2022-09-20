package oauth

import (
	"backend/models"
	"backend/util"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

var (
	oauthConfig = oauth2.Config{
		ClientID:     util.Config.BanchoOauthClientID,
		ClientSecret: util.Config.BanchoOauthClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://osu.ppy.sh/oauth/authorize",
			TokenURL:  "https://osu.ppy.sh/oauth/token",
			AuthStyle: oauth2.AuthStyleAutoDetect,
		},
		RedirectURL: util.Config.BanchoOauthRedirectURL,
		Scopes:      []string{""},
	}
)

func Login(ctx *fiber.Ctx) error {
	sess, err := util.GetSession(ctx)
	if err != nil {
		return err
	}

	code := ctx.Query("code")
	if code == "" {
		sess.Regenerate()
		id := sess.ID()
		sess.Set("oauth_state", id)
		// Save session
		if err := sess.Save(); err != nil {
			return err
		}

		return ctx.Redirect(oauthConfig.AuthCodeURL(id))
	}

	token, err := getOauth(ctx, &oauthConfig, code)

	if err != nil {
		log.Println(err)
		return err
	}

	client := oauthConfig.Client(context.Background(), token)

	resp, err := client.Get("https://osu.ppy.sh/api/v2/me/")
	if err != nil {
		return err
	}

	converted := util.Convert[models.BanchoUserResponse](resp.Body)

	return ctx.JSON(converted.ToUser())
}

func getOauth(ctx *fiber.Ctx, oauthConfig *oauth2.Config, code string) (*oauth2.Token, error) {
	sess, err := util.GetSession(ctx)
	if err != nil {
		return nil, err
	}
	// Read oauthState from Cookie
	oauth_state := sess.Get("oauth_state")

	if ctx.Query("state") != oauth_state {
		return nil, fiber.ErrBadRequest
	}

	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	return token, nil
}
