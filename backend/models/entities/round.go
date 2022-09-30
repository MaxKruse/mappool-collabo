package entities

import (
	"gorm.io/gorm"
)

type Round struct {
	gorm.Model
	Name         string `gorm:"index:test,unique"`
	TournamentId uint   `gorm:"index:test,unique"`
	Mappool      []Map
	Suggestions  []Suggestion
}
