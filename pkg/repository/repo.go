package repository

import (
	"fiberBoilerplate/pkg/ent"
	"github.com/gofiber/fiber/v2"
)

type Repo struct {

	// Database connection
	Conn ent.Client
}

type IRepo interface {
	Create(ctx *fiber.Ctx) error
	Index(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
