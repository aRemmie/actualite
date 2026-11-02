package domain

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

var AllowedRoles = []string{"admin", "user"}

type User struct {
	ID                     string `validate:"required,uuid"`
	Email                  string `validate:"email"`
	Password               string `validate:"min=8,max=64,ascii"`
	Role                   string `validate:"oneof=admin user"`
	CreatedAt              time.Time
	Username               string `validate:"min=4,max=15,username_validation"`
	DisplayName            string `validate:"omitempty,min=1,max=30"`
	AvatarURL              string
	Status                 string `validate:"omitempty,max=100"`
	EmailVerified          bool
	PasswordResetToken     string
	PasswordResetExpiresAt time.Time
}

func (u *User) Validate() error {
	validator := validator.New()

	err := validator.RegisterValidation("username_validation", validateUsername)
	if err != nil {
		return fmt.Errorf("validator err: %w", err)
	}

	if err := validator.Struct(u); err != nil {
		return fmt.Errorf("validator struct: %w", err)
	}

	return nil
}

func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return re.MatchString(username)
}
