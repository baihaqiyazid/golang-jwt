package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go_fiber/exception"
	"go_fiber/helper"
	"go_fiber/model/entity"
	"go_fiber/model/web"
	"go_fiber/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserServiceImpl struct{
	Repository repository.UserRepository
	DB *sql.DB
	Validator *validator.Validate
}

func NewUserService(Repository repository.UserRepository, DB *sql.DB, Validator *validator.Validate) UserService {
	return &UserServiceImpl{
		Repository: Repository,
		DB: DB,
		Validator: Validator,
	}
}

func (service *UserServiceImpl) Create(c *fiber.Ctx, request web.UserCreateRequest) web.UserResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.Repository.FindAll(c, tx)
	for _, u :=range users{
		if u.Email == request.Email{
			panic(exception.NewBadRequestError(errors.New("Email already to use").Error()))
		}
	}

	hashPassword := helper.GenerateHashPassword(&request)

	email := request.Email
	code := helper.RandomStringBytes()
	err = helper.SendEmail(email, code)

	if err != nil {
		fmt.Println("Error sending email:", err)
	}

	user := entity.User{
		Name: request.Name,
		Email: email,
		Password: hashPassword,
		Address: request.Address,
		VerifyCode: code,
		Phone: request.Phone,
	}

	helper.AddToken(c, user)

	user = service.Repository.Create(c, tx, user)
	
	return web.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(c *fiber.Ctx, request web.UserUpdateRequest) web.UserResponse {
	
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.FindById(c, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	users := service.Repository.FindAll(c, tx)
	for _, u :=range users{
		if u.Email == request.Email && u.Id != request.Id{
			panic(exception.NewBadRequestError(errors.New("Email already to use").Error()))
		}
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Address = request.Address
	user.Phone = request.Phone

	user = service.Repository.Update(c, tx, user)

	return web.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(c *fiber.Ctx) []web.UserResponse {

	tx, err:= service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	users := service.Repository.FindAll(c, tx)

	var userResponses []web.UserResponse
	for _, user := range users{
		userResponses = append(userResponses, web.UserResponse{
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
			Address: user.Address,
			IsVerified: user.IsVerified,
			Phone: user.Phone,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return userResponses
}

func (service *UserServiceImpl) FindById(c *fiber.Ctx, userId int) web.UserResponse {

	tx, err:= service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.FindById(c, tx, userId)

	if err != nil {	
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	return web.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(c *fiber.Ctx,  UserId int){
	
	tx, err:= service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.FindById(c, tx, UserId)
	if err != nil {	
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.Repository.Delete(c, tx, user)
}

func (service *UserServiceImpl) Login(c *fiber.Ctx, request web.UserLoginRequest) web.UserTokenResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)
	
	tx, err:= service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.FindByEmail(c, tx, request.Email)
	if err != nil{
		panic(exception.NewBadRequestError(err.Error()))
	}
	log.Printf("email: %s", user.Email);

	log.Println(user);
	log.Println(user.IsVerified)
	if user.IsVerified != "y"{
		panic(exception.NewBadRequestError(errors.New("You haven't verified email. Please click resend code!").Error()))
	}
	
	if helper.CheckHashAndPassword(user.Password, request.Password){
		
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.Email,
			"exp": time.Now().Add(time.Minute * 10).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
		if err != nil {
			panic(exception.NewBadRequestError(err.Error()))
		}

		c.Cookie(&fiber.Cookie{
			Name: "Authorization",
			Value: tokenString,
			Expires: time.Now().Add(time.Hour * 24 * 30),
			HTTPOnly: true,
		})

		return web.ToUserTokenResponse(user)

	}else{
		panic(exception.NewBadRequestError(errors.New("Email or Password wrong").Error()))
	}	
}

func (service *UserServiceImpl) Verify(c *fiber.Ctx, request web.UserVerifyRequest){
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Check cookie verification
	tokenString := c.Cookies("Verification") 

	if tokenString == "" {
		panic(exception.NewUserUnauthorized(errors.New("User Unauthorized").Error()))
	}

	// Parse token
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", token.Header["sub"])
		}

		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	// Check token isValid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			panic(exception.NewUserUnauthorized(errors.New("User Unauthorized").Error()))
		}

		user, err := service.Repository.FindByEmail(c, tx, request.Email)

		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		if request.Code == user.VerifyCode {
			tx, err:= service.DB.Begin()
			helper.PanicIfError(err)
			defer helper.CommitOrRollback(tx)

			service.Repository.SetStatusIsVerified(c, tx, user.Email)
			
		}else {
			panic(exception.NewBadRequestError(errors.New("Verification Code is wrong").Error()))
		}

		c.Locals("user", user.Email)
		c.Next()

	}else {
		panic(exception.NewUserUnauthorized(errors.New("User Unauthorized").Error()))
	}

	c.ClearCookie("Verification")
}

