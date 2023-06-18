package dto

import (
	"electronic_diary/internal/constants"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateUserDTO struct {
	Surname     string         `json:"surname" maxLength:"40" validate:"required"`
	Name        string         `json:"name" maxLength:"40" validate:"required"`
	Patronymic  string         `json:"patronymic" maxLength:"40" validate:"required"`
	Address     string         `json:"address" maxLength:"255" validate:"required"`
	Phone       string         `json:"phone" maxLength:"30" validate:"required"`
	Email       string         `json:"email" maxLength:"100" validate:"required"`
	Password    string         `json:"password" maxLength:"30" validate:"required"`
	Role        constants.Role `json:"role" enums:"student,teacher,parent" validate:"required"`
	IsActive    bool           `json:"isActive" default:"false"`
	IsSuperUser bool           `json:"isSuperUser" default:"false"`
}

func (u CreateUserDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Surname, validation.Required, validation.Length(0, 40)),
		validation.Field(&u.Name, validation.Required, validation.Length(0, 40)),
		validation.Field(&u.Patronymic, validation.Required, validation.Length(0, 40)),
		validation.Field(&u.Address, validation.Required, validation.Length(0, 255)),
		validation.Field(&u.Phone, validation.Required, is.E164, validation.Length(0, 30)),
		validation.Field(&u.Email, validation.Required, is.Email, validation.Length(0, 100)),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 30)),
		validation.Field(&u.Role, validation.Required, validation.In(constants.TeacherRole, constants.ParentRole, constants.StudentRole)),
		validation.Field(&u.IsActive, validation.NotNil),
		validation.Field(&u.IsSuperUser, validation.NotNil),
	)
}
