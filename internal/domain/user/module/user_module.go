package module

import (
	"electronic_diary/internal/domain/user"
	"electronic_diary/internal/domain/user/delivery/http"
	"electronic_diary/internal/domain/user/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB

	UseCase user.UseCase
}

func NewUserModule(db *gorm.DB) user.Module {
	userUC := usecase.NewUser(db)

	return &User{
		db:      db,
		UseCase: userUC,
	}
}

func (u User) RegisterController(router *gin.RouterGroup) {
	http.NewDeliveryHttpUser(router, u.UseCase)
}

func (u User) GetUseCase() user.UseCase {
	return u.UseCase
}
