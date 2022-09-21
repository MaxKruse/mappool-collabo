package users

import (
	"backend/models"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

func List(ctx *fiber.Ctx) error {
	users := services.GetUsers()

	var res []models.UserDto
	for _, user := range users {
		res = append(res, models.UserDtoFromEntity(user))
	}

	return ctx.JSON(res)
}

func Get(ctx *fiber.Ctx) error {
	res, err := services.GetUserFromId(ctx.Params("id"))
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
	// get Authorization Bearer token from header
	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "no token provided",
		})
	}

	// strip the "Bearer" part
	token = token[7:]

	// get user from token
	user, err := services.GetUserFromToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(models.UserDtoFromEntity(user))
}
