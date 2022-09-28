package entities

import (
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserId    int
	AuthToken string
}

type Token struct {
	gorm.Model
	oauth2.Token
	UserId uint
}

type User struct {
	gorm.Model
	Sessions  []Session
	AvatarUrl string
	Username  string
	Token     Token
}
