package main

import (
	"github.com/sirupsen/logrus"

	mod "backend/modules"
	"backend/modules/config"
	"backend/modules/fiber"
	"backend/modules/mongo"
)

func main() {
	config.Init()
	mongo.Init()
	fiber.Init()

	// * Startup
	err := mod.Fiber.Listen(mod.Conf.Address)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to start server")
	}
}
