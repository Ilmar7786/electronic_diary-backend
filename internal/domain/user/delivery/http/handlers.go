package http

import (
	"net/http"

	"electronic_diary/internal/domain/user/dto"
	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
)

// @Summary Создать
// @Tags 	Пользователь
// @Accept 	json
// @Produce json
// @Param 	input body   dto.CreateUserDTO true "credentials"
// @Success 200 {object} user.Model
// @Failure 400 {object} api.ResponseError
// @Router /users [post]
// handlerCreate - create new user
func (d DeliveryHttpUser) handlerCreate(ctx *gin.Context) {
	body, err := api.ParseAndValidateJSON[dto.CreateUserDTO](ctx)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	create, err := d.UserUC.Create(body)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, create)
}

// @Summary Список пользователей
// @Tags 	Пользователь
// @Accept 	json
// @Produce json
// @Success 200 {object} []user.Model
// @Failure 400 {object} api.ResponseError
// @Router /users [get]
// handlerFindAll - find all users
func (d DeliveryHttpUser) handlerFindAll(ctx *gin.Context) {
	users := d.UserUC.FindAll()
	ctx.JSON(http.StatusOK, users)
}

// @Summary Получить одного пользователя
// @Tags 	Пользователь
// @Accept 	json
// @Produce json
// @Param 	id	  path	 string false "ID роли"
// @Success 200 {object} user.Model
// @Failure 400 {object} api.ResponseError
// @Router /users/{id} [get]
// handlerFindByID - find user by id
func (d DeliveryHttpUser) handlerFindByID(ctx *gin.Context) {
	userId := ctx.Param("id")
	if userId == "" {
		api.NewErrorResponse(ctx, http.StatusBadRequest, "params - id not found")
		return
	}

	user, err := d.UserUC.FindByID(userId)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary Обновить
// @Tags 	Пользователь
// @Accept 	json
// @Produce json
// @Param 	input 	  body   dto.UpdateUserDTO true "credentials"
// @Param 	id	  path	 string false "ID пользователя"
// @Success 200 {bool} 	 true
// @Failure 400 {object} api.ResponseError
// @Router /users/{id} [patch]
// handlerUpdate - update user
func (d DeliveryHttpUser) handlerUpdate(ctx *gin.Context) {
	userId := ctx.Param("id")
	body, err := api.ParseAndValidateJSON[dto.UpdateUserDTO](ctx)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = d.UserUC.UpdateById(userId, body)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, true)
}

// @Summary Удалить
// @Tags 	Пользователь
// @Accept 	json
// @Produce json
// @Param 	id	  path	 string false "ID пользователя"
// @Success 200 {bool} true
// @Failure 400 {bool} api.ResponseError
// @Router /users/{id} [delete]
// handlerDelete - delete user by id
func (d DeliveryHttpUser) handlerDelete(ctx *gin.Context) {
	userId := ctx.Param("id")

	err := d.UserUC.Delete(userId)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, true)
}
