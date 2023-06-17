package auth

import (
	"electronic_diary/internal/domain/user"

	"github.com/golang-jwt/jwt/v5"
)

type ResponseAuth struct {
	User   *user.Model `json:"user"`
	Tokens *Tokens     `json:"tokens"`
}

type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type Claims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}
