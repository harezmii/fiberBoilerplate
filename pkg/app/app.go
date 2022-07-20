package app

import (
	"fiberBoilerplate/pkg/ent"
	"fiberBoilerplate/pkg/route"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	// Database Connection

	client, connError := ent.Conn()
	if connError != nil {
		return
	}
	// app routes

	route.SetupRoute(app, client)

	serverError := app.Listen(":1453")
	if serverError != nil {
		return
	}
}
