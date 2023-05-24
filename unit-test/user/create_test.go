package user

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"go_fiber/controller"
	"go_fiber/route"

	"go_fiber/repository"

	"go_fiber/service"
	"io"

	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupDB() *sql.DB {

	var (
		username = "root"
		password = "root"
		host     = "localhost"
		port     = "3306"
		db_name  = "go-fiber-test"
	)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", username, password, host, port, db_name)

	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) *fiber.App {

	validator := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validator)
	userController := controller.NewUserController(userService)

	router := route.RouteInit(userController)

	return router
}

func truncateUsers(db *sql.DB) {
	db.Exec("TRUNCATE users")
}

func TestCreateUsersSuccess(t *testing.T) {
	db := setupDB()
	truncateUsers(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"name": "baihaqi",
		"email": "baihaqicool329@gmail.com",
		"password": "baihaqi!@#",
		"address": "Jl. nanananaaana",
		"phone": "0812388773342"
	}`)
	
	request := httptest.NewRequest("POST", "http://localhost:3000/users/create", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, _ := router.Test(request, -1)

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)


	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "success", responseBody["status"])
	assert.Equal(t, "baihaqi", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateUsersAlreadyEmail(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"name": "baihaqi",
		"email": "baihaqicool329@gmail.com",
		"password": "baihaqi!@#",
		"address": "Jl. nanananaaana",
		"phone": "0812388773342"
	}`)
	
	request := httptest.NewRequest("POST", "http://localhost:3000/users/create", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, _ := router.Test(request, -1)

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// log.Printf("response : %v", responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "bad request", responseBody["status"])
	assert.Equal(t, "Email already to use", responseBody["message"])
}

func TestCreateUserNoName(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"name": "",
		"email": "baihaqicool329@gmail.com",
		"password": "baihaqi!@#",
		"address": "Jl. nanananaaana",
		"phone": "0812388773342"
	}`)
	
	request := httptest.NewRequest("POST", "http://localhost:3000/users/create", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, _ := router.Test(request, -1)

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// log.Printf("response : %v", responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "bad request", responseBody["status"])
	assert.Equal(t, "required", responseBody["message"].(map[string]interface{})["name"])
}

func TestCreateUserNoEmail(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"name": "baihaqi",
		"email": "",
		"password": "baihaqi!@#",
		"address": "Jl. nanananaaana",
		"phone": "0812388773342"
	}`)
	
	request := httptest.NewRequest("POST", "http://localhost:3000/users/create", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, _ := router.Test(request, -1)

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	log.Printf("response : %v", responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "bad request", responseBody["status"])
	assert.Equal(t, "required", responseBody["message"].(map[string]interface{})["email"])
}

func TestCreateUserNoPassword(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"name": "baihaqi",
		"email": "baihaqicool329@gmail.com",
		"password": "",
		"address": "Jl. nanananaaana",
		"phone": "0812388773342"
	}`)
	
	request := httptest.NewRequest("POST", "http://localhost:3000/users/create", requestBody)
	request.Header.Set("Content-Type", "application/json")

	response, _ := router.Test(request, -1)

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// log.Printf("response : %v", responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "bad request", responseBody["status"])
	assert.Equal(t, "required", responseBody["message"].(map[string]interface{})["password"])
}

func TestCreateUser(t *testing.T)()  {
	db := setupDB()
	truncateUsers(db)
	router := setupRouter(db)

	var responseBody map[string]interface{}

	tests := []struct{
		name string
		request string
		expected map[string]interface{}
	}{
		{
			name: "CreateUserSuccess",
			request: `{
				"name": "baihaqi",
				"email": "baihaqicool329@gmail.com",
				"password": "baihaqi!@#",
				"address": "Jl. nanananaaana",
				"phone": "0812388773342"
			}`,
			expected: map[string]interface{}{
				"code" : 200,
				"status": "success",
				"name": "baihaqi",
				"email": "baihaqicool329@gmail.com",
				"password": "baihaqi!@#",
				"address": "Jl. nanananaaana",
				"phone": "0812388773342",
			},
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			requestTest := bytes.NewReader([]byte(test.request))
			request := httptest.NewRequest("POST", "http://localhost:3000/users/create", requestTest)
			request.Header.Set("Content-Type", "application/json")

			response, _ := router.Test(request, -1)

			assert.Equal(t, 200, response.StatusCode)

			body, _ := io.ReadAll(response.Body)
			json.Unmarshal(body, &responseBody)

			if test.name == "CreateUserSuccess" {
				assert.Equal(t, test.expected["code"], int(responseBody["code"].(float64)))
				assert.Equal(t, test.expected["status"], responseBody["status"])
				assert.Equal(t, test.expected["name"], responseBody["data"].(map[string]interface{})["name"])
				assert.Equal(t, test.expected["email"], responseBody["data"].(map[string]interface{})["email"])
				assert.Equal(t, test.expected["address"], responseBody["data"].(map[string]interface{})["address"])
				assert.Equal(t, test.expected["phone"], responseBody["data"].(map[string]interface{})["phone"])
			}
		})
	}
}