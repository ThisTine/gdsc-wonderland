package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"backend/modules"
)

var Recover = func() fiber.Handler {
	if mod.Conf.Environment == 1 {
		return func(c *fiber.Ctx) error {
			return c.Next()
		}
	}
	return recover.New()
}
