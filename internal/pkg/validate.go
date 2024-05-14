package pkg

import (
	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	faTranslations "github.com/go-playground/validator/v10/translations/fa"
)

var trans ut.Translator

func NewValidate() *validator.Validate {

	faTranslator := fa.New()
	uni := ut.New(faTranslator, faTranslator)

	trans, _ = uni.GetTranslator("fa")
	validate := validator.New()
	_ = faTranslations.RegisterDefaultTranslations(validate, trans)

	return validate
}
func GetTrans() ut.Translator {
	return trans
}
