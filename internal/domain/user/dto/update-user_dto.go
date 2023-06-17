package dto

import (
	"electronic_diary/internal/constants"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UpdateUserDTO struct {
	Surname     *string         `json:"surname"`
	Name        *string         `json:"name"`
	Patronymic  *string         `json:"patronymic"`
	Address     *string         `json:"address"`
	Phone       *string         `json:"phone"`
	Email       *string         `json:"email"`
	Password    *string         `json:"password"`
	Role        *constants.Role `json:"role" enums:"STUDENT,TEACHER,PARENT"`
	IsActive    *bool           `json:"isActive"`
	IsSuperUser *bool           `json:"isSuperUser"`
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
