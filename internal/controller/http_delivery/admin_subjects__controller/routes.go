package subjectController

import (
	"electronic_diary/internal/domain/subject"
	authService "electronic_diary/internal/services/auth"

	"github.com/gin-gonic/gin"
)

const pathUrlAPI = "admin"

type DeliveryHttpAdmin struct {
	router *gin.RouterGroup

	subjectUC   subject.UseCase
	authService authService.Service
}

func Register(router *gin.RouterGroup, auth authService.Service, subjectUC subject.UseCase) {
	prefix := router.Group(pathUrlAPI)

	deliveryHttp := DeliveryHttpAdmin{
		router:      router,
		subjectUC:   subjectUC,
		authService: auth,
	}

	prefix.Use(auth.Middleware(&authService.MiddlewareOptions{IsAdmin: true}))
	subjects := prefix.Group("subject")
	{
		subjects.POST("/", deliveryHttp.handlerSubjectCreate)
		subjects.GET("/", deliveryHttp.handlerSubjectFindAll)
		subjects.PATCH("/:subjectId", deliveryHttp.handlerSubjectUpdateByID)
		subjects.DELETE("/:subjectId", deliveryHttp.handlerSubjectDelete)
	}
}
