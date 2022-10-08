package tournamentservice

import (
	"backend/models/entities"
	"backend/services/database"
	"backend/util/modenum"
	"errors"
	"math"
)

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
		return entities.Round{}, errors.New("could not find the round: " + err.Error())
	}

	return round, nil
}

func GetSuggestion[k comparable](suggestionId k) (entities.Suggestion, error) {
	// get the round from the database
	dbSession := database.GetDBSession()
	suggestion := entities.Suggestion{}

	// get the round from the database
	err := dbSession.Where("id = ?", suggestionId).First(&suggestion).Error
	if err != nil {
		return entities.Suggestion{}, errors.New("could not find the suggestion: " + err.Error())
	}

	return suggestion, nil
}

func GetVote[k comparable](voteId k) (entities.Vote, error) {
	// get the round from the database
	dbSession := database.GetDBSession()
	vote := entities.Vote{}

	// get the round from the database
	err := dbSession.Where("id = ?", voteId).First(&vote).Error
	if err != nil {
		return entities.Vote{}, errors.New("could not find the vote: " + err.Error())
	}

	return vote, nil
}
