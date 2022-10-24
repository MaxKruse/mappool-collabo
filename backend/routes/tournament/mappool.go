package tournament

import (
	"backend/models"
	"backend/services/exportservice"
	"backend/services/tournamentservice"
	"os"
	"strconv"

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

	// if the service returns an error, return a 500
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// return the suggestion
	return c.JSON(models.SuggestionDtoFromEntity(suggestion))

}

func RemoveRound(c *fiber.Ctx) error {
	roundId := c.Params("id")

	token := c.Get("Authorization")
	// call the service to remove the round
	err := tournamentservice.RemoveRound(token, roundId)

	// if the service returns an error, return a 500
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

func AddVote(c *fiber.Ctx) error {
	suggestionId := c.Params("suggestionId")
	var vote models.VoteDto
	if err := c.BodyParser(&vote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	token := c.Get("Authorization")

	// convert roundId and suggestionId to uint
	suggestionIdUint, err := strconv.ParseUint(suggestionId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	// call the service to add the vote
	err = tournamentservice.AddVote(token, uint(suggestionIdUint), vote)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func RemoveVote(c *fiber.Ctx) error {
	voteId := c.Params("voteId")

	token := c.Get("Authorization")
	err := tournamentservice.RemoveVote(token, voteId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// This is a download endpoint, therefore we either send the file or a regular string
func GetMappool(c *fiber.Ctx) error {
	format := c.Params("format")
	roundId := c.Params("roundId")

	// make a temporary file
	f, err := os.CreateTemp("", "mappool")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer os.Remove(f.Name())

	if format == exportservice.EXPORT_CSV {
		mappool, err := tournamentservice.GetMappool(roundId, format)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		exportservice.ExportCSV(f, mappool)

	} else if format == exportservice.EXPORT_LAZER {
		round, err := tournamentservice.GetRoundDeep(roundId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		exportservice.ExportLazer(f, round)
	}

	// check for the format. If we return CSV, we need to return a CSV file
	if format == exportservice.EXPORT_CSV {
		return c.Download(f.Name(), "mappool.csv")
	}

	// if we return lazer, we need to return a json file
	if format == exportservice.EXPORT_LAZER {
		return c.Download(f.Name(), "mappool.json")
	}

	return c.SendStatus(fiber.StatusNotImplemented)
}
