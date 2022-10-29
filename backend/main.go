package main

import (
	"backend/routes"
	"backend/services/sessioncleanerservice"

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
	routes.AddRoutes(app)

	// start the session cleaner service
	sessCleaner := sessioncleanerservice.NewService(sessioncleanerservice.Config{
		Margin: 60,
	})
	sessCleaner.Run()

	// run the app
	app.Listen(":5000")
}
