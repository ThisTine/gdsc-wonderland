package endpoints

import (
	"github.com/gofiber/fiber/v2"

	"backend/endpoints/pair"
)

func Init(router fiber.Router) {
	// * Map group
	pair := router.Group("/pair")
	pair.Post("/commit", pairEndpoint.CommitPostHandler)
}
