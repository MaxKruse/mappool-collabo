package beatmapservice

import (
	"backend/models/entities"
	"backend/services/database"
	"backend/services/tournamentservice"
	"backend/services/userservice"
	"errors"
	"strconv"
)

func GetBeatmap[k comparable](auth_token string, beatmapId k) (entities.Map, error) {
	user, err := userservice.GetUserFromToken(auth_token)
	if err != nil {
		return entities.Map{}, err
	}

	// get all tournaments. If the user is a testplayer, pooler or the owner, check if the beatmap is in the tournament
	// if it is, return the beatmap
	// if it isn't, return an error

	tournaments := tournamentservice.GetTournaments()

	canView := false

	for _, tournament := range tournaments {
		if tournament.Owner.ID == user.ID {
			canView = true
			break
		}

		for _, pooler := range tournament.Poolers {
			if pooler.ID == user.ID {
				canView = true
				break
			}
		}

		for _, testplayer := range tournament.Testplayers {
			if testplayer.ID == user.ID {
				canView = true
				break
			}
		}
	}

	if !canView {
		return entities.Map{}, errors.New("user is not allowed to view this beatmap")
	}

	dbSession := database.GetDBSession()
	var beatmap entities.Map
	err = dbSession.Where("id = ?", beatmapId).First(&beatmap).Error
	if err != nil {
		return entities.Map{}, errors.New("beatmap not found")
	}

	return beatmap, nil
}

func AddReplay(auth_token string, beatmapId string, replayPath string) error {
	user, err := userservice.GetUserFromToken(auth_token)
	if err != nil {
		return err
	}

	// get all tournaments. If the user is a testplayer, pooler or the owner, check if the beatmap is in the tournament
	// if it is, return the beatmap
	// if it isn't, return an error

	tournaments := tournamentservice.GetTournaments()

	canView := false

	for _, tournament := range tournaments {
		if tournament.Owner.ID == user.ID {
			canView = true
			break
		}

		for _, pooler := range tournament.Poolers {
			if pooler.ID == user.ID {
				canView = true
				break
			}
		}

		for _, testplayer := range tournament.Testplayers {
			if testplayer.ID == user.ID {
				canView = true
				break
			}
		}
	}

	if !canView {
		return errors.New("user is not allowed to add replays to this beatmap")
	}

	dbSession := database.GetDBSession()
	var beatmap entities.Map
	err = dbSession.Where("id = ?", beatmapId).First(&beatmap).Error
	if err != nil {
		return errors.New("beatmap not found")
	}

	// convert beatmapId to uint
	beatmapIdUint, err := strconv.ParseUint(beatmapId, 10, 64)

	replay := entities.Replay{
		MapId:    uint(beatmapIdUint),
		Filepath: replayPath,
	}

	err = dbSession.Create(&replay).Error
	if err != nil {
		return err
	}

	return nil
}

func GetReplay(auth_token string, replayId string) (entities.Replay, error) {
	user, err := userservice.GetUserFromToken(auth_token)
	if err != nil {
		return entities.Replay{}, err
	}

	dbSession := database.GetDBSession()
	var replay entities.Replay
	err = dbSession.Preload("User").Preload("Map").Where("id = ?", replayId).First(&replay).Error
	if err != nil {
		return entities.Replay{}, errors.New("replay not found")
	}

	// get all tournaments. If the user is a testplayer, pooler or the owner, check if the beatmap is in the tournament
	// if it is, return the beatmap
	// if it isn't, return an error

	tournaments := tournamentservice.GetTournaments()

	canView := false

	for _, tournament := range tournaments {
		if tournament.Owner.ID == user.ID {
			canView = true
			break
		}

		for _, pooler := range tournament.Poolers {
			if pooler.ID == user.ID {
				canView = true
				break
			}
		}

		for _, testplayer := range tournament.Testplayers {
			if testplayer.ID == user.ID {
				canView = true
				break
			}
		}
	}

	if !canView {
		return entities.Replay{}, errors.New("user is not allowed to view this replay")
	}

	return replay, nil
}
