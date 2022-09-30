package tournamentservice

import (
	"backend/models"
	"backend/models/entities"
	"backend/services/database"
	"backend/services/userservice"
	"errors"
)

func AddRound(token string, roundDto models.RoundDto) (entities.Round, error) {
	// get the user from the token
	user, err := userservice.GetUserFromToken(token)
	if err != nil {
		return entities.Round{}, err
	}

	// get the tournament from the database
	tournament, err := GetTournament(roundDto.TournamentID, DepthRounds|DepthPoolers)
	if err != nil {
		return entities.Round{}, err
	}

	// check if the user is a mappooler or the owner of the tournament
	canEdit := false
	for _, pooler := range tournament.Poolers {
		if pooler.ID == user.ID {
			canEdit = true
			break
		}
	}
	if tournament.Owner.ID == user.ID {
		canEdit = true
	}

	// if the user is not a mappooler or the owner, return an error
	if !canEdit {
		return entities.Round{}, errors.New("you are not allowed to edit this tournament")
	}

	// create the round
	round := entities.Round{
		Name:         roundDto.Name,
		TournamentId: roundDto.TournamentID,
	}

	// add the round to the database

	dbSession := database.GetDBSession()
	err = dbSession.Create(&round).Error

	if err != nil {
		return entities.Round{}, err
	}

	return round, nil
}
