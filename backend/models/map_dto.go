package models

import "backend/models/entities"

type MapDto struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Slot        string `json:"slot"`
	Description string `json:"description"`
}

func MapDtoFromEntity(mapEntity entities.Map) MapDto {
	return MapDto{
		ID:          mapEntity.ID,
		Name:        mapEntity.Name,
		Slot:        mapEntity.SlotName(),
		Description: mapEntity.Description,
	}
}

func MapDtoListFromEntityList(mapEntities []entities.Map) []MapDto {
	var res []MapDto
	for _, mapEntity := range mapEntities {
		res = append(res, MapDtoFromEntity(mapEntity))
	}
	return res
}
