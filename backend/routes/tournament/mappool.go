package tournament

import (
	"backend/models"
	"backend/services/tournamentservice"

	"github.com/gofiber/fiber/v2"
)

func CreateRound(c *fiber.Ctx) error {
	// bodyparse the RoundDto
	// if the body is not parsable, return a 400
	// if the body is not parsable, return a 400
	var roundDto models.RoundDto
	if err := c.BodyParser(&roundDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	token := c.Get("Authorization")
	// call the service to remove the testplayer
	round, err := tournamentservice.AddRound(token, roundDto)

	// if the service returns an error, return a 400
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// return the round
	return c.JSON(models.RoundDtoFromEntity(round))
}

func AddSuggestion(c *fiber.Ctx) error {
	// bodyparse the SuggestionDto
	// if the body is not parsable, return a 400
	var suggestionDto models.SuggestionDto
	if err := c.BodyParser(&suggestionDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	roundId := c.Params("id")

	token := c.Get("Authorization")
	// call the service to remove the testplayer
	suggestion, err := tournamentservice.AddSuggestion(token, suggestionDto, roundId)

	// if the service returns an error, return a 400
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// return the suggestion
	return c.JSON(models.SuggestionDtoFromEntity(suggestion))

}
