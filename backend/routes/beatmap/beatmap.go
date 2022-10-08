package beatmap

import (
	"backend/models"
	"backend/services/beatmapservice"
	"backend/util"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Get(c *fiber.Ctx) error {
	beatmapId := c.Params("id")

	token := c.Get("Authorization")

	beatmap, err := beatmapservice.GetBeatmap(token, beatmapId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	beatmapDto := models.MapDtoFromEntity(beatmap)

	return c.JSON(beatmapDto)
}

func AddReplay(c *fiber.Ctx) error {
	beatmapId := c.Params("id")

	token := c.Get("Authorization")

	// parse the replay file from formdata
	replayFile, err := c.FormFile("replay")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// save the replay file to disk
	replayPath := util.Config.StorageURI + "/" + uuid.NewString() + ".osr"
	err = c.SaveFile(replayFile, replayPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// add the replay to the database
	err = beatmapservice.AddReplay(token, beatmapId, replayPath)
	if err != nil {

		// delete the file from disk if the database insertion failed
		err = os.Remove(replayPath)
		if err != nil {
			log.Println("There was an error attempting to remove a stale replay file from disk: " + err.Error())
			log.Println("The file is located at: " + replayPath)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func DownloadReplay(c *fiber.Ctx) error {
	replayId := c.Params("id")

	token := c.Get("Authorization")

	replay, err := beatmapservice.GetReplay(token, replayId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	downloadName := replay.Map.Name + " - " + replay.Map.SlotName() + " by " + replay.User.Username + ".osr"

	return c.Download(replay.Filepath, downloadName)
}
