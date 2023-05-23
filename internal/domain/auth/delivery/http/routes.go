package http

import (
	"electronic_diary/internal/domain/auth"

	"github.com/gin-gonic/gin"
)

const pathUrlAPI = "auth"

type DeliveryHttpAuth struct {
	router *gin.RouterGroup

	AuthUC auth.UseCase
}

func NewDeliveryHttpAuth(router *gin.RouterGroup, authUC auth.UseCase) {
	prefix := router.Group(pathUrlAPI)
	deliveryHttp := DeliveryHttpAuth{
		router: router,
		AuthUC: authUC,
	}

	prefix.POST("/sign-in", deliveryHttp.handlerSignIn)
	prefix.POST("/sign-up", deliveryHttp.handlerSignUp)
}
