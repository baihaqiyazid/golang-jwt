package helper

import (
	"fmt"
	"go_fiber/exception"
	"go_fiber/model/entity"
	"go_fiber/model/web"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RandomStringBytes() string {
	
	rand.Seed(time.Now().UnixNano())
	number := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	
	b := make([]byte, 6)
	
	for i := range b{
	  b[i] =  number[rand.Intn(len(number))] 
	}
	
	  return string(b)
}
  
func SendEmail(to string, code string) error {
	
	from := "baihaqiyazid16@gmail.com"
	password := "dknlujktqwkefzvd"
	smptServer := "smtp.gmail.com:587"
	subject := "Verification Code"
	body := "Your verification code is: "+code

	message := "From: "+ from + "\n" +
	"To: " + to + "\n" +
	"Subject: " + subject + "\n\n" + 
	body
	
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	return smtp.SendMail(smptServer, auth, from, []string{to}, []byte(message))
}

func EmailVerify(c *fiber.Ctx, user entity.User, requestEmail string) {

	email := requestEmail
	code := RandomStringBytes()
	fmt.Println("Code:", code)
	err := SendEmail(email, code)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	var inputCode string
	fmt.Print("Enter code: ")
	fmt.Scanln(&inputCode)
	if inputCode == code {
		fmt.Println("Email verified.")
	} else {
		fmt.Println("Invalid code.")
	}
}

func AddToken(c *fiber.Ctx, user entity.User) web.UserTokenResponse {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	c.Cookie(&fiber.Cookie{
		Name: "Verification",
		Value: tokenString,
		Expires: time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
	})

	return web.ToUserTokenResponse(user)
}

