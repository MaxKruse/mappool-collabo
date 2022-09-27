package entities

import (
	"gorm.io/gorm"
)

type Round struct {
	gorm.Model
	Name         string
	TournamentId uint
	Mappool      []Map
	Suggestions  []Suggestion
}
