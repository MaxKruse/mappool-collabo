package entities

import (
	"gorm.io/gorm"
)

type Round struct {
	gorm.Model
	Name         string `gorm:"index:round_tournament,unique"`
	TournamentId uint   `gorm:"index:round_tournament,unique"`
	Mappool      []Map
	Suggestions  []Suggestion
}
