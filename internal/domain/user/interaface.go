package user

import (
	"electronic_diary/internal/domain/user/dto"

	"github.com/gin-gonic/gin"
)

type Module interface {
	RegisterController(router *gin.RouterGroup)
	GetUseCase() UseCase
}

type UseCase interface {
	Create(dto dto.CreateUserDTO) (*Model, error)
	FindByID(id string) (*Model, error)
	FindByEmail(email string) (*Model, error)
	FindAll() []*Model
	Delete(id string) error
	UpdateById(id string, dto dto.UpdateUserDTO) error
}
