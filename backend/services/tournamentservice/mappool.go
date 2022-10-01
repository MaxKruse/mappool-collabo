package tournamentservice

import (
	"backend/models"
	"backend/models/entities"
	"backend/services/database"
	"backend/services/osuapiservice"
	"backend/services/userservice"
	"errors"
	"strconv"

	"gorm.io/gorm"
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

func AddSuggestion(token string, suggestionsDto models.SuggestionDto, roundId string) (entities.Suggestion, error) {
	// get the user from the token
	user, err := userservice.GetUserFromToken(token)
	if err != nil {
		return entities.Suggestion{}, err
	}

	// get the round from the database
	round, err := GetRound(roundId, DepthSuggestions)
	if err != nil {
		return entities.Suggestion{}, err
	}

	// get the tournament from the database
	tournament, err := GetTournament(round.TournamentId, DepthBasic|DepthSuggestions)
	if err != nil {
		return entities.Suggestion{}, err
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
		return entities.Suggestion{}, errors.New("you are not allowed to edit this tournament")
	}

	// first, get the beatmap data from the osuapiservice
	// TODO
	_, err = osuapiservice.NewClient(token)
	if err != nil {
		return entities.Suggestion{}, err
	}

	// parse the suggestionsDto.Map.Slot. Format should be "NM1", "HD2", "HR3", "DT4", "FM5" etc.

	if len(suggestionsDto.Map.Slot) < 3 {
		return entities.Suggestion{}, errors.New("invalid slot")
	}

	slotName := suggestionsDto.Map.Slot[:2]
	slotNumber, err := strconv.Atoi(suggestionsDto.Map.Slot[2:])
	if err != nil {
		return entities.Suggestion{}, errors.New("invalid slot index")
	}

	// TODO: Actually fill the beatmap data from the osuapiservice
	mapEntity := entities.Map{
		Model: gorm.Model{
			ID: 696969,
		},
		Name: "todo",
		Link: "todo",
		PlaySlot: entities.Slot{
			Name:  slotName,
			Index: slotNumber,
		},
		Description: suggestionsDto.Comment,
		RoundId:     round.ID,
		Difficulty: entities.DifficultyAttributes{
			HP:      5,
			OD:      5,
			AR:      5,
			CS:      5,
			Stars:   6.9,
			ModInts: 72,
		},
	}

	// create the suggestion
	suggestion := entities.Suggestion{
		Map:    mapEntity,
		Round:  round,
		Author: user,
	}

	// add the suggestion to the database
	dbSession := database.GetDBSession()
	err = dbSession.Create(&suggestion).Error

	if err != nil {
		return entities.Suggestion{}, err
	}

	return suggestion, nil
}

func GetRound(roundId string, depth int) (entities.Round, error) {
	// get the round from the database
	dbSession := database.GetDBSession()
	round := entities.Round{}

	// get the round from the database
	err := dbSession.Where("id = ?", roundId).First(&round).Error
	if err != nil {
		return entities.Round{}, err
	}

	return round, nil
}
