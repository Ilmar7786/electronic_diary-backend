package auth

import (
	"electronic_diary/internal/domain/auth/dto"
	"electronic_diary/internal/domain/user"

	"github.com/gin-gonic/gin"
)

type Module interface {
	RegisterController(router *gin.RouterGroup)
}

type UseCase interface {
	SignIn(dto dto.SignInDTO) (*user.Model, error)
	SignUp(dto dto.SignUpDTO) (*user.Model, error)
	LogOut()
}
