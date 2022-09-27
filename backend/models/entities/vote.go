package entities

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	Author       User
	AuthorID     uint
	Suggestion   Suggestion
	SuggestionID uint
	Comment      string
	Value        int
}
