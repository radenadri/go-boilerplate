package pkg

import (
	"github.com/go-playground/validator/v10"
	"github.com/radenadri/go-boilerplate/internal/delivery/dto/response"
)

var Validator *validator.Validate

func InitValidator() {
	Validator = validator.New()

	Validator.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		return len(password) >= 8 && len(password) <= 32
	})
}

func GetValidatorErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "password":
		return "Password must be at least 8 characters long"
	case "email":
		return "Invalid email format"
	case "min":
		return "Should be at least " + err.Param() + " characters long"
	case "max":
		return "Should not be longer than " + err.Param() + " characters"
	case "gte":
		return "Should be greater than or equal to " + err.Param()
	default:
		return "Invalid value"
	}
}

func FormatValidationErrors(err error) []response.ValidationError {
	var errors []response.ValidationError

	for _, err := range err.(validator.ValidationErrors) {
		var element response.ValidationError
		element.Field = err.Field()
		element.Rule = err.Tag()
		element.Value = err.Value().(string)
		element.Reason = GetValidatorErrorMessage(err)
		errors = append(errors, element)
	}

	return errors
}
