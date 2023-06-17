package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type RefreshTokenDTO struct {
	Token string `json:"token"`
}

func (r RefreshTokenDTO) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Token, validation.Required),
	)
}
