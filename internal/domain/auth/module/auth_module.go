package module

import (
	"electronic_diary/internal/domain/auth"
	"electronic_diary/internal/domain/auth/delivery/http"
	"electronic_diary/internal/domain/auth/usecase"
	"electronic_diary/internal/domain/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Auth struct {
	db      *gorm.DB
	UseCase auth.UseCase
}

func NewAuthModule(db *gorm.DB, userUC user.UseCase) auth.Module {
	authUseCase := usecase.NewAuth(db, userUC)

	return &Auth{
		db:      db,
		UseCase: authUseCase,
	}
}

func (a Auth) RegisterController(router *gin.RouterGroup) {
	http.NewDeliveryHttpAuth(router, a.UseCase)
}
