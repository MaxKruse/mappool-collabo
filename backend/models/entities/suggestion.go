package entities

import (
	"gorm.io/gorm"
)

type Suggestion struct {
	gorm.Model
	Round    Round
	RoundId  uint `gorm:"index:round_suggestion,unique"`
	Author   User
	AuthorId uint
	Map      Map
	MapId    uint `gorm:"index:round_suggestion,unique"`
	Comment  string
	Votes    []Vote
}
