package dto

import (
	"electronic_diary/internal/constants"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UpdateUserDTO struct {
	Surname     *string        `json:"surname" maxLength:"40"`
	Name        *string        `json:"name" maxLength:"40"`
	Patronymic  *string        `json:"patronymic" maxLength:"40"`
	Address     *string        `json:"address" maxLength:"255"`
	Phone       *string        `json:"phone" maxLength:"30"`
	Email       *string        `json:"email" maxLength:"100"`
	Password    *string        `json:"password" minLength:"8" maxLength:"30"`
	Role        constants.Role `json:"role" enums:"student,teacher,parent"`
	IsActive    *bool          `json:"isActive" default:"false"`
	IsSuperUser *bool          `json:"isSuperUser" default:"false"`
}

func (u UpdateUserDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Surname, validation.Length(0, 40)),
		validation.Field(&u.Name, validation.Length(0, 40)),
		validation.Field(&u.Patronymic, validation.Length(0, 40)),
		validation.Field(&u.Address, validation.Length(0, 255)),
		validation.Field(&u.Phone, is.E164, validation.Length(0, 30)),
		validation.Field(&u.Email, is.Email, validation.Length(0, 100)),
		validation.Field(&u.Password, validation.Length(8, 30)),
		validation.Field(&u.Role, validation.Required, validation.In(constants.TeacherRole, constants.StudentRole, constants.StudentRole)),
		validation.Field(&u.IsActive),
		validation.Field(&u.IsSuperUser),
	)
}
