package user

import (
	"github.com/gin-gonic/gin"
)

type Module interface {
	RegisterController(router *gin.RouterGroup)
	GetUseCase() UseCase
}
