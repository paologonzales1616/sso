package main

import (
	"fmt"
	"sso/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func RunServer() error {
	port := utils.Env.Port

	engine := django.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	Middlewares(app)

	return app.Listen(fmt.Sprintf(":%s", port))
}
