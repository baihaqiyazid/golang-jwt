package web

import (
	"go_fiber/model/entity"
	"time"
)

type UserResponse struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Address    string    `json:"address"`
	Phone      string    `json:"phone"`
	IsVerified string    `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToUserResponse(user entity.User) UserResponse {
	return UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		IsVerified: user.IsVerified,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
