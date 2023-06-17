package authController

import (
	"electronic_diary/internal/services/auth"

	"github.com/gin-gonic/gin"
)

const pathUrlAPI = "auth"

type DeliveryHttpAuth struct {
	router *gin.RouterGroup

	AuthService auth.Service
}

func Register(router *gin.RouterGroup, AuthService auth.Service) {
	prefix := router.Group(pathUrlAPI)
	deliveryHttp := DeliveryHttpAuth{
		router:      router,
		AuthService: AuthService,
	}

	prefix.POST("/sign-in", deliveryHttp.handlerSignIn)
	prefix.POST("/refresh", deliveryHttp.handlerRefresh)
}
