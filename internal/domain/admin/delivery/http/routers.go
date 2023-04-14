package http

import (
	"electronic_diary/internal/domain/admin"

	"github.com/gin-gonic/gin"
)

const pathUrlAPI = "admin"

type DeliveryHttpAdmin struct {
	router *gin.RouterGroup

	AdminUC admin.UseCase
}

func NewDeliveryHttpAdmin(router *gin.RouterGroup, adminUC admin.UseCase) {
	prefix := router.Group(pathUrlAPI)
	deliveryHttp := DeliveryHttpAdmin{
		router:  router,
		AdminUC: adminUC,
	}

	prefix.POST("/", deliveryHttp.handlerCreate)
	prefix.GET("/", deliveryHttp.handlerFindAll)
}
