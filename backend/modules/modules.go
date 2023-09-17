package mod

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"backend/types/module"
)

var Conf *module.Config
var Fiber *fiber.App
var Database *mongo.Database
var Client *mongo.Client
