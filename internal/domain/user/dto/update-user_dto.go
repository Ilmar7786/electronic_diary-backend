package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UpdateUserDTO struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsEmail    bool   `json:"isEmail"`
}

func (u UpdateUserDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Surname, validation.Required, validation.Length(1, 40)),
		validation.Field(&u.Name, validation.Required, validation.Length(1, 40)),
		validation.Field(&u.Patronymic, validation.Required, validation.Length(1, 40)),
		validation.Field(&u.Email, validation.Required, is.Email, validation.Length(1, 50)),
		validation.Field(&u.Password, validation.Required, validation.Length(1, 30)),
		validation.Field(&u.IsEmail, validation.Empty),
	)
}
