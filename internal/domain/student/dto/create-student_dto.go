package dto

import (
	"electronic_diary/internal/domain/user/dto"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type CreateStudentDTO struct {
	ResidentialAddress string            `json:"residentialAddress"`
	UserID             uuid.UUID         `json:"userId"`
	ParentID           uuid.UUID         `json:"parentId"`
	User               dto.CreateUserDTO `json:"user"`
}

func (c CreateStudentDTO) Validate() error {
	if err := c.User.Validate(); err != nil {
		return err
	}

	return validation.ValidateStruct(&c,
		validation.Field(&c.ResidentialAddress, validation.Required),
		validation.Field(&c.UserID, validation.Required, is.UUID),
		validation.Field(&c.ParentID, validation.Required, is.UUID),
	)
}
