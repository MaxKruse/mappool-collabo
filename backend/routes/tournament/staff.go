package tournament

import (
	"backend/models"
	"backend/services/tournamentservice"

	"github.com/gofiber/fiber/v2"
)

func AddMappooler(ctx *fiber.Ctx) error {
	// make sure the user is logged in by checking for the Authorization header
	// if the user is not logged in, return a 401

	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// bodyparse the MappoolerDto
	// if the body is not parsable, return a 400
	var mappoolerDto models.MappoolerDto
	if err := ctx.BodyParser(&mappoolerDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	// call the service to add the mappooler
	err := tournamentservice.AddMappooler(token[7:], mappoolerDto.TournamentID, mappoolerDto.UserID)

	// if the service returns an error, return a 400
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// return a 200
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Mappooler added"})
}
