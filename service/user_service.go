package service

import (
	"go_fiber/model/web"

	"github.com/gofiber/fiber/v2"
)

type UserService interface{
	Create(c *fiber.Ctx, request web.UserCreateRequest) web.UserResponse
	Update(c *fiber.Ctx, request web.UserUpdateRequest) web.UserResponse
	FindAll(c *fiber.Ctx) []web.UserResponse
	FindById(c *fiber.Ctx, userId int) web.UserResponse
	Delete(c *fiber.Ctx, userId int)
	Login(c *fiber.Ctx, request web.UserLoginRequest) web.UserTokenResponse
	Verify(c *fiber.Ctx, request web.UserVerifyRequest)
}