package entities

import (
	"fmt"

	"gorm.io/gorm"
)

type Map struct {
	gorm.Model
	Name        string
	Link        string
	PlaySlot    Slot
	Description string
	Suggestions []Suggestion
	Round       Round
	RoundId     uint
}

func (m Map) SlotName() string {
	slotIndexStr := fmt.Sprintf("%d", m.PlaySlot.Index)

	return m.PlaySlot.Name + slotIndexStr
}
