package models

type BanchoUserResponse struct {
	ID        uint   `json:"id"`
	AvatarUrl string `json:"avatar_url"`
	Username  string `json:"username"`
}

func (self BanchoUserResponse) ToUserDto() UserDto {
	return UserDto{
		BanchoUserResponse: self,
	}
}
