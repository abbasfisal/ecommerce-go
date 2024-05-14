package requests

type LoginRequest struct {
	Mobile   string `form:"mobile" validate:"required"`
	Password string `form:"password" validate:"required"`
}
