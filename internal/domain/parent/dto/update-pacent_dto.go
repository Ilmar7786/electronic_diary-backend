package dto

import (
	"electronic_diary/internal/domain/user/dto"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UpdateParentDTO struct {
	Guardian *string            `json:"guardian"`
	User     *dto.UpdateUserDTO `json:"user"`
}

func (d UpdateParentDTO) Validate() error {
	if err := d.User.Validate(); err != nil {
		return err
	}

	return validation.ValidateStruct(&d,
		validation.Field(&d.Guardian, validation.Length(0, 20)),
	)
}
