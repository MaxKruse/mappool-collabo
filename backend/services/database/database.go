package database

import (
	"backend/models/entities"
	"backend/util"

	"github.com/davecgh/go-spew/spew"

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
		Rounds: []entities.Round{
			{
				Name: "First Round",
				Mappool: []entities.Map{
					{
						Name:        "VINXIS - Sidetracked Day",
						Link:        "https://osu.ppy.sh/beatmapsets/838182#osu/1754777",
						Description: "Stream Marathon",
						PlaySlot: entities.Slot{
							Name:  "NM",
							Index: 2,
						},
					},
				},
				Suggestions: []entities.Suggestion{
					{
						Author: userEntity,
						Map: entities.Map{
							Name: "40mP feat.yuikonnu - Ame to Asphalt [Lonely]",
							Link: "https://osu.ppy.sh/beatmapsets/709244#osu/1499305",
							PlaySlot: entities.Slot{
								Name:  "NM",
								Index: 1,
							},
							Description: "Consistency map",
						},
						Comment: "Maybe this is good",
						Votes: []entities.Vote{
							{
								Author:  userEntity,
								Comment: "Shit map",
								Value:   -2,
							},
							{
								Author:  userEntity,
								Comment: "Love it",
								Value:   3,
							},
							{
								Author:  userEntity,
								Comment: "ok map",
								Value:   1,
							},
						},
					},
				},
			},
		},
	}

	// print tournamentEntity as json
	spew.Dump(tournamentEntity)

	db.Create(&tournamentEntity)
}

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
	db.AutoMigrate(&entities.Slot{})
	db.AutoMigrate(&entities.Tournament{})
	db.AutoMigrate(&entities.Round{})
	db.AutoMigrate(&entities.Map{})
	db.AutoMigrate(&entities.Suggestion{})
	db.AutoMigrate(&entities.Vote{})

	// check if we are using in-memory database
	if util.Config.DatabaseURI == ":memory:" {
		populateDummyData()
	}

}

func GetDBSession() *gorm.DB {
	return db.Session(&gorm.Session{})
}

func GetDebugDBSession() *gorm.DB {
	return db.Debug().Session(&gorm.Session{})
}
