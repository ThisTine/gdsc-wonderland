package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	mod "backend/modules"
)

var Cors = func() fiber.Handler {
	// origins is the value of allowed CORS addresses, separated by comma (,).
	// Example: "https://www.google.com, https://www.bsthun.com, http://localhost:8080"
	origins := ""
	for i, s := range mod.Conf.Cors {
		origins += s
		if i < len(mod.Conf.Cors) {
			origins += ", "
		}
	}

	config := cors.Config{
		AllowOrigins:     origins,
		AllowCredentials: true,
	}

	return cors.New(config)
}
