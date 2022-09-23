package entities

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserId    int
	AuthToken string
}

type User struct {
	gorm.Model
	Sessions  []Session
	AvatarUrl string
	Username  string
}
