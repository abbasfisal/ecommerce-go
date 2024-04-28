package requests

type LoginRequest struct {
	PhoneNumber string `form:"phone_number" validate:"required"`
	Password    string `form:"password" validate:"required"`
}
