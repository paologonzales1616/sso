package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type LoginData struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func Login(c *fiber.Ctx) error {
	payload := fiber.Map{}

	return c.Render("pages/login", payload)
}

func DoLogin(c *fiber.Ctx) error {
	user := new(LoginData)

	if err := c.BodyParser(user); err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fmt.Println(user.Email)
	fmt.Println(user.Password)

	return c.SendStatus(fiber.StatusOK)
}
