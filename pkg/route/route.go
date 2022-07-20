package route

import (
	"fiberBoilerplate/pkg/ent"
	"fiberBoilerplate/pkg/repository/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App, client *ent.Client) {

	// User Routes
	userRepo := user.RepoUser{
		Repo: struct {
			Conn ent.Client
		}{Conn: *client}}

	app.Get("/user", userRepo.Index)
}
