package http

import (
	"electronic_diary/internal/domain/role"

	"github.com/gin-gonic/gin"
)

const pathUrlAPI = "roles"

type DeliveryHttpRole struct {
	router *gin.RouterGroup

	RoleUC role.UseCase
}

func NewDeliveryHttpRole(router *gin.RouterGroup, roleUC role.UseCase) {
	prefix := router.Group(pathUrlAPI)
	deliveryHttp := DeliveryHttpRole{
		router: router,
		RoleUC: roleUC,
	}

	prefix.POST("/", deliveryHttp.handlerCreate)
	prefix.GET("/", deliveryHttp.handlerFindAll)
	prefix.GET("/:id", deliveryHttp.handlerFindByID)
	prefix.PUT("/:id", deliveryHttp.handlerUpdate)
	prefix.DELETE("/:id", deliveryHttp.handlerDelete)
}
