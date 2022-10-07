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

func RemoveRound(token string, roundId string) error {
	// get the user from the token
	user, err := userservice.GetUserFromToken(token)
	if err != nil {
		return err
	}

	// get the round from the database
	round, err := GetRound(roundId)
	if err != nil {
		return err
	}

	// get the tournament of this round
	tournament, err := GetTournament(round.TournamentId, DepthRounds|DepthPoolers)
	if err != nil {
		return err
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
		return errors.New("you are not allowed to edit this tournament")
	}

	// remove the round from the database
	dbSession := database.GetDBSession()
	err = dbSession.Delete(&round).Error

	if err != nil {
		return err
	}

	return nil
}

func AddSuggestion(token string, suggestionDto models.SuggestionDto, roundId string) (entities.Suggestion, error) {
	// get the user from the token
	user, err := userservice.GetUserFromToken(token)
	if err != nil {
		return entities.Suggestion{}, err
	}

	// get the round from the database
	round, err := GetRound(roundId)
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

	// parse the suggestionDto.Map.Slot. Format should be "NM1", "HD2", "HR3", "DT4", "FM5" etc.

	if len(suggestionDto.Map.Slot.Name) < 2 {
		return entities.Suggestion{}, errors.New("invalid slot name")
	}

	if suggestionDto.Map.Slot.Index == 0 {
		return entities.Suggestion{}, errors.New("invalid slot index")
	}

	modStrings := []string{suggestionDto.Map.Slot.Name}
	modInts := modenum.ModStringsToInt64(modStrings)

	// if we already have a map with this ID and modInts, return, dont create a new entity
	dbSession := database.GetDBSession()
	var mapEntity entities.Map

	err = dbSession.Preload("Difficulty").Preload("PlaySlot").Where("beatmap_id = ? AND mod_ints = ?", suggestionDto.Map.ID, modInts).First(&mapEntity).Error

	if err != nil {
		beatmapRes, err := apiclient.GetBeatmap(suggestionDto.Map.ID)
		if err != nil {
			return entities.Suggestion{}, err
		}
		beatmapAttribsRes, err := apiclient.GetBeatmapWithMods(suggestionDto.Map.ID, "osu", modInts)
		if err != nil {
			return entities.Suggestion{}, err
		}

		// format the name of the map
		// Artist - Song [Difficulty]
		beatmapName := beatmapRes.Beatmapset.Artist + " - " + beatmapRes.Beatmapset.Title + " [" + beatmapRes.Version + "]"

		mapEntity = entities.Map{
			BeatmapId: uint(beatmapRes.ID),
			Name:      beatmapName,
			Link:      beatmapRes.URL,
			PlaySlot: entities.Slot{
				Name:  suggestionDto.Map.Slot.Name,
				Index: suggestionDto.Map.Slot.Index,
			},
			Description: suggestionDto.Map.Description,
			Difficulty: entities.DifficultyAttributes{
				AR:      beatmapAttribsRes.Attributes.ApproachRate,
				OD:      beatmapAttribsRes.Attributes.OverallDifficulty,
				CS:      convertCS(beatmapRes.Cs, modInts),
				HP:      convertDrain(beatmapRes.Drain, modInts),
				Length:  float64(beatmapRes.TotalLength),
				Stars:   beatmapAttribsRes.Attributes.StarRating,
				ModInts: modInts,
			},
		}

	}

	// create the suggestion
	suggestion := entities.Suggestion{
		Map:     mapEntity,
		Round:   round,
		Author:  user,
		Comment: suggestionDto.Comment,
	}

	// add the suggestion to the database
	err = dbSession.Create(&suggestion).Error

	if err != nil {
		return entities.Suggestion{}, err
	}

	// add the suggestion to the round
	round.Suggestions = append(round.Suggestions, suggestion)
	// save the round
	err = dbSession.Save(&round).Error

	if err != nil {
		return entities.Suggestion{}, err
	}

	return suggestion, nil
}

func AddVote(token string, roundId uint, suggestionId uint, vote models.VoteDto) error {
	user, err := userservice.GetUserFromToken(token)
	if err != nil {
		return err
	}

	// get the round from the database
	round, err := GetRound(roundId)
	if err != nil {
		return err
	}

	// get the tournament from the database
	tournament, err := GetTournament(round.TournamentId, DepthBasic|DepthSuggestions)
	if err != nil {
		return err
	}

	// check if the user is a mappooler or the owner of the tournament
	canEdit := false
	for _, pooler := range tournament.Poolers {
		if pooler.ID == user.ID {
			canEdit = true
			break
		}
	}
	for _, testplayer := range tournament.Testplayers {
		if testplayer.ID == user.ID {
			canEdit = true
			break
		}
	}

	if tournament.Owner.ID == user.ID {
		canEdit = true
	}

	// if the user is not a mappooler or the owner, return an error
	if !canEdit {
		return errors.New("you are not allowed to vote for this tournament")
	}

	var roundToUse entities.Round
	var suggestionToUse entities.Suggestion

	// find the round and suggestion in the tournament
	for _, round := range tournament.Rounds {
		if round.ID == roundId {
			roundToUse = round
			for _, suggestion := range round.Suggestions {
				if suggestion.ID == suggestionId {
					suggestionToUse = suggestion
					break
				}
			}
			break
		}
	}

	// if the round or suggestion could not be found, return an error
	if roundToUse.ID == 0 || suggestionToUse.ID == 0 {
		return errors.New("could not find the round or suggestion")
	}

	// construct our vote entity
	var voteEntity entities.Vote

	voteEntity.Comment = vote.Comment
	voteEntity.Value = vote.Value
	voteEntity.Author = user
	voteEntity.SuggestionID = suggestionToUse.ID

	// add the vote to the database
	dbSession := database.GetDBSession()
	err = dbSession.Create(&voteEntity).Error

	if err != nil {
		return err
	}

	return err
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

func GetRound[k comparable](roundId k) (entities.Round, error) {
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
