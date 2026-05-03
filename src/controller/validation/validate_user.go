package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translator "github.com/go-playground/validator/v10/translations/en"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/rest_err"
)

var (
	validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translator.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid JSON format")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCause := []rest_err.Cause{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCause = append(errorsCause, cause)
		}
		return rest_err.NewBadRequestValidationError("some fields are invalid", errorsCause)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
