package pkg

import (
	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	faTranslations "github.com/go-playground/validator/v10/translations/fa"
	"regexp"
)

var trans ut.Translator

func NewValidate() *validator.Validate {

	faTranslator := fa.New()
	uni := ut.New(faTranslator, faTranslator)

	trans, _ = uni.GetTranslator("fa")
	validate := validator.New(validator.WithRequiredStructEnabled())
	_ = validate.RegisterValidation("mobile", ValidateMobileNumber)

	_ = faTranslations.RegisterDefaultTranslations(validate, trans)

	return validate
}
func GetTrans() ut.Translator {
	return trans
}

// CollectAndTranslateValidationErrors will return slice of validation errors which translated to selected language
func CollectAndTranslateValidationErrors(err error) []map[string]any {
	var validationError []map[string]any
	for _, err := range err.(validator.ValidationErrors) {
		validationError = append(validationError, map[string]any{
			err.Field(): err.Translate(GetTrans()),
		})
	}
	return validationError
}
func ValidateMobileNumber(fl validator.FieldLevel) bool {
	mobileNumber := fl.Field().String()
	regex := regexp.MustCompile(`^09\d{9}$`)
	return regex.MatchString(mobileNumber)
}
