package user

import (
	"backend/models"
	"backend/services/userservice"

	"github.com/gofiber/fiber/v2"
)

func List(ctx *fiber.Ctx) error {
	users := userservice.GetUsers()

	var res []models.UserDto
	for _, user := range users {
		res = append(res, models.UserDtoFromEntity(user))
	}

	return ctx.JSON(res)
}

func Get(ctx *fiber.Ctx) error {
	res, err := userservice.GetUserFromId(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if res.ID == 0 && res.Username == "" {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
	}

	return ctx.JSON(models.UserDtoFromEntity(res))
}

func GetSelf(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")

	// get user from token
	user, err := userservice.GetUserFromToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(models.UserDtoFromEntity(user))
}
