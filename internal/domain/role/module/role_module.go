package module

import (
	"electronic_diary/internal/domain/role"
	"electronic_diary/internal/domain/role/delivery/http"
	"electronic_diary/internal/domain/role/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Role struct {
	db *gorm.DB

	UseCase role.UseCase
}

func NewRoleModule(db *gorm.DB) *Role {
	roleUC := usecase.NewRoleUseCase(db)

	return &Role{
		db:      db,
		UseCase: roleUC,
	}
}

func (r Role) RegisterController(router *gin.RouterGroup) {
	http.NewDeliveryHttpRole(router, r.UseCase)
}
