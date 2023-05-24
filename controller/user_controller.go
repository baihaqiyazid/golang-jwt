package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error

	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	Verify(c *fiber.Ctx) error
}
