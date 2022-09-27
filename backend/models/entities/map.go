package entities

import (
	"fmt"

	"gorm.io/gorm"
)

type Map struct {
	gorm.Model
	Name        string `json:"name"`
	Slot        Slot   `json:"slot" gorm:"many2many:map_slots;"`
	Description string `json:"description"`
}

func (m Map) SlotName() string {
	slotIndexStr := fmt.Sprintf("%d", m.Slot.Index)

	return m.Slot.Name + slotIndexStr
}
