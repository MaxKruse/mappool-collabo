package services

import (
	"backend/models/entities"

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
