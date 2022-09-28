package tournamentservice

import (
	"backend/models"
	"backend/services/database"
	"backend/services/userservice"
	"errors"
)

func AddMappooler(auth_token string, tournamentID uint, userID uint) error {
	self, err := userservice.GetUserFromToken(auth_token)
	if err != nil {
		return err
	}

	tourney, err := GetTournament(tournamentID, DepthPoolers|DepthBasic)
	if err != nil {
		return err
	}

	// if the user already is a pooler, return pre-emptively
	for _, pooler := range tourney.Poolers {
		if pooler.ID == userID {
			return errors.New("user is already a pooler")
		}
	}

	// check if we are the owner and can actually add poolers
	if self.ID != tourney.Owner.ID {
		return errors.New("you are not the owner of this tournament")
	}

	userToAdd, err := userservice.GetUserFromId(userID)
	if err != nil {
		return err
	}

	tourney.Poolers = append(tourney.Poolers, userToAdd)

	tourneyDto := models.TournamentDto{}
	tourneyDto.ID = tourney.ID
	tourneyDto.Poolers = models.UserDtoListFromEntityList(tourney.Poolers)

	_, err = UpdateTournament(tourneyDto)
	return err
}

func RemoveMappooler(auth_token string, tournamentID uint, userID uint) error {
	self, err := userservice.GetUserFromToken(auth_token)
	if err != nil {
		return err
	}

	tourney, err := GetTournament(tournamentID, DepthPoolers|DepthBasic)
	if err != nil {
		return err
	}

	// check if we are the owner and can actually add poolers
	if self.ID != tourney.Owner.ID {
		return errors.New("you are not the owner of this tournament")
	}

	// find the user in the poolers
	poolerIndex := -1
	for i, pooler := range tourney.Poolers {
		if pooler.ID == userID {
			poolerIndex = i
			break
		}
	}

	// if the user is not a pooler, return pre-emptively
	if poolerIndex == -1 {
		return errors.New("user is not a pooler")
	}

	// remove the pooler
	pooler := tourney.Poolers[poolerIndex]

	dbSession := database.GetDBSession()
	err = dbSession.Model(&tourney).Association("Poolers").Delete(&pooler)
	return err
}
