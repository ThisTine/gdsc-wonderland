package fiber

import (
	"densomap-backend/types/response"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
		Success: false,
		Message: fmt.Sprintf("%s %s not found", c.Method(), c.Path()),
		Error:   "404_NOT_FOUND",
	})
}
