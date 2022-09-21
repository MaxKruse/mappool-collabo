package users

import (
	"backend/models"
	"backend/models/entities"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

func List(ctx *fiber.Ctx) error {
	dbSession := services.GetDebugDBSession()

	var users []entities.User

	dbSession.Find(&users)

	res := []models.UserDto{}
	for _, user := range users {
		res = append(res, models.UserDtoFromEntity(user))
	}

	return ctx.JSON(res)
}

func Get(ctx *fiber.Ctx) error {
	dbSession := services.GetDebugDBSession()

	res := entities.User{}
	dbSession.Find(&res, "id = ?", ctx.Params("id"))

	if res.ID == 0 && res.Username == "" {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
	}

	return ctx.JSON(models.UserDtoFromEntity(res))
}
