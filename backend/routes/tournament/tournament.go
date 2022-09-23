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
	tournament := tournamentservice.GetTournament(id)
	return c.JSON(tournament)
}

func Create(c *fiber.Ctx) error {
	// get self from context
	self, err := userservice.GetUserFromToken(c.Get("Authorization")[7:])
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// get tournament from request
	var tournament models.TournamentDto
	if err := c.BodyParser(&tournament); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad request"})
	}

	// set owner to self
	tournament.Owner = models.UserDtoFromEntity(self)

	// create tournament
	newTournament, err := tournamentservice.CreateTournament(tournament)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad request"})
	}

	return c.JSON(newTournament)
}

func Update(c *fiber.Ctx) error {
	return nil
}

func Delete(c *fiber.Ctx) error {
	return nil
}
