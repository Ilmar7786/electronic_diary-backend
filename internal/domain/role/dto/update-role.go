package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type UpdateRoleDTO struct {
	Name *string `json:"name"`
	Key  *string `json:"key"`
}

func (r UpdateRoleDTO) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Length(4, 50)),
		validation.Field(&r.Key, validation.Length(4, 50)),
	)
}
