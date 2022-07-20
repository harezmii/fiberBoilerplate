package app

import (
	_ "fiberBoilerplate/pkg/docs"
	"fiberBoilerplate/pkg/ent"
	"fiberBoilerplate/pkg/route"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func Run() {
	app := fiber.New()

	//Swagger Initialize
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Database Connection

	client, connError := ent.Conn()
	if connError != nil {
		return
	}
	defer client.Close()
	// app routes

	route.SetupRoute(app, client)

	serverError := app.Listen(":1453")
	if serverError != nil {
		return
	}
}
