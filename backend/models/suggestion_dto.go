package models

import "backend/models/entities"

type SuggestionDto struct {
	ID        uint      `json:"id"`
	Comment   string    `json:"comment,omitempty"`
	Map       MapDto    `json:"map,omitempty"`
	Votes     []VoteDto `json:"votes,omitempty"`
	VoteScore int       `json:"voteScore"`
	Round     *RoundDto `json:"round,omitempty"`
}

func SuggestionDtoFromEntity(suggestion entities.Suggestion) SuggestionDto {
	votes := 0
	for _, vote := range suggestion.Votes {
		votes += vote.Value
	}

	return SuggestionDto{
		ID:        suggestion.ID,
		Comment:   suggestion.Comment,
		Map:       MapDtoFromEntity(suggestion.Map),
		Votes:     VoteDtoListFromEntityList(suggestion.Votes),
		VoteScore: votes,
	}
}

func SuggestionDtoListFromEntityList(suggestions []entities.Suggestion) []SuggestionDto {
	var res []SuggestionDto
	for _, suggestion := range suggestions {
		res = append(res, SuggestionDtoFromEntity(suggestion))
	}
	return res
}
