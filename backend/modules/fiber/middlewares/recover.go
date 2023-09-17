package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"backend/modules"
)

var Recover = func() fiber.Handler {
	if mod.Conf.Environment == 1 {
		return func(c *fiber.Ctx) error {
			// defer logrus.Debug("CALL " + c.Method() + " " + c.Path() + " " + string(c.Request().Body()))
			return c.Next()
		}
	}
	return recover.New()
}
