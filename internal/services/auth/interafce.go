package auth

import (
	"electronic_diary/internal/services/auth/dto"

	"github.com/gin-gonic/gin"
)

type Service interface {
	SignIn(dto dto.SignInDTO) (*ResponseAuth, error)
	Middleware() gin.HandlerFunc
	RefreshToken(token string) (*Tokens, error)
	GetUserID(ctx *gin.Context) string
}
