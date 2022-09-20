package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	sess *session.Store
)

func init() {
	sess = session.New()
}

func GetSession(ctx *fiber.Ctx) (*session.Session, error) {
	return sess.Get(ctx)
}
