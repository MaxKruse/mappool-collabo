package entities

import (
	"gorm.io/gorm"
)

type Slot struct {
	gorm.Model
	Name  string
	Index int
	MapId uint
}
