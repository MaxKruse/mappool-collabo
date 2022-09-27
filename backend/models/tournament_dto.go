package models

import "backend/models/entities"

type TournamentDto struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Owner       UserDto    `json:"owner"`
	Poolers     []UserDto  `json:"poolers,omitempty"`
	Testplayers []UserDto  `json:"testplayers,omitempty"`
	Rounds      []RoundDto `json:"rounds,omitempty"`
}

func TournamentDtoFromEntity(tournament entities.Tournament) TournamentDto {
	return TournamentDto{
		ID:          tournament.ID,
		Name:        tournament.Name,
		Description: tournament.Description,
		Owner:       UserDtoFromEntity(tournament.Owner),
		Poolers:     UserDtoListFromEntityList(tournament.Poolers),
		Testplayers: UserDtoListFromEntityList(tournament.Testplayers),
		Rounds:      RoundDtoListFromEntityList(tournament.Rounds),
	}
}

func TournamentDtoListFromEntityList(tournaments []entities.Tournament) []TournamentDto {
	var res []TournamentDto
	for _, tournament := range tournaments {
		res = append(res, TournamentDtoFromEntity(tournament))
	}
	return res
}

func (t TournamentDto) IsPooler(user entities.User) bool {
	for _, pooler := range t.Poolers {
		if pooler.ID == user.ID {
			return true
		}
	}
	return false
}

func (t TournamentDto) IsTestplayer(user entities.User) bool {
	for _, testplayer := range t.Testplayers {
		if testplayer.ID == user.ID {
			return true
		}
	}
	return false
}
