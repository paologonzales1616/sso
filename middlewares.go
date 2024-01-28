package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Middlewares(app *fiber.App) {
	app.Use(logger.New())
}
