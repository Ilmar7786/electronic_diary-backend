package adminController

import (
	"electronic_diary/internal/domain/subject"
	"electronic_diary/internal/domain/user"
	authService "electronic_diary/internal/services/auth"

	"github.com/gin-gonic/gin"
)

const pathUrlAPI = "admin"

type DeliveryHttpAdmin struct {
	router      *gin.RouterGroup
	authService authService.Service

	subjectUC subject.UseCase
	userUC    user.UseCase
}

func Register(
	router *gin.RouterGroup,
	auth authService.Service,
	subjectUC subject.UseCase,
	userUC user.UseCase,
) {
	prefix := router.Group(pathUrlAPI)

	deliveryHttp := DeliveryHttpAdmin{
		router:      router,
		authService: auth,
		subjectUC:   subjectUC,
		userUC:      userUC,
	}

	prefix.Use(auth.Middleware(&authService.MiddlewareOptions{IsAdmin: true}))
	subjects := prefix.Group("subject")
	{
		subjects.POST("/", deliveryHttp.handlerSubjectCreate)
		subjects.GET("/", deliveryHttp.handlerSubjectFindAll)
		subjects.GET("/:id", deliveryHttp.handlerSubjectById)
		subjects.PATCH("/:id", deliveryHttp.handlerSubjectUpdateByID)
		subjects.DELETE("/:id", deliveryHttp.handlerSubjectDelete)
	}

	users := prefix.Group("users")
	{
		users.POST("/", deliveryHttp.handlerUserCreate)
		users.GET("/", deliveryHttp.handlerUserFindAll)
		users.GET("/:id", deliveryHttp.handlerUserFindById)
		users.PATCH("/:id", deliveryHttp.handlerUserUpdateByID)
		users.DELETE("/:id", deliveryHttp.handlerUserDelete)
	}
}
