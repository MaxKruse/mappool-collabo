package database

import (
	"backend/models/entities"
	"backend/util"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open(sqlite.Open(util.Config.DatabaseURI), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		panic(err)
	}

	// migrate tables
	db.AutoMigrate(&entities.Session{})
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Tournament{})
	db.AutoMigrate(&entities.Round{})
	db.AutoMigrate(&entities.Map{})
	db.AutoMigrate(&entities.Suggestion{})
	db.AutoMigrate(&entities.Vote{})

}

func GetDBSession() *gorm.DB {
	return db.Session(&gorm.Session{})
}

func GetDebugDBSession() *gorm.DB {
	return db.Debug().Session(&gorm.Session{})
}
