package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UpdateUserDTO struct {
	Surname    *string `json:"surname"`
	Name       *string `json:"name"`
	Patronymic *string `json:"patronymic"`
	Email      *string `json:"email"`
	Password   *string `json:"password"`
}

func (u UpdateUserDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Surname, validation.Length(1, 40)),
		validation.Field(&u.Name, validation.Length(1, 40)),
		validation.Field(&u.Patronymic, validation.Length(1, 40)),
		validation.Field(&u.Email, is.Email, validation.Length(1, 50)),
		validation.Field(&u.Password, validation.Length(8, 30)),
	)
}
