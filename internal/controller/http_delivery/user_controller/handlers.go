package userController

import (
	"net/http"

	"electronic_diary/pkg/api"

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
	userId := d.authService.GetUserID(ctx)

	user, err := d.userUC.FindByID(userId)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}
