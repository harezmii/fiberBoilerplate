package app

import "github.com/gofiber/fiber/v2"

func Run() {
	app := fiber.New()

	serverError := app.Listen(":1453")
	if serverError != nil {
		return
	}
}
