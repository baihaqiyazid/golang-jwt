package helper

import (
	"go_fiber/model/web"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(user *web.UserCreateRequest) string {
	byte, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(byte)

	return user.Password
}

func CheckHashAndPassword(userPassword string, password string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	return err == nil
}