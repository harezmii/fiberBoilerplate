package route

import (
	"context"
	"fiberBoilerplate/pkg/ent"
	"fiberBoilerplate/pkg/repository/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App, client *ent.Client) {

	// User Routes
	userRepo := user.RepoUser{
		Repo: struct {
			Conn    ent.Client
			Context context.Context
			Limit   int
		}{Conn: *client, Context: context.Background(), Limit: 10}}

	app.Get("/users", userRepo.Index)
	app.Get("/users/:id", userRepo.Show)
	app.Post("/users", userRepo.Store)
	app.Put("/users/:id", userRepo.Update)
	app.Delete("/users/:id", userRepo.Delete)
}
