package adminController

import (
	"net/http"

	"electronic_diary/internal/domain/user/dto"
	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
)

// @Tags 	 	Администратор
// @Summary 	Создать пользователя
// @Security 	ApiKeyAuth
// @Accept 	 	json
// @Produce  	json
// @Param	 	request body dto.CreateUserDTO true " "
// @Success  	200 {object} user.Model
// @Failure  	400,401 {object} api.ResponseError
// @Router 		/admin/users [post]
func (d DeliveryHttpAdmin) handlerUserCreate(ctx *gin.Context) {
	body, err := api.ParseAndValidateJSON[dto.CreateUserDTO](ctx)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user, err := d.userUC.Create(body)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Tags 	 	Администратор
// @Summary  	Список пользователей
// @Security 	ApiKeyAuth
// @Accept 	 	json
// @Produce  	json
// @Success  	200 {object} []user.Model
// @Failure	 	400,401,404 {object}  api.ResponseError
// @Router 		/admin/users [get]
func (d DeliveryHttpAdmin) handlerUserFindAll(ctx *gin.Context) {
	users := d.userUC.FindAll()
	ctx.JSON(http.StatusOK, users)
}

// @Tags 	 	Администратор
// @Summary  	Обновить пользователя
// @Security 	ApiKeyAuth
// @Accept 	 	json
// @Produce  	json
// @Param  	 	userId path string true "Индефикатор пользователя"
// @Param  	 	request body dto.UpdateUserDTO true " "
// @Success  	200 {object} api.Response
// @Failure	 	400,401,404 {object}  api.ResponseError
// @Router 		/admin/users/{userId} [patch]
func (d DeliveryHttpAdmin) handlerUserUpdateByID(ctx *gin.Context) {
	userId := ctx.Param("userId")
	body, err := api.ParseAndValidateJSON[dto.UpdateUserDTO](ctx)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := d.userUC.UpdateById(userId, body); err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "ok"})
}

// @Tags 	 	Администратор
// @Summary  	Удалить пользователя
// @Security 	ApiKeyAuth
// @Accept 	 	json
// @Produce  	json
// @Param  	 	userId path string true "Индефикатор пользователя"
// @Success  	200 {object} api.Response
// @Failure 	400,401,404 {object}  api.ResponseError
// @Router 		/admin/users/{userId} [delete]
func (d DeliveryHttpAdmin) handlerUserDelete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	err := d.userUC.Delete(userId)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "ok"})
}
