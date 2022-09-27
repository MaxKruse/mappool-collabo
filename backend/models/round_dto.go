package models

import "backend/models/entities"

type RoundDto struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Mappool     []MapDto
	Suggestions []SuggestionDto
}

func RoundDtoFromEntity(round entities.Round) RoundDto {
	return RoundDto{
		ID:          round.ID,
		Name:        round.Name,
		Mappool:     MapDtoListFromEntityList(round.Mappool),
		Suggestions: SuggestionDtoListFromEntityList(round.Suggestions),
	}
}

func RoundDtoListFromEntityList(rounds []entities.Round) []RoundDto {
	var res []RoundDto
	for _, round := range rounds {
		res = append(res, RoundDtoFromEntity(round))
	}
	return res
}
