package models

import (
	"backend/models/entities"
	"backend/util/modenum"
)

type SlotDto struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
}

type MapDto struct {
	ID          uint    `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Slot        SlotDto `json:"slot,omitempty"`
	Description string  `json:"description,omitempty"`
}

type DifficultyAttributeDto struct {
	HP         float64  `json:"hp,omitempty"`
	OD         float64  `json:"od,omitempty"`
	AR         float64  `json:"ar,omitempty"`
	CS         float64  `json:"cs,omitempty"`
	Stars      float64  `json:"stars,omitempty"`
	ModStrings []string `json:"modStrings,omitempty"`
	ModInts    int64    `json:"modInts,omitempty"`
}

func MapDtoFromEntity(mapEntity entities.Map) MapDto {
	return MapDto{
		ID:          mapEntity.ID,
		Name:        mapEntity.Name,
		Slot:        SlotDtoFromEntity(mapEntity.PlaySlot),
		Description: mapEntity.Description,
	}
}

func SlotDtoFromEntity(slotEntity entities.Slot) SlotDto {
	return SlotDto{
		Name:  slotEntity.Name,
		Index: slotEntity.Index,
	}
}

func MapDtoListFromEntityList(mapEntities []entities.Map) []MapDto {
	var res []MapDto
	for _, mapEntity := range mapEntities {
		res = append(res, MapDtoFromEntity(mapEntity))
	}
	return res
}

func DifficultyAttributeDtoFromEntity(difficultyAttributeEntity entities.DifficultyAttributes) DifficultyAttributeDto {
	var modStrings []string
	var modInts int64

	if difficultyAttributeEntity.ModInts != 0 {
		modInts = difficultyAttributeEntity.ModInts
		modStrings = modenum.ModIntsToStringArray(modInts)
	}

	return DifficultyAttributeDto{
		HP:         difficultyAttributeEntity.HP,
		OD:         difficultyAttributeEntity.OD,
		AR:         difficultyAttributeEntity.AR,
		CS:         difficultyAttributeEntity.CS,
		Stars:      difficultyAttributeEntity.Stars,
		ModStrings: modStrings,
		ModInts:    modInts,
	}
}

func DifficultyAttributeDtoListFromEntityList(difficultyAttributeEntities []entities.DifficultyAttributes) []DifficultyAttributeDto {
	var res []DifficultyAttributeDto
	for _, difficultyAttributeEntity := range difficultyAttributeEntities {
		res = append(res, DifficultyAttributeDtoFromEntity(difficultyAttributeEntity))
	}
	return res
}
