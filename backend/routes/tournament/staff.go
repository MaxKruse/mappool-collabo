package tournament

import (
	"backend/models"
	"backend/services/tournamentservice"

	"github.com/gofiber/fiber/v2"
)

func AddMappooler(c *fiber.Ctx) error {
	// bodyparse the MappoolerDto
	// if the body is not parsable, return a 400
	var mappoolerDto models.MappoolerDto
	if err := c.BodyParser(&mappoolerDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	token := c.Get("Authorization")
	// call the service to add the mappooler
	err := tournamentservice.AddMappooler(token[7:], mappoolerDto.TournamentID, mappoolerDto.UserID)

	// if the service returns an error, return a 400
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// return a 200
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Mappooler added"})
}

func RemoveMappooler(c *fiber.Ctx) error {
	// bodyparse the MappoolerDto
	// if the body is not parsable, return a 400
	var mappoolerDto models.MappoolerDto
	if err := c.BodyParser(&mappoolerDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	token := c.Get("Authorization")
	// call the service to remove the mappooler
	err := tournamentservice.RemoveMappooler(token, mappoolerDto.TournamentID, mappoolerDto.UserID)

	// if the service returns an error, return a 400
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// return a 200
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Mappooler removed"})
}
