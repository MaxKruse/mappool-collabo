package models

import "backend/models/entities"

type TournamentDto struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Owner       UserDto   `json:"owner"`
	Poolers     []UserDto `json:"poolers"`
	Testplayers []UserDto `json:"testplayers"`
}

func TournamentDtoFromEntity(tournament entities.Tournament) TournamentDto {
	return TournamentDto{
		ID:          tournament.ID,
		Name:        tournament.Name,
		Description: tournament.Description,
		Owner:       UserDtoFromEntity(tournament.Owner),
		Poolers:     UserDtoListFromEntityList(tournament.Poolers),
		Testplayers: UserDtoListFromEntityList(tournament.Testplayers),
	}
}

func TournamentDtoListFromEntityList(tournaments []entities.Tournament) []TournamentDto {
	var res []TournamentDto
	for _, tournament := range tournaments {
		res = append(res, TournamentDtoFromEntity(tournament))
	}
	return res
}
