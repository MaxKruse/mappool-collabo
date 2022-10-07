package routes

import (
	"backend/routes/oauth"
	"backend/routes/tournament"
	"backend/routes/user"

	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
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

	testplayerGroup := app.Group("/testplayers")
	{
		testplayerGroup.Post("/", tournament.AddTestplayer)
		testplayerGroup.Delete("/", tournament.RemoveTestplayer)
	}

	rounds := app.Group("/rounds")
	{
		rounds.Post("/", tournament.CreateRound)
		rounds.Post("/:id/suggest", tournament.AddSuggestion)
		rounds.Delete("/:id", tournament.RemoveRound)
	}

	votes := app.Group("/votes")
	{
		votes.Post("/:id/:suggestionId", tournament.AddVote)
	}
}
