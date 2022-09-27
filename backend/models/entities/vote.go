package entities

import (
	"gorm.io/gorm"
)

type Vote struct {
	gorm.Model
	Author       User
	AuthorID     uint
	SuggestionID uint
	Comment      string
	Value        int
}
