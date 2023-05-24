package web

type UserCreateRequest struct {
	Name       string `validate:"required,min=3,max=50"  json:"name"`
	Email      string `validate:"required,email" json:"email"`
	Password   string `validate:"required,min=8" json:"password"`
	Address    string `validate:"required" json:"address"`
	VerifyCode string `json:"verify_code"`
	Phone      string `validate:"required,number" json:"phone"`
}

type UserUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Name     string `validate:"required,min=3,max=50" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `json:"password"`
	Address  string `validate:"required" json:"address"`
	Phone    string `validate:"required,number" json:"phone"`
}

type UserLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UserVerifyRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Code string `validate:"required" json:"code"`
}