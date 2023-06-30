package dto

import (
	"electronic_diary/internal/domain/subject"
	"electronic_diary/internal/domain/user/dto"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateTeacherDTO struct {
	Subject []*subject.Model  `json:"subject"`
	User    dto.CreateUserDTO `json:"user"`
}

func (c CreateTeacherDTO) Validate() error {
	if err := c.User.Validate(); err != nil {
		return err
	}

	return validation.ValidateStruct(&c,
		validation.Field(&c.Subject, validation.Required, validation.Length(1, 100)),
	)
}
