package userservice

import (
	"backend/models/entities"
	"backend/services/database"
	"errors"
	"fmt"

	"github.com/Code-Hex/go-generics-cache/policy/lru"
)

var userCache *lru.Cache[string, entities.User]

func init() {
	userCache = lru.NewCache[string, entities.User](lru.WithCapacity(100))
}

func GetUserFromToken(token string) (entities.User, error) {
	dbSession := database.GetDBSession()

	// ensure the token has the format of
	// "Bearer <token>"
	if len(token) < 7 || token[:7] != "Bearer " {
		return entities.User{}, errors.New("invalid token format")
	}

	// truncate the "Bearer " part
	token = token[7:]

	authToken := entities.Session{}
	err := dbSession.Find(&authToken, "auth_token = ?", token).Error
	if err != nil {
		return entities.User{}, err
	}

	if authToken.ID == 0 {
		return entities.User{}, errors.New("token not found")
	}

	return GetUserFromId(authToken.UserId)
}

func GetUserFromId[k comparable](id k) (entities.User, error) {

	if user, ok := userCache.Get(fmt.Sprintf("%v", id)); ok {
		return user, nil
	}

	dbSession := database.GetDBSession()

	res := entities.User{}
	err := dbSession.Find(&res, "id = ?", id).Error

	if err != nil {
		return entities.User{}, err
	}

	userCache.Set(fmt.Sprintf("%v", id), res)

	return res, err
}

func GetUsers() []entities.User {
	dbSession := database.GetDBSession()

	var res []entities.User
	dbSession.Find(&res)

	return res
}
