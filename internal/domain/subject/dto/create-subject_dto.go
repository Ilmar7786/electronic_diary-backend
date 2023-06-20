package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateSubjectDTO struct {
	Title string `json:"title" maxLength:"40" validate:"required"`
}

func (c CreateSubjectDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required, validation.Length(0, 40)),
	)
}
