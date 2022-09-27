package entities

import "gorm.io/gorm"

type Round struct {
	gorm.Model
	Name         string
	TournamentId uint
	Mappool      []Map        `gorm:"many2many:round_mappools;"`
	Suggestions  []Suggestion `gorm:"many2many:round_suggestions;"`
}
