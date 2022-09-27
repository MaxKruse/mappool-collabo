package tournamentservice

import (
	"backend/models"
	"backend/models/entities"
	"backend/services/database"
	"backend/services/userservice"

	"gorm.io/gorm"
)

// enum for level of depth
const (
	DepthNone        = 1 << iota
	DepthOwner       = 1 << iota
	DepthPoolers     = 1 << iota
	DepthTestplayers = 1 << iota
	DepthRounds      = 1 << iota
	DepthMappool     = 1 << iota
	DepthSuggestions = 1 << iota
	DepthMaps        = 1 << iota
	DepthAll         = DepthOwner | DepthPoolers | DepthTestplayers | DepthRounds | DepthMappool | DepthSuggestions | DepthMaps
	DepthBasic       = DepthOwner | DepthPoolers | DepthTestplayers
)

func preloadFromDepth(db *gorm.DB, depth int) *gorm.DB {
	preloads := db

	if depth&DepthOwner != 0 {
		preloads = preloads.Preload("Owner")
	}

	if depth&DepthPoolers != 0 {
		preloads = preloads.Preload("Poolers")
	}

	if depth&DepthTestplayers != 0 {
		preloads = preloads.Preload("Testplayers")
	}

	if depth&DepthRounds != 0 {
		preloads = preloads.Preload("Rounds")
	}

	if depth&DepthMappool != 0 {
		preloads = preloads.Preload("Rounds.Mappool.PlaySlot")
	}

	if depth&DepthSuggestions != 0 {
		preloads = preloads.Preload("Rounds.Suggestions.Votes.Author")
		preloads = preloads.Preload("Rounds.Suggestions.Map.PlaySlot")
		preloads = preloads.Preload("Rounds.Suggestions.Author")
	}

	if depth&DepthMaps != 0 {
		preloads = preloads.Preload("Rounds.Mappool.PlaySlot")

		preloads = preloads.Preload("Rounds.Suggestions.Votes.Author")
		preloads = preloads.Preload("Rounds.Suggestions.Map.PlaySlot")
		preloads = preloads.Preload("Rounds.Suggestions.Author")
	}

	return preloads
}

func GetTournament[k comparable](id k, depth int) (models.TournamentDto, error) {
	dbSession := database.GetDBSession()
	var tournament entities.Tournament

	preloads := preloadFromDepth(dbSession, depth)

	err := preloads.First(&tournament, id).Error
	return models.TournamentDtoFromEntity(tournament), err
}

func GetTournaments() []models.TournamentDto {
	dbSession := database.GetDBSession()
	var tournaments []entities.Tournament
	preloadFromDepth(dbSession, DepthBasic).Find(&tournaments)
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

	res, err := GetTournament(tournament.ID, DepthAll)
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
