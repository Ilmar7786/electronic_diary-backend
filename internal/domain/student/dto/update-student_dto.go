package dto

import (
	"electronic_diary/internal/domain/user/dto"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type UpdateStudentDTO struct {
	ResidentialAddress *string            `json:"residentialAddress"`
	UserID             *uuid.UUID         `json:"userId"`
	ParentID           *uuid.UUID         `json:"parentId"`
	User               *dto.UpdateUserDTO `json:"user"`
}

func (c UpdateStudentDTO) Validate() error {
	if err := c.User.Validate(); err != nil {
		return err
	}

	return validation.ValidateStruct(&c,
		validation.Field(&c.UserID, is.UUID),
		validation.Field(&c.ParentID, is.UUID),
	)
}
