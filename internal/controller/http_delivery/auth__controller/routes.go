package authController

import (
	"electronic_diary/internal/services/auth"

	"github.com/gin-gonic/gin"
)

const pathUrlAPI = "auth"

type DeliveryHttpAuth struct {
	router *gin.RouterGroup

	authService auth.Service
}

func Register(router *gin.RouterGroup, authService auth.Service) {
	prefix := router.Group(pathUrlAPI)
	deliveryHttp := DeliveryHttpAuth{
		router:      router,
		authService: authService,
	}

	prefix.POST("/sign-in", deliveryHttp.handlerSignIn)
	prefix.POST("/refresh", deliveryHttp.handlerRefresh)

	prefix.Use(authService.Middleware(&auth.MiddlewareOptions{}))
	prefix.GET("/user-info", deliveryHttp.handlerUserInfo)
}
