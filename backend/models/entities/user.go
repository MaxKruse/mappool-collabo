package entities

import (
	"backend/models"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string    `json:"username"`
	Sessions     []Session `json:"-"`
	BanchoUser   models.User
	BanchoUserId uint
}
