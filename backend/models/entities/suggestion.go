package entities

import (
	"gorm.io/gorm"
)

type Suggestion struct {
	gorm.Model
	RoundId  uint
	Author   User
	AuthorId uint
	Map      Map
	MapId    uint
	Comment  string
	Votes    []Vote
}
