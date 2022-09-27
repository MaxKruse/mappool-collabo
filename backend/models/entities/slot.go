package entities

import (
	"gorm.io/gorm"
)

// A slot can be used by many maps, and many maps can have this specific slot.
type Slot struct {
	gorm.Model
	Name  string `json:"name"`
	Index int    `json:"index"`
}
