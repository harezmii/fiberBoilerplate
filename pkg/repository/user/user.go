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

// Store ShowAccount godoc
// @Summary      Create Data
// @Description  create users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body body  model.User  false   "User form"
// @Success      201  {object}  model.Response
// @Router       /users [post]
func (r RepoUser) Store(ctx *fiber.Ctx) error {
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

// Index  ShowAccount godoc
// @Summary      All  Data
// @Description  get all users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Response
// @Router       /users [get]
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

// Show ShowAccount godoc
// @Summary      Show Data
// @Description  get string by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "User ID"
// @Success      200  {object}  model.Response
// @Router       /users/{id} [get]
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

// Update ShowAccount godoc
// @Summary      User Update Data
// @Description  update user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path  string  true   "User Id"
// @Param        body  body model.User  false   "User update form"
// @Success      200  {object}  model.User
// @Router       /users/{id} [put]
func (r RepoUser) Update(ctx *fiber.Ctx) error {
	return nil
}

// Delete ShowAccount godoc
// @Summary      Delete Data
// @Description  delete users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "User ID"
// @Success      200  {object}  model.Response
// @Router       /users/{id} [delete]
func (r RepoUser) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	userDelete := r.Conn.User.DeleteOneID(id).Exec(r.Context)

	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		StatusCode: fiber.StatusOK,
		Message:    "User Deleted",
		Data:       userDelete,
	})
}
