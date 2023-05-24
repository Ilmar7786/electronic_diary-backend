package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignUpDTO struct {
	Surname    string `json:"surname"`    // @required Обязательное поле
	Name       string `json:"name"`       // @required Обязательное поле
	Patronymic string `json:"patronymic"` // @required Обязательное поле
	Email      string `json:"email"`      // @required Обязательное поле
	Password   string `json:"password"`   // @required Обязательное поле
}

func (u SignUpDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Surname, validation.Required, validation.Length(1, 40)),
		validation.Field(&u.Name, validation.Required, validation.Length(1, 40)),
		validation.Field(&u.Patronymic, validation.Required, validation.Length(1, 40)),
		validation.Field(&u.Email, validation.Required, is.Email, validation.Length(1, 50)),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 30)),
	)
}
