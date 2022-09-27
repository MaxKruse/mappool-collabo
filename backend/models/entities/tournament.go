package entities

import "gorm.io/gorm"

type Tournament struct {
	gorm.Model
	Name        string
	Description string
	Owner       User `gorm:"foreignKey:OwnerID"`
	OwnerID     uint
	Poolers     []User `gorm:"many2many:tournament_poolers;"`
	Testplayers []User `gorm:"many2many:tournament_testplayers;"`
	Rounds      []Round
}
