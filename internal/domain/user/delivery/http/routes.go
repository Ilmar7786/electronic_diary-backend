package http

import (
	"electronic_diary/internal/domain/user"

	"github.com/gin-gonic/gin"
)

const pathUrlAPI = "users"

type DeliveryHttpUser struct {
	router *gin.RouterGroup

	UserUC user.UseCase
}

func NewDeliveryHttpUser(router *gin.RouterGroup, userUC user.UseCase) {
	prefix := router.Group(pathUrlAPI)
	deliveryHttp := DeliveryHttpUser{
		router: router,
		UserUC: userUC,
	}

	prefix.POST("/", deliveryHttp.handlerCreate)
	prefix.GET("/", deliveryHttp.handlerFindAll)
	prefix.GET("/:id", deliveryHttp.handlerFindByID)
	prefix.PUT("/:id", deliveryHttp.handlerUpdate)
	prefix.DELETE("/:id", deliveryHttp.handlerDelete)
}
