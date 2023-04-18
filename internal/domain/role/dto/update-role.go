package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type UpdateRoleDTO struct {
	Name string `json:"name"`
}

func (r UpdateRoleDTO) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.Length(4, 50)),
	)
}
