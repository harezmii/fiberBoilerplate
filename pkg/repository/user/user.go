package user

import (
	"fiberBoilerplate/pkg/model"
	"fiberBoilerplate/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type RepoUser struct {
	repository.Repo
}

func (r RepoUser) Create(ctx *fiber.Ctx) error {
	name := ctx.FormValue("name", "suat")
	email := ctx.FormValue("email", "suatcnby06@gmail.com")

	createError := r.Conn.User.Create().SetName(name).SetEmail(email).Exec(r.Context)
	if createError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "User Create Error: " + createError.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.Response{
		StatusCode: fiber.StatusCreated,
		Message:    "User created",
	})
}

func (r RepoUser) Index(ctx *fiber.Ctx) error {
	users, userError := r.Conn.User.Query().Limit(r.Limit).All(r.Context)
	if userError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "User Index Error: " + userError.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		StatusCode: fiber.StatusOK,
		Message:    "User Index",
		Data:       users,
	})
}
func (r RepoUser) Show(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	user, userGetError := r.Conn.User.Get(r.Context, id)

	if userGetError != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(model.Response{
			StatusCode: fiber.StatusNotFound,
			Message:    "User Show Error: " + userGetError.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		StatusCode: fiber.StatusOK,
		Message:    "User Show",
		Data:       user,
	})
}
func (r RepoUser) Update(ctx *fiber.Ctx) error {
	return nil
}
func (r RepoUser) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	userDelete := r.Conn.User.DeleteOneID(id).Exec(r.Context)

	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		StatusCode: fiber.StatusOK,
		Message:    "User Deleted",
		Data:       userDelete,
	})
}
