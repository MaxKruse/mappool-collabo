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
