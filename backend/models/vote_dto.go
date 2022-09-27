package models

import "backend/models/entities"

type VoteDto struct {
	ID      uint    `json:"id"`
	Value   int     `json:"value"`
	Comment string  `json:"comment"`
	Author  UserDto `json:"author"`
}

func VoteDtoFromEntity(vote entities.Vote) VoteDto {
	return VoteDto{
		ID:      vote.ID,
		Value:   vote.Value,
		Comment: vote.Comment,
		Author:  UserDtoFromEntity(vote.Author),
	}
}

func VoteDtoListFromEntityList(votes []entities.Vote) []VoteDto {
	var res []VoteDto
	for _, vote := range votes {
		res = append(res, VoteDtoFromEntity(vote))
	}
	return res
}
