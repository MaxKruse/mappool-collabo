package oauth

import (
	"backend/models"
	"backend/models/entities"
	"backend/services/database"
	"backend/services/userservice"
	"backend/util"
	"context"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
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

	oauthToken, err := getOauth(ctx, &oauthConfig, code)

	if err != nil {
		log.Println(err)
		return err
	}

	client := oauthConfig.Client(context.Background(), oauthToken)

	resp, err := client.Get("https://osu.ppy.sh/api/v2/me/")
	if err != nil {
		return err
	}

	converted, err := util.Convert[models.BanchoUserResponse](resp.Body)
	if err != nil {
		return err
	}

	sessionToken := makeSessionToken(converted)

	err = saveUser(converted, sessionToken, oauthToken)
	if err != nil {
		return err
	}

	return ctx.Redirect(util.Config.FrontendURL + "/login?token=" + sessionToken)
}

func saveUser(user models.BanchoUserResponse, sessionToken string, oauthToken *oauth2.Token) error {
	db := database.GetDebugDBSession()

	// get existing user from db, by id
	var existingUser entities.User
	db.First(&existingUser, user.ID)

	// make the session variable in any case
	session := entities.Session{
		UserId:    int(user.ID),
		AuthToken: sessionToken,
	}

	// if user exists, add the session token to the existing user
	if existingUser.ID != 0 {
		existingUser.Sessions = append(existingUser.Sessions, session)
		existingUser.Token = *oauthToken
		db.Save(&existingUser)
	} else {
		// if user does not exist, create a new user and add the session token
		newUser := entities.User{
			Model:     gorm.Model{ID: uint(user.ID)},
			Sessions:  []entities.Session{session},
			AvatarUrl: user.AvatarUrl,
			Username:  user.Username,
			Token:     *oauthToken,
		}
		db.Create(&newUser)
	}

	// find the user for the session we just saved
	savedUser, err := userservice.GetUserFromToken(sessionToken)
	if err != nil {
		return err
	}

	if savedUser.ID == 0 {
		return errors.New("could not save user correctly to database")
	}
	return nil
}

func makeSessionToken(user models.BanchoUserResponse) string {
	// generate a random uuid
	res, err := uuid.NewUUID()
	if err != nil {
		return ""
	}

	// return the uuid as a string
	return res.String()
}

func getOauth(ctx *fiber.Ctx, oauthConfig *oauth2.Config, code string) (*oauth2.Token, error) {
	sess, err := util.GetSession(ctx)
	if err != nil {
		return nil, err
	}
	// Read oauthState from Cookie
	oauth_state := sess.Get("oauth_state")

	if ctx.Query("state") != oauth_state {
		return nil, errors.New("invalid oauth state")
	}

	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	return token, nil
}
