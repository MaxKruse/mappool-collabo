package tournamentservice

import (
	"backend/models"
	"backend/models/entities"
	"backend/services/database"
	"backend/services/osuapiservice"
	"backend/services/userservice"
	"backend/util/modenum"
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
		return entities.Round{}, errors.New("round with this name already exists for this tournament")
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
		return errors.New("could not delete the round: " + err.Error())
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

	if len(suggestionDto.Map.Slot.Name) != 2 {
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
			Creator:     beatmapRes.Beatmapset.Creator,
			Description: suggestionDto.Map.Description,
			Difficulty: entities.DifficultyAttributes{
				AR:      beatmapAttribsRes.Attributes.ApproachRate,
				OD:      beatmapAttribsRes.Attributes.OverallDifficulty,
				CS:      convertCS(beatmapRes.Cs, modInts),
				HP:      convertDrain(beatmapRes.Drain, modInts),
				Length:  float64(beatmapRes.TotalLength),
				Stars:   beatmapAttribsRes.Attributes.StarRating,
				ModInts: modInts,
				BPM:     beatmapRes.Bpm,
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
		return entities.Suggestion{}, errors.New("could not create the suggestion: " + err.Error())
	}

	// add the suggestion to the round
	round.Suggestions = append(round.Suggestions, suggestion)
	// save the round
	err = dbSession.Save(&round).Error

	if err != nil {
		return entities.Suggestion{}, errors.New("could not save the round: " + err.Error())
	}

	return suggestion, nil
}

func AddVote(token string, suggestionId uint, vote models.VoteDto) error {
	user, err := userservice.GetUserFromToken(token)
	if err != nil {
		return err
	}

	// get the round from the database
	suggestion, err := GetSuggestion(suggestionId)
	if err != nil {
		return err
	}

	round, err := GetRound(suggestion.RoundId)
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
		if round.ID == suggestion.RoundId {
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

	// if we already voted, return an error
	for _, v2 := range suggestionToUse.Votes {
		if v2.Author.ID == user.ID {
			return errors.New("you already voted for this suggestion")
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
		return errors.New("could not create the vote: " + err.Error())
	}

	return nil
}

func RemoveVote(token string, voteId string) error {
	user, err := userservice.GetUserFromToken(token)
	if err != nil {
		return err
	}

	// get the vote from the database
	vote, err := GetVote(voteId)
	if err != nil {
		return err
	}

	// check if the user is trying to remove his own vote
	if vote.Author.ID != user.ID {
		return errors.New("you are not allowed to remove this vote")
	}

	// delete the vote from the database
	dbSession := database.GetDBSession()
	err = dbSession.Delete(&vote).Error

	if err != nil {
		return errors.New("could not delete the vote: " + err.Error())
	}

	return nil
}

func GetMappool(roundId string, format string) ([]entities.Map, error) {
	var mappool []entities.Map

	// we only support json and csv
	if format != "json" && format != "csv" {
		return mappool, errors.New("invalid format")
	}

	// get the round from the database
	round, err := GetRoundDeep(roundId)
	if err != nil {
		return mappool, err
	}

	// check if the round has a mappool
	if len(round.Mappool) == 0 {
		return mappool, errors.New("this round does not have a mappool yet")
	}

	return round.Mappool, nil
}
