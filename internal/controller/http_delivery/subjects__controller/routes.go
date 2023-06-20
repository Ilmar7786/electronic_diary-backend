package subjectsController

import (
	"electronic_diary/internal/domain/subject"

	"github.com/gin-gonic/gin"
)

const pathUrlAPI = "subjects"

type DeliveryHttp struct {
	router *gin.RouterGroup

	subjectUC subject.UseCase
}

func Register(router *gin.RouterGroup, subjectUC subject.UseCase) {
	prefix := router.Group(pathUrlAPI)

	deliveryHttp := DeliveryHttp{
		router:    router,
		subjectUC: subjectUC,
	}

	prefix.GET("/", deliveryHttp.handlerSubjectFindAll)
	prefix.GET("/:id", deliveryHttp.handlerSubjectById)
}
