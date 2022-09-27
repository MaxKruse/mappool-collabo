package models

import "backend/models/entities"

type SuggestionDto struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Votes       int    `json:"votes"`
}

func SuggestionDtoFromEntity(suggestion entities.Suggestion) SuggestionDto {
	votes := 0
	for _, vote := range suggestion.Votes {
		votes += vote.Value
	}

	return SuggestionDto{
		ID:          suggestion.ID,
		Name:        suggestion.Map.Name,
		Description: suggestion.Map.Description,
		Votes:       votes,
	}
}

func SuggestionDtoListFromEntityList(suggestions []entities.Suggestion) []SuggestionDto {
	var res []SuggestionDto
	for _, suggestion := range suggestions {
		res = append(res, SuggestionDtoFromEntity(suggestion))
	}
	return res
}
