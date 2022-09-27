package entities

import (
	"gorm.io/gorm"
)

type Suggestion struct {
	gorm.Model
	Author   User
	AuthorID uint
	Map      Map
	MapID    uint
	RoundID  uint
	Comment  string
	Votes    []Vote
}
