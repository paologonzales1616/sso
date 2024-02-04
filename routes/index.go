package routes

import (
	controllers "sso/controllers/auth"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Get("/login", controllers.Login)
	auth.Post("/login", controllers.DoLogin)
	auth.Post("/verify", controllers.Verify)
}
