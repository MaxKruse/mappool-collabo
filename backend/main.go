package main

import (
	"backend/routes/oauth"
	"backend/routes/tournament"
	"backend/routes/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// app
	app := fiber.New(fiber.Config{
		AppName: "Mappool-Collab Backend",
	})

	// middlewares
	app.Use(compress.New(compress.Config{Level: compress.LevelBestCompression}))
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] (${latency}) [${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "2006-01-02T15:04:05.000",
	}))
	app.Use(recover.New())

	// make session

	// Our groups
	usersGroup := app.Group("/users")
	{
		usersGroup.Get("/", user.List)
		usersGroup.Get("/self", user.GetSelf)
		usersGroup.Get("/:id", user.Get)
	}

	oauthGroup := app.Group("/oauth")
	{
		oauthGroup.Get("/login", oauth.Login)
	}

	tournamentGroup := app.Group("/tournaments")
	{
		tournamentGroup.Get("/", tournament.List)
		tournamentGroup.Get("/:id", tournament.Get)
		tournamentGroup.Post("/", tournament.Create)
		tournamentGroup.Put("/:id", tournament.Update)
		tournamentGroup.Delete("/:id", tournament.Delete)
	}

	mappoolerGroup := app.Group("/mappoolers")
	{
		mappoolerGroup.Post("/", tournament.AddMappooler)
		mappoolerGroup.Delete("/", tournament.RemoveMappooler)
	}

	// run the app
	app.Listen(":5000")
}
