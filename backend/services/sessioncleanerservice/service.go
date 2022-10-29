package sessioncleanerservice

import (
	"backend/models/entities"
	"backend/services/database"
	"log"
	"time"

	"gorm.io/gorm"
)

type Config struct {
	// The amount of time left before a session expires
	Margin int
}

type Service struct {
	config    Config
	dbSession *gorm.DB
}

const defaultMargin = 10

func NewService(configs ...Config) *Service {
	var config Config
	if len(configs) > 0 {
		config = configs[0]
	}

	if config.Margin == 0 {
		config.Margin = defaultMargin
	}

	dbSession := database.GetDBSession()

	return &Service{
		dbSession: dbSession,
		config:    config,
	}
}

func (s *Service) Run() {
	// run an infinite loop that runs every (config.margin / 2) minutes
	go s.internalRun()
}

func (s *Service) internalRun() {
	// on first start, clean
	s.Clean()
	for {
		<-time.After(time.Duration(s.config.Margin/2) * time.Minute)
		s.Clean()
	}
}

func (s *Service) Clean() {
	// get all users
	var users []entities.User
	s.dbSession.Preload("Token").Preload("Sessions").Find(&users)

	// for each user, check if their oauth token is expired.
	for _, user := range users {
		// if the user has no sessions, skip
		if len(user.Sessions) == 0 {
			continue
		}

		// if the user has no token, skip
		if user.Token.AccessToken == "" {
			continue
		}

		// if the token is going to expire in the next (margin) minutes, delete it and all sessions
		expiry := user.Token.Expiry
		lastViableTime := time.Now().Add(time.Duration(s.config.Margin) * time.Minute)
		if expiry.Before(lastViableTime) {
			s.dbSession.Delete(&user.Token)
			s.dbSession.Delete(&user.Sessions)
			log.Printf("Deleted token and sessions for user (%d) %s", user.ID, user.Username)
		}
	}
}
