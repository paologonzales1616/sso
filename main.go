package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func main() {
	port := 3000
	engine := django.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	error := app.Listen(fmt.Sprintf(":%d", port))
	if error != nil {
		log.Fatal("Error occured while starting server at port:", port)
	}
}
