package tournamentservice

import (
	"backend/models"
	"backend/models/entities"
	"backend/services/database"
	"backend/services/userservice"
)

func GetTournament[k comparable](id k) (models.TournamentDto, error) {
	dbSession := database.GetDBSession()
	var tournament entities.Tournament
	err := dbSession.Preload("Owner").Preload("Testplayers").Preload("Poolers").First(&tournament, id).Error
	return models.TournamentDtoFromEntity(tournament), err
}

func GetTournaments() []models.TournamentDto {
	dbSession := database.GetDBSession()
	var tournaments []entities.Tournament
	dbSession.Preload("Owner").Preload("Testplayers").Preload("Poolers").Find(&tournaments)
	return models.TournamentDtoListFromEntityList(tournaments)
}

func CreateTournament(tournament models.TournamentDto) (models.TournamentDto, error) {
	dbSession := database.GetDBSession()
	var err error
	var tournamentEntity entities.Tournament

	// Name, Description and Owner are required
	// The owner has to be identified by his ID
	tournamentEntity.Name = tournament.Name
	tournamentEntity.Description = tournament.Description
	tournamentEntity.Owner, err = userservice.GetUserFromId(tournament.Owner.ID)
	if err != nil {
		return models.TournamentDto{}, err
	}

	// Poolers and Testplayers are optional
	// They have to be identified by their IDs
	for _, pooler := range tournament.Poolers {
		newPooler, _ := userservice.GetUserFromId(pooler.ID)
		tournamentEntity.Poolers = append(tournamentEntity.Poolers, newPooler)
	}
	for _, testplayer := range tournament.Testplayers {
		newTestplayer, _ := userservice.GetUserFromId(testplayer.ID)
		tournamentEntity.Testplayers = append(tournamentEntity.Testplayers, newTestplayer)
	}

	dbSession.Create(&tournamentEntity)
	return models.TournamentDtoFromEntity(tournamentEntity), nil
}

func UpdateTournament(tournament models.TournamentDto) (models.TournamentDto, error) {
	// Update if values are not empty
	dbSession := database.GetDBSession()

	var tournamentEntity entities.Tournament
	dbSession.First(&tournamentEntity, tournament.ID)

	if tournament.Name != "" {
		tournamentEntity.Name = tournament.Name
	}
	if tournament.Description != "" {
		tournamentEntity.Description = tournament.Description
	}
	if tournament.Owner.ID != 0 {
		tournamentEntity.Owner, _ = userservice.GetUserFromId(tournament.Owner.ID)
	}

	// Poolers and Testplayers are optional
	// They have to be identified by their IDs
	for _, pooler := range tournament.Poolers {
		newPooler, _ := userservice.GetUserFromId(pooler.ID)
		tournamentEntity.Poolers = append(tournamentEntity.Poolers, newPooler)
	}
	for _, testplayer := range tournament.Testplayers {
		newTestplayer, _ := userservice.GetUserFromId(testplayer.ID)
		tournamentEntity.Testplayers = append(tournamentEntity.Testplayers, newTestplayer)
	}

	dbSession.Save(&tournamentEntity)

	res, err := GetTournament(tournament.ID)
	if err != nil {
		return models.TournamentDto{}, err
	}

	return res, nil
}

func DeleteTournament[k comparable](id k) error {
	dbSession := database.GetDBSession()
	var tournament entities.Tournament
	err := dbSession.Delete(&tournament, id).Error
	return err
}
