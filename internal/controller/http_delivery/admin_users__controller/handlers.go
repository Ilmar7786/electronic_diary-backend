package adminUsersController

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
// @Success  	200 {object} user.Model
// @Failure  	400,401 {object} api.ResponseError
// @Param	 	request body dto.CreateUserDTO true " "
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
// @Failure	 	400,401 {object}  api.ResponseError
// @Router 		/admin/users [get]
func (d DeliveryHttpAdmin) handlerUserFindAll(ctx *gin.Context) {
	users := d.userUC.FindAll()
	ctx.JSON(http.StatusOK, users)
}

// @Tags 	 	Администратор
// @Summary  	Получить пользователя
// @Security 	ApiKeyAuth
// @Accept 	 	json
// @Produce  	json
// @Success  	200 {object} user.Model
// @Failure	 	400,401,404 {object}  api.ResponseError
// @Param  	 	id path string true "Индефикатор пользователя"
// @Router 		/admin/users/{id} [get]
func (d DeliveryHttpAdmin) handlerUserFindById(ctx *gin.Context) {
	id := ctx.Param("id")
	candidate, err := d.userUC.FindByID(id)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, candidate)
}

// @Tags 	 	Администратор
// @Summary  	Обновить пользователя
// @Security 	ApiKeyAuth
// @Accept 	 	json
// @Produce  	json
// @Success  	200 {object} api.Response
// @Failure	 	400,401,404 {object}  api.ResponseError
// @Param  	 	id path string true "Индефикатор пользователя"
// @Param  	 	request body dto.UpdateUserDTO true " "
// @Router 		/admin/users/{id} [patch]
func (d DeliveryHttpAdmin) handlerUserUpdateByID(ctx *gin.Context) {
	id := ctx.Param("id")
	body, err := api.ParseAndValidateJSON[dto.UpdateUserDTO](ctx)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := d.userUC.UpdateById(id, body); err != nil {
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
// @Success  	200 {object} api.Response
// @Failure 	400,401,404 {object}  api.ResponseError
// @Param  	 	id path string true "Индефикатор пользователя"
// @Router 		/admin/users/{id} [delete]
func (d DeliveryHttpAdmin) handlerUserDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := d.userUC.Delete(id)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "ok"})
}
