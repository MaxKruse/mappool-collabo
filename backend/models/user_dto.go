package models

import "backend/models/entities"

type UserDto struct {
	BanchoUserResponse
}

func UserDtoFromEntity(user entities.User) UserDto {
	return UserDto{
		BanchoUserResponse: BanchoUserResponse{
			ID:        user.ID,
			Username:  user.Username,
			AvatarUrl: user.AvatarUrl},
	}
}

func UserDtoListFromEntityList(users []entities.User) []UserDto {
	var res []UserDto
	for _, user := range users {
		res = append(res, UserDtoFromEntity(user))
	}
	return res
}
