package entities

import (
	"fmt"

	"gorm.io/gorm"
)

type Map struct {
	gorm.Model
	BeatmapId   uint `gorm:"index:idx_map_mods,unique"`
	ModInts     int  `gorm:"index:idx_map_mods,unique"`
	Name        string
	Link        string
	PlaySlot    Slot
	Description string
	Difficulty  DifficultyAttributes
}

type DifficultyAttributes struct {
	gorm.Model
	HP      float64
	OD      float64
	AR      float64
	CS      float64
	Stars   float64
	Length  float64
	ModInts int64
	MapId   uint
}

func (m Map) SlotName() string {
	slotIndexStr := fmt.Sprintf("%d", m.PlaySlot.Index)

	return m.PlaySlot.Name + slotIndexStr
}
