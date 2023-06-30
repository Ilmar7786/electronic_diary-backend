package parentController

import (
	"electronic_diary/internal/domain/parent"

	"github.com/gin-gonic/gin"
)

const pathUrlAPI = "parent"

type DeliveryHttp struct {
	router   *gin.RouterGroup
	parentUC parent.UseCase
}

func Register(router *gin.RouterGroup, parentUC parent.UseCase) {
	prefix := router.Group(pathUrlAPI)
	deliveryHttp := DeliveryHttp{
		router:   router,
		parentUC: parentUC,
	}

	prefix.POST("/", deliveryHttp.handlerCreate)
}
