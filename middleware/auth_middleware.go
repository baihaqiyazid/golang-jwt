package middleware

import (
	"errors"
	"fmt"
	"go_fiber/app"
	"go_fiber/exception"
	"go_fiber/helper"

	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(c *fiber.Ctx) error{
	tokenString := c.Cookies("Authorization") 

	if tokenString == "" {
		panic(exception.NewUserUnauthorized(errors.New("User Unauthorized").Error()))
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", token.Header["sub"])
		}

		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})


	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			panic(exception.NewUserUnauthorized(errors.New("User Unauthorized").Error()))
		}

		db := app.NewDB()
		tx, err := db.Begin()
		helper.PanicIfError(err)

		query :=`
			SELECT email FROM users 
			WHERE email = ? 
			LIMIT 1
		`
		user, err := tx.QueryContext(c.Context(), query, claims["sub"])

		if err != nil {
			panic(exception.NewUserUnauthorized(errors.New("User Unauthorized").Error()))
		}

		c.Locals("user", user)
		c.Next()

	}else {
		panic(exception.NewUserUnauthorized(errors.New("User Unauthorized").Error()))
	}

	return nil
}