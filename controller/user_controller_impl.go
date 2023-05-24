package controller

import (
	"go_fiber/helper"
	"go_fiber/model/web"
	"go_fiber/service"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(UserService service.UserService) UserController {
	return &UserControllerImpl{
		service: UserService,
	}
}

func (controller *UserControllerImpl) FindAll(c *fiber.Ctx) error {
	
	userResponse := controller.service.FindAll(c)
	webResponse := web.ToWebResponse(fiber.StatusOK, "success", nil, userResponse)

	return c.JSON(webResponse)
}

func (controller *UserControllerImpl) Create(c *fiber.Ctx) error {
	
	userCreateRequest := web.UserCreateRequest{}
	err := c.BodyParser(&userCreateRequest)
	helper.PanicIfError(err)

	userResponse := controller.service.Create(c, userCreateRequest)
	webResponse := web.ToWebResponse(fiber.StatusOK, "success", "Successful user registration! Please check the verification code on your email!", userResponse)

	return c.JSON(webResponse)
}

func (controller *UserControllerImpl) Update(c *fiber.Ctx) error {
	
	userUpdateRequest := web.UserUpdateRequest{}
	
	id, err := strconv.Atoi(c.Params("userId"))
	helper.PanicIfError(err)
	userUpdateRequest.Id = id

	err = c.BodyParser(&userUpdateRequest)
	helper.PanicIfError(err)

	userResponse := controller.service.Update(c, userUpdateRequest)
	webResponse := web.ToWebResponse(fiber.StatusOK, "success", "Successfully updated user data!", userResponse)
	
	return c.JSON(webResponse)
}

func (controller *UserControllerImpl) FindById(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("userId"))
	userResponse := controller.service.FindById(c, id)
	webResponse := web.ToWebResponse(fiber.StatusOK, "success", nil, userResponse)

	return c.JSON(webResponse)
}

func (controller *UserControllerImpl) Delete(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("userId"))
	controller.service.Delete(c, id)
	
	c.Cookie(&fiber.Cookie{
        Name:     "Authorization",
        Value:    "",
        Expires:  time.Unix(0, 0),
        HTTPOnly: true,
    })

	webResponse := web.ToWebResponse(fiber.StatusOK, "success", "Successfuly delete user!", nil)

	return c.JSON(webResponse)
}

func (controller *UserControllerImpl) Login(c *fiber.Ctx) error {
	
	userLoginRequest := web.UserLoginRequest{}
	err := c.BodyParser(&userLoginRequest)
	helper.PanicIfError(err)
	
	userResponse := controller.service.Login(c, userLoginRequest)

	webResponse := web.ToWebResponse(fiber.StatusOK, "success", "Successfuly logged in!", userResponse)

	return c.JSON(webResponse)

}

func (controller *UserControllerImpl) Logout(c *fiber.Ctx) error{
	c.ClearCookie("Authorization")

	webResponse := web.ToWebResponse(fiber.StatusOK,"success", "Successfuly logged out!", nil)

	return c.JSON(webResponse)
}

func (controller *UserControllerImpl) Verify(c *fiber.Ctx) error{
	
	userVerifyRequest := web.UserVerifyRequest{}
	err := c.BodyParser(&userVerifyRequest)
	helper.PanicIfError(err)

	controller.service.Verify(c, userVerifyRequest)

	webResponse := web.ToWebResponse(fiber.StatusOK,"success", "Successfuly verified email!", nil)

	return c.JSON(webResponse)
}