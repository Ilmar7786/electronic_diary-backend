package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CreateAdminDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (a CreateAdminDTO) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Login, validation.Required, validation.Length(4, 30)),
		validation.Field(&a.Password, validation.Required, validation.Length(8, 30)),
	)
}
