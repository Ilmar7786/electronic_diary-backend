package teacherController

import (
	"fmt"

	"electronic_diary/internal/domain/teacher"

	"github.com/gin-gonic/gin"
)

const pathUrlAPI = "teachers"

type DeliveryHttp struct {
	router *gin.RouterGroup

	teacherUC teacher.UseCase
}

func Register(router *gin.RouterGroup, teacherUC teacher.UseCase) {
	prefix := router.Group(pathUrlAPI)

	deliveryHttp := DeliveryHttp{
		router:    router,
		teacherUC: teacherUC,
	}

	fmt.Println(prefix, deliveryHttp)
}
