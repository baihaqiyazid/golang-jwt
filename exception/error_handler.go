package exception

import (
	"go_fiber/model/web"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err interface{}) {
	
	if notFoundError(c, err) {
		return
	}

	if badRequest(c, err){
		return
	}

	if validationError(c, err){
		return
	}

	if unauthorized(c, err){
		return
	}

	internalServerError(c, err)
}

func validationError(c *fiber.Ctx, err interface{}) bool {
	
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		errors := make(map[string]string)
		for _, fieldError := range exception{
			errors[strings.ToLower(fieldError.Field())] = fieldError.Tag()
		}

		webResponse := web.ToWebResponse(fiber.StatusBadRequest, "bad request", errors, nil)
		c.JSON(webResponse)

		return true
	}else {
		return false
	}
}

func internalServerError(c *fiber.Ctx, err interface{}){

	webResponse := web.ToWebResponse(fiber.StatusInternalServerError, "internal server error", nil ,nil)

	c.JSON(webResponse)
}

func notFoundError(c *fiber.Ctx, err interface{}) bool {
	
	exception, ok := err.(NotFoundError)

	if ok {
		webResponse := web.ToWebResponse(fiber.StatusNotFound, "not found", exception.Error, nil)
		c.JSON(webResponse)

		return true
		
	} else {
		return false
	}
}

func unauthorized(c *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(UnauthorizedError) 

	if ok {
		webResponse := web.ToWebResponse(fiber.StatusUnauthorized, "unauthorized", exception.Error, nil)
		c.JSON(webResponse)
		return true
	}else {
		return false
	}
}

func badRequest(c *fiber.Ctx, err interface{}) bool {
	
	exception, ok := err.(BadRequest)

	if ok {
		webResponse := web.ToWebResponse(fiber.StatusBadRequest, "bad request", exception.Error, nil)
		c.JSON(webResponse)

		return true
		
	} else {
		return false
	}
}

