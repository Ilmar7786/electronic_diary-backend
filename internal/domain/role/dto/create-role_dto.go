package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CreateRoleDTO struct {
	Name string `json:"name"` // @required Обязательное поле
	Key  string `json:"key"`  // @required Обязательное поле
}

func (r CreateRoleDTO) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.Length(4, 50)),
		validation.Field(&r.Key, validation.Required, validation.Length(4, 50)),
	)
}
