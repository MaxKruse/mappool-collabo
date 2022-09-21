package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Sessions  []Session
	AvatarUrl string
	Username  string
}
