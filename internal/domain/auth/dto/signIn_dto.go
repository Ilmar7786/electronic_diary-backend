package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignInDTO struct {
	Email    string
	Password string
}

func (s SignInDTO) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Email, validation.Required, is.Email, validation.Length(1, 50)),
		validation.Field(&s.Password, validation.Required, validation.Length(8, 30)),
	)
}
