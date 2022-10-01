package models

import "backend/models/entities"

type SuggestionDto struct {
	ID        uint      `json:"id"`
	Comment   string    `json:"comment"`
	Map       MapDto    `json:"map"`
	Votes     []VoteDto `json:"votes,omitempty"`
	VoteScore int       `json:"voteScore"`
	Round     RoundDto  `json:"round"`
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
		Round:     RoundDtoFromEntity(suggestion.Round),
	}
}

func SuggestionDtoListFromEntityList(suggestions []entities.Suggestion) []SuggestionDto {
	var res []SuggestionDto
	for _, suggestion := range suggestions {
		res = append(res, SuggestionDtoFromEntity(suggestion))
	}
	return res
}
