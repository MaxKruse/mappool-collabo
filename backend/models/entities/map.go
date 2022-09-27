package entities

import (
	"fmt"

	"gorm.io/gorm"
)

type Map struct {
	gorm.Model
	Name        string `json:"name"`
	Link        string `json:"link"`
	PlaySlot    Slot   `json:"slot" gorm:"many2many:map_slots;"`
	Description string `json:"description"`
}

func (m Map) SlotName() string {
	slotIndexStr := fmt.Sprintf("%d", m.PlaySlot.Index)

	return m.PlaySlot.Name + slotIndexStr
}
