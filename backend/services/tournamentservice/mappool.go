package tournamentservice

import (
	"backend/models"
	"backend/models/entities"
	"backend/services/database"
	"backend/services/osuapiservice"
	"backend/services/userservice"
	"backend/util/modenum"
	"errors"
	"math"
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
	apiclient, err := osuapiservice.NewClient(token)
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

	modStrings := []string{slotName}

	beatmapRes, err := apiclient.GetBeatmap(suggestionsDto.Map.ID)
	if err != nil {
		return entities.Suggestion{}, err
	}
	beatmapAttribsRes, err := apiclient.GetBeatmapWithMods(suggestionsDto.Map.ID, "osu", modenum.ModStringsToInt64(modStrings))
	if err != nil {
		return entities.Suggestion{}, err
	}

	// format the name of the map
	// Artist - Song [Difficulty]
	beatmapName := beatmapRes.Beatmapset.Artist + " - " + beatmapRes.Beatmapset.Title + " [" + beatmapRes.Version + "]"

	mapEntity := entities.Map{
		Model: gorm.Model{
			ID: uint(beatmapRes.ID),
		},
		Name: beatmapName,
		Link: beatmapRes.URL,
		PlaySlot: entities.Slot{
			Name:  slotName,
			Index: slotNumber,
		},
		Description: suggestionsDto.Comment,
		RoundId:     round.ID,
		Difficulty: entities.DifficultyAttributes{
			AR:         beatmapAttribsRes.Attributes.ApproachRate,
			OD:         beatmapAttribsRes.Attributes.OverallDifficulty,
			CS:         convertCS(beatmapRes.Cs, modenum.ModStringsToInt64(modStrings)),
			HP:         convertDrain(beatmapRes.Drain, modenum.ModStringsToInt64(modStrings)),
			Length:     float64(beatmapRes.TotalLength),
			Stars:      beatmapAttribsRes.Attributes.StarRating,
			ModStrings: modStrings,
			ModInts:    modenum.ModStringsToInt64(modStrings),
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

func convertDrain(drain float64, mods int64) float64 {
	// easy halves it
	if mods&modenum.Easy != 0 && drain > 0 {
		drain /= 2
	}

	// hardrock multiplies by 1.4, with 10.0 clamp
	if mods&modenum.HardRock != 0 {
		drain *= 1.4
		if drain > 10.0 {
			drain = 10.0
		}
	}

	// Doubletime increases it by 50% (artifically, still want to display this)
	if mods&modenum.DoubleTime != 0 {
		drain *= 1.5
	}

	// Halftime decreases it by 25%
	if mods&modenum.HalfTime != 0 {
		drain *= 0.75
	}

	return drain
}

func convertCS(cs float64, modInts int64) float64 {
	if modInts&modenum.HardRock > 0 {
		// Multiply by 1.3, but clamp to 10.0
		cs = math.Min(cs*1.3, 10.0)
	}

	// prevent divide by zero
	if modInts&modenum.Easy > 0 && cs > 0 {
		// Divide by 2
		cs = cs / 2
	}

	return cs
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
