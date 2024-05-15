package requests

type CreateRegisterRequest struct {
	FirstName       string `form:"first_name" validate:"required,min=3,max=10"`
	LastName        string `form:"last_name" validate:"required,min=3,max=10"`
	Mobile          string `form:"mobile" validate:"required,len=11,mobile"`
	Password        string `form:"password" validate:"required,min=5"`
	ConfirmPassword string `form:"confirm_password" validate:"required,eqfield=Password"`
	NationalCode    string `form:"national_code" validate:"required"` //todo:not necessary
	Number          string `form:"number" validate:"required"`
	Floor           int    `form:"floor" validate:"required"`
	Description     string `form:"description" validate:"required"`
}
