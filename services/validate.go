package services

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/nbittich/factsfood/types"
)

var (
	Validate    *validator.Validate
	initialized bool
)

func init() {
	if initialized {
		return
	}
	Validate = validator.New(validator.WithRequiredStructEnabled())
	Validate.RegisterValidation("password", validatePassword)
	initialized = true
}

func ValidateStruct(s interface{}) error {
	err := Validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			panic(err) // should never happen
		}

		validationErrors := err.(validator.ValidationErrors)
		errors := make(types.InvalidMessage, len(validationErrors))
		for i, err := range validationErrors {
			errors[i] = types.ErrorMessage{
				Field: err.Field(),
				Error: err.Tag(),
			}
		}
		return types.InvalidFormError{Messages: errors}
	}
	return nil
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Check for at least one uppercase letter
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	}

	// Check for at least one lowercase letter
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return false
	}

	// Check for at least one special character
	specialChars := regexp.MustCompile(`[^a-zA-Z0-9]`)

	return specialChars.MatchString(password)
}
