package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vashp/ms-user/src/cmd/domain"
	"strconv"
)

func (c *Controller) CreateUser(ctx *fiber.Ctx) error {
	var user domain.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	if err := c.userService.Create(ctx.Context(), &user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "User not created",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(user)
}

func (c *Controller) GetUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is bad",
		})
	}

	user, err := c.userService.GetByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return ctx.JSON(user)
}

func (c *Controller) UpdateUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID is bad"})
	}

	var user domain.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	user.ID = id
	if err := c.userService.Update(ctx.Context(), &user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(user)
}

func (c *Controller) DeleteUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is bad",
		})
	}

	if err := c.userService.Delete(ctx.Context(), id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
