package adminController

import (
	"electronic_diary/internal/domain/user"
	authService "electronic_diary/internal/services/auth"

	"github.com/gin-gonic/gin"
)

const pathUrlAPI = "admin"

type DeliveryHttpAdmin struct {
	router *gin.RouterGroup

	userUC      user.UseCase
	authService authService.Service
}

func Register(router *gin.RouterGroup, auth authService.Service, userUC user.UseCase) {
	prefix := router.Group(pathUrlAPI)

	deliveryHttp := DeliveryHttpAdmin{
		router:      router,
		userUC:      userUC,
		authService: auth,
	}

	prefix.Use(auth.Middleware(&authService.MiddlewareOptions{IsAdmin: true}))
	users := prefix.Group("users")
	{
		users.POST("/", deliveryHttp.handlerUserCreate)
		users.GET("/", deliveryHttp.handlerUserFindAll)
		users.PATCH("/:userId", deliveryHttp.handlerUserUpdateByID)
		users.DELETE("/:userId", deliveryHttp.handlerUserDelete)
	}
}
