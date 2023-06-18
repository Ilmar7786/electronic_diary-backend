package auth

import (
	"electronic_diary/internal/domain/user"
	"electronic_diary/internal/services/auth/dto"

	"github.com/gin-gonic/gin"
)

type Service interface {
	SignIn(dto dto.SignInDTO) (*ResponseAuth, error)
	Middleware(options *MiddlewareOptions) gin.HandlerFunc
	RefreshToken(token string) (*Tokens, error)
	GetUser(ctx *gin.Context) user.Model
}
