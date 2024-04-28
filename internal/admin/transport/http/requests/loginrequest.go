package requests

type LoginRequest struct {
	UserName string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}
