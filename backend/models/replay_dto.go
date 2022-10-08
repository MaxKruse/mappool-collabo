package models

import "backend/models/entities"

type ReplayDto struct {
	User     UserDto
	Map      MapDto
	Filepath string
}

func ReplayDtoFromEntity(replay entities.Replay) ReplayDto {
	return ReplayDto{
		User:     UserDtoFromEntity(replay.User),
		Map:      MapDtoFromEntity(replay.Map),
		Filepath: replay.Filepath,
	}
}

func ReplayDtoListFromEntityList(replays []entities.Replay) []ReplayDto {
	var res []ReplayDto
	for _, replay := range replays {
		res = append(res, ReplayDtoFromEntity(replay))
	}
	return res
}
