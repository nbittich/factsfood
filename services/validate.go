package services

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
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
