package user

import (
	"fiberBoilerplate/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

type RepoUser struct {
	repository.Repo
}

func (r RepoUser) Create(ctx *fiber.Ctx) error {
	return nil
}

func (r RepoUser) Index(ctx *fiber.Ctx) error {
	return ctx.JSON("İndex")
}
func (r RepoUser) Show(ctx *fiber.Ctx) error {
	return nil
}
func (r RepoUser) Update(ctx *fiber.Ctx) error {
	return nil
}
func (r RepoUser) Delete(ctx *fiber.Ctx) error {
	return nil
}
