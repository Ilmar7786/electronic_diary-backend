package module

import (
	"electronic_diary/internal/domain/admin"
	"electronic_diary/internal/domain/admin/delivery/http"
	"electronic_diary/internal/domain/admin/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Admin struct {
	db *gorm.DB

	UseCase admin.UseCase
}

func NewAdminModule(db *gorm.DB) *Admin {
	adminUC := usecase.NewAdminUseCase(db)

	return &Admin{
		db:      db,
		UseCase: adminUC,
	}
}

func (a Admin) RegisterController(router *gin.RouterGroup) {
	http.NewDeliveryHttpAdmin(router, a.UseCase)
}
