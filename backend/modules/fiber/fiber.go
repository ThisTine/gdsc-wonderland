package fiber

import (
	"github.com/gofiber/fiber/v2"

	"backend/endpoints"
	mod "backend/modules"
	"backend/modules/fiber/middlewares"
	"backend/types/response"
)

func Init() {
	// * Initialize fiber instance
	app := fiber.New(fiber.Config{
		ErrorHandler:  ErrorHandler,
		Prefork:       false,
		StrictRouting: true,
	})

	// * Register root endpoint
	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(response.Info("GDSC_WONDERLAND"))
	})

	// * Register API endpoints
	apiGroup := app.Group("api")
	apiGroup.Use(middlewares.Cors())
	apiGroup.Use(middlewares.Recover())
	endpoints.Init(apiGroup)

	// Register static files
	app.Static("/static", "./resources/static")

	// * Register not found endpoint
	app.Use(NotFoundHandler)

	// * Bind fiber instance to module
	mod.Fiber = app
}
