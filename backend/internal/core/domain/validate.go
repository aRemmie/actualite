package domain

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	validate = newValidator()

	usernameRe = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
)

func newValidator() *validator.Validate {
	v := validator.New()

	_ = v.RegisterValidation("username_validation", validateUsername)

	return v
}

func validateUsername(fl validator.FieldLevel) bool {
	return usernameRe.MatchString(fl.Field().String())
}
