package route

import (
	"go_fiber/controller"
	"go_fiber/exception"
	"go_fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(userController controller.UserController) *fiber.App {

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				exception.ErrorHandler(c, err)
			}
		}()
	
		return c.Next()
	})

	app.Get("/users", middleware.AuthMiddleware, userController.FindAll)
	app.Get("/users/:userId", userController.FindById)
	app.Put("/users/:userId", middleware.AuthMiddleware, userController.Update)
	app.Delete("/users/:userId", middleware.AuthMiddleware, userController.Delete)
	app.Post("/users/create", userController.Create)

	app.Post("/login", userController.Login)
	app.Get("/logout", middleware.AuthMiddleware, userController.Logout)
	app.Post("/verify", userController.Verify)

	return app
}
