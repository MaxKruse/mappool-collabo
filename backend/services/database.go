package services

import (
	"backend/models/entities"
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		panic(err)
	}

	// migrate tables
	db.AutoMigrate(&entities.Session{})
	db.AutoMigrate(&entities.User{})

}

func GetDBSession() *gorm.DB {
	return db.Session(&gorm.Session{})
}

func GetDebugDBSession() *gorm.DB {
	return db.Debug().Session(&gorm.Session{})
}

func GetUserFromToken(token string) (entities.User, error) {
	dbSession := GetDBSession()

	authToken := entities.Session{}
	err = dbSession.Find(&authToken, "auth_token = ?", token).Error
	if err != nil {
		return entities.User{}, err
	}

	if authToken.ID == 0 {
		return entities.User{}, errors.New("token not found")
	}

	user := entities.User{}
	err = dbSession.Find(&user, "id = ?", authToken.UserId).Error

	return user, err
}

func GetUserFromId(id string) (entities.User, error) {
	dbSession := GetDBSession()

	res := entities.User{}
	dbSession.Find(&res, "id = ?", id)

	return res, nil
}

func GetUsers() []entities.User {
	dbSession := GetDBSession()

	var res []entities.User
	dbSession.Find(&res)

	return res
}
