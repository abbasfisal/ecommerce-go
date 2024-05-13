package requests

type CreateRegisterRequest struct {
	FirstName       string `form:"first_name"`
	LastName        string `form:"last_name"`
	Mobile          string `form:"mobile"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm_password"`
	NationalCode    string `form:"national_code"`
	Number          string `form:"number"`
	Floor           int    `form:"floor"`
	Description     string `form:"description"`
}
