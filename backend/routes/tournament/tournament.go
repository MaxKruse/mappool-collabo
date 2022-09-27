package tournament

import (
	"backend/models"
	"backend/services/tournamentservice"
	"backend/services/userservice"

	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	tournaments := tournamentservice.GetTournaments()
	return c.JSON(tournaments)
}

func Get(c *fiber.Ctx) error {
	id := c.Params("id")
	tournament, err := tournamentservice.GetTournament(id, tournamentservice.DepthBasic)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
	}
	return c.JSON(tournament)
}

func Create(c *fiber.Ctx) error {
	// get self from context
	self, err := userservice.GetUserFromToken(c.Get("Authorization")[7:])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// get tournament from request
	var tournament models.TournamentDto
	if err := c.BodyParser(&tournament); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	// set owner to self
	tournament.Owner = models.UserDtoFromEntity(self)

	// create tournament
	newTournament, err := tournamentservice.CreateTournament(tournament)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	return c.JSON(newTournament)
}

func Update(c *fiber.Ctx) error {
	// get self from context
	self, err := userservice.GetUserFromToken(c.Get("Authorization")[7:])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// get tournament from path param
	id := c.Params("id")
	tournament, err := tournamentservice.GetTournament(id, tournamentservice.DepthBasic)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
	}

	// check if self is owner
	if tournament.Owner.ID != self.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden"})
	}

	// get tournament from request
	var tournamentDto models.TournamentDto
	if err := c.BodyParser(&tournamentDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	tournamentDto.ID = tournament.ID

	// update tournament
	updatedTournament, err := tournamentservice.UpdateTournament(tournamentDto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	return c.JSON(updatedTournament)
}

func Delete(c *fiber.Ctx) error {
	// get self from context
	self, err := userservice.GetUserFromToken(c.Get("Authorization")[7:])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// get tournament from path param
	id := c.Params("id")
	tournament, err := tournamentservice.GetTournament(id, tournamentservice.DepthNone)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
	}

	// check if self is owner
	if tournament.Owner.ID != self.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden"})
	}

	// delete tournament
	tournamentservice.DeleteTournament(id)

	return c.SendStatus(fiber.StatusNoContent)
}
