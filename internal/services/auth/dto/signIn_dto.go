package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignInDTO struct {
	Email    string `json:"email" maxLength:"100" validate:"required"`
	Password string `json:"password" minLength:"8" maxLength:"30" validate:"required"`
	Remember bool   `json:"remember" default:"false"`
}

func (s SignInDTO) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Email, validation.Required, is.Email, validation.Length(1, 50)),
		validation.Field(&s.Password, validation.Required, validation.Length(8, 30)),
	)
}
