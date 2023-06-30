package dto

import (
	"electronic_diary/internal/domain/subject"
	"electronic_diary/internal/domain/user/dto"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type UpdateTeacherDTO struct {
	UserID  *uuid.UUID         `json:"userId"`
	Subject []*subject.Model   `json:"subject"`
	User    *dto.UpdateUserDTO `json:"user"`
}

func (c UpdateTeacherDTO) Validate() error {
	if err := c.User.Validate(); err != nil {
		return err
	}

	return validation.ValidateStruct(&c,
		validation.Field(&c.UserID, is.UUID),
		validation.Field(&c.Subject, validation.Length(0, 100)),
	)
}
