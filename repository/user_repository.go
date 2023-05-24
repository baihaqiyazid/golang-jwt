package repository

import (
	"database/sql"
	"go_fiber/model/entity"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	Create(c *fiber.Ctx, tx *sql.Tx, user entity.User) entity.User
	Update(c *fiber.Ctx, tx *sql.Tx, user entity.User) entity.User 
	FindAll(c *fiber.Ctx, tx *sql.Tx) []entity.User
	FindById(c *fiber.Ctx, tx *sql.Tx, userId int) (entity.User, error) 
	FindByEmail(c *fiber.Ctx, tx *sql.Tx, email string) (entity.User, error)
	Delete(c *fiber.Ctx, tx *sql.Tx, user entity.User)
	SetStatusIsVerified(c *fiber.Ctx, tx *sql.Tx, email string)
}