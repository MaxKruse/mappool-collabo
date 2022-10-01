package database

import (
	"backend/models/entities"
	"backend/util"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func populateDummyData() {
	userEntity := entities.User{
		Username:  "[BH]Lithium",
		AvatarUrl: "https://a.ppy.sh/1199528?1654635999.jpeg",
		Sessions: []entities.Session{
			{
				AuthToken: "e1ae0177-3b7d-11ed-83c0-ceac6f8c603c",
			},
		},
	}
	db.Create(&userEntity)

	tournamentEntity := entities.Tournament{
		Name:        "Test Tournament",
		Owner:       userEntity,
		Description: "This is a test tournament",
		Poolers: []entities.User{
			userEntity,
		},
		Testplayers: []entities.User{
			userEntity,
		},
	}

	db.Create(&tournamentEntity)

	roundEntity := entities.Round{
		TournamentId: tournamentEntity.ID,
		Name:         "Test Round",
		Mappool:      []entities.Map{},
	}

	db.Create(&roundEntity)

	mapEntity := entities.Map{
		Name:        "40mp feat.yuikonnu - Ame to Asphalt",
		Link:        "https://osu.ppy.sh/beatmapsets/709244#osu/1499305",
		Description: "This is a test map",
		RoundId:     roundEntity.ID,
		PlaySlot: entities.Slot{
			Name:  "NM",
			Index: 1,
		},
	}

	db.Create(&mapEntity)

	// add suggestions to the round
	suggestionEntity := entities.Suggestion{
		Author:  userEntity,
		Map:     mapEntity,
		RoundId: roundEntity.ID,
		Comment: "This is a test suggestion",
	}

	db.Create(&suggestionEntity)

	// add votes to the suggestion
	voteEntity := []entities.Vote{
		{
			Author:       userEntity,
			Comment:      "This is a test vote",
			Value:        2,
			SuggestionID: suggestionEntity.ID,
		},
		{
			Author:       userEntity,
			Comment:      "This is another test vote",
			Value:        5,
			SuggestionID: suggestionEntity.ID,
		},
	}

	db.Create(&voteEntity)

}

func isDevelopment() bool {
	// check if we are using in-memory database
	return util.Config.DatabaseURI == ":memory:" || util.Config.DatabaseURI == "development.db"
}

func init() {

	if isDevelopment() {
		// attempt to delete the development database file
		os.Remove(util.Config.DatabaseURI)
	}

	db, err = gorm.Open(sqlite.Open(util.Config.DatabaseURI), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		panic(err)
	}

	// migrate tables
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Session{})
	db.AutoMigrate(&entities.Token{})
	db.AutoMigrate(&entities.Tournament{})
	db.AutoMigrate(&entities.Round{})
	db.AutoMigrate(&entities.Map{})
	db.AutoMigrate(&entities.DifficultyAttributes{})
	db.AutoMigrate(&entities.Suggestion{})
	db.AutoMigrate(&entities.Vote{})
	db.AutoMigrate(&entities.Slot{})

	// check if we are using in-memory database
	if isDevelopment() {
		populateDummyData()
	}

}

func GetDBSession() *gorm.DB {
	return db.Session(&gorm.Session{})
}

func GetDebugDBSession() *gorm.DB {
	return db.Debug().Session(&gorm.Session{})
}
