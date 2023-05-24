package entity

import "time"

type User struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Address    string    `json:"address"`
	Phone      string    `json:"phone"`
	VerifyCode string    `json:"verify_code"`
	IsVerified string    `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
