package beatmap

import (
	"backend/models"
	"backend/services/beatmapservice"
	"backend/util"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type downloadable struct {
	UUID     string
	Filepath string
	Filename string
}

var (
	availableReplayDownloads = map[uint]downloadable{}
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

	return c.Status(fiber.StatusOK).JSON(beatmapDto)
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

func GetReplayDownload(c *fiber.Ctx) error {
	replayId := c.Params("id")

	token := c.Get("Authorization")

	replay, err := beatmapservice.GetReplay(token, replayId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// if we already have the replay in our cache, send its uuid directly
	if entry, ok := availableReplayDownloads[replay.ID]; ok {
		return c.Status(fiber.StatusCreated).SendString(entry.UUID)
	}

	// make a uuid for the key
	downloadId := uuid.NewString()
	downloadName := replay.Map.Name + " - " + replay.Map.SlotName() + " by " + replay.User.Username + ".osr"

	availableReplayDownloads[replay.ID] = downloadable{
		UUID:     downloadId,
		Filepath: replay.Filepath,
		Filename: downloadName,
	}

	// spin up a goroutine to delete the entry for this identifier after 5 minutes
	go func() {
		<-time.After(5 * time.Minute)
		// Note that this will always try to delete the entry, even if it is already deleted
		delete(availableReplayDownloads, replay.ID)
	}()

	return c.Status(fiber.StatusCreated).SendString(downloadId)
}

func DownloadReplay(c *fiber.Ctx) error {
	identifier := c.Params("identifier")

	// find the downloadable entry for this identifier
	for _, entry := range availableReplayDownloads {
		if entry.UUID == identifier {
			return c.Download(entry.Filepath, entry.Filename)
		}
	}

	return c.Status(fiber.StatusNoContent).SendString("replay not found")
}
