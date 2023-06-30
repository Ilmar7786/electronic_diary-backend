package dto

import (
	"electronic_diary/internal/domain/user/dto"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateParentDTO struct {
	Guardian string            `json:"guardian"`
	User     dto.CreateUserDTO `json:"user"`
}

func (d CreateParentDTO) Validate() error {
	if err := d.User.Validate(); err != nil {
		return err
	}

	return validation.ValidateStruct(&d,
		validation.Field(&d.Guardian, validation.Required, validation.Length(0, 20)),
	)
}
