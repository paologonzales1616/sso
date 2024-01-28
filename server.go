package main

import (
	"fmt"
	"sso/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/template/django/v3"
)

func RunServer() error {
	port := utils.Env.Port

	engine := django.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	Middlewares(app)

	app.Use(healthcheck.New())

	return app.Listen(fmt.Sprintf("127.0.0.1:%s", port))
}
