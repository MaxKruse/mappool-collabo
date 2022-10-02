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
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Slot        SlotDto `json:"slot"`
	Description string  `json:"description"`
}

type DifficultyAttributeDto struct {
	HP         float64  `json:"hp"`
	OD         float64  `json:"od"`
	AR         float64  `json:"ar"`
	CS         float64  `json:"cs"`
	Stars      float64  `json:"stars"`
	ModStrings []string `json:"modStrings"`
	ModInts    int64    `json:"modInts"`
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
	} else if len(difficultyAttributeEntity.ModStrings) != 0 {
		modStrings = difficultyAttributeEntity.ModStrings
		modInts = modenum.ModStringsToInt64(modStrings)
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
