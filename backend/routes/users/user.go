package users

import (
	"backend/models"
	"backend/models/entities"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

func List(ctx *fiber.Ctx) error {
	dbSession := services.GetDebugSession()

	var users []entities.User

	dbSession.Find(&users)

	res := []models.User{}
	for _, user := range users {
		res = append(res, user.BanchoUser)
	}

	return ctx.JSON(res)
}

func Get(ctx *fiber.Ctx) error {
	dbSession := services.GetDebugSession()

	user := entities.User{}
	dbSession.Find(&user, "id = ?", ctx.Params("id"))

	res := user.BanchoUser

	if res.ID == 0 && res.Username == "" {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
	}

	return ctx.JSON(res)
}
