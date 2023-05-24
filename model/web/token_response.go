package web

import "go_fiber/model/entity"

type UserTokenResponse struct{
	Name      string    `json:"name"`
	Email     string    `json:"email"`
}

func ToUserTokenResponse(user entity.User) UserTokenResponse{
	return UserTokenResponse{
		Name: user.Name,
		Email: user.Email,
	}
}