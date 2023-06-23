package userController

import (
	"electronic_diary/internal/domain/user"
	authService "electronic_diary/internal/services/auth"

	"github.com/gin-gonic/gin"
)

const pathUrlAPI = "user"

type DeliveryHttpUser struct {
	router *gin.RouterGroup

	userUC      user.UseCase
	authService authService.Service
}

func Register(router *gin.RouterGroup, auth authService.Service, userUC user.UseCase) {
	prefix := router.Group(pathUrlAPI)
	deliveryHttp := DeliveryHttpUser{
		router:      router,
		userUC:      userUC,
		authService: auth,
	}

	prefix.Use(auth.Middleware(&authService.MiddlewareOptions{}))
	prefix.GET("/info", deliveryHttp.handlerUserInfo)
}
