package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags 	 Пользователь
// @Summary  Информация о пользователи
// @Security ApiKeyAuth
// @Accept 	 json
// @Produce  json
// @Success  200 {object} user.Model
// @Failure 400 {object}  api.ResponseError
// @Failure 401 {object}  api.ResponseError
// @Router /user/info [get]
func (d DeliveryHttpUser) handlerUserInfo(ctx *gin.Context) {
	user := d.authService.GetUser(ctx)

	ctx.JSON(http.StatusOK, user)
}
