package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UpdateSubjectDTO struct {
	Title *string `json:"title" maxLength:"40"`
}

func (c UpdateSubjectDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Length(0, 40)),
	)
}
