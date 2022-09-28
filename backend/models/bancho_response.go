package models

type BanchoUserResponse struct {
	ID        uint   `json:"id"`
	AvatarUrl string `json:"avatar_url"`
	Username  string `json:"username"`
}
