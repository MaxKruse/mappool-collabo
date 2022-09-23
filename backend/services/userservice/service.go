package userservice

import (
	"backend/models/entities"
	"backend/services/database"
	"errors"
)

func GetUserFromToken(token string) (entities.User, error) {
	dbSession := database.GetDBSession()

	authToken := entities.Session{}
	err := dbSession.Find(&authToken, "auth_token = ?", token).Error
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

func GetUserFromId[k comparable](id k) (entities.User, error) {
	dbSession := database.GetDBSession()

	res := entities.User{}
	dbSession.Find(&res, "id = ?", id)

	return res, nil
}

func GetUsers() []entities.User {
	dbSession := database.GetDBSession()

	var res []entities.User
	dbSession.Find(&res)

	return res
}