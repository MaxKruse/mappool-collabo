package entities

import "gorm.io/gorm"

type Replay struct {
	gorm.Model
	UserId   uint
	User     User
	MapId    uint
	Map      Map
	Filepath string
}
