package entities

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	UserId    int
	AuthToken string
}
