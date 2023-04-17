package http

import (
	"net/http"

	"electronic_diary/internal/domain/user/dto"
	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
)

// @Summary Авторизация
// @Tags 	users
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

// handlerFindByID - find user by id
func (d DeliveryHttpUser) handlerFindByID(ctx *gin.Context) {
	userId := ctx.Param("id")
	if userId == "" {
		api.NewErrorResponse(ctx, http.StatusBadRequest, "params - id not found")
		return
	}

	user, err := d.UserUC.FindById(userId)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// @Summary Список пользователей
// @Tags 	users
// @Accept 	json
// @Produce json
// @Success 200 {object} []user.Model
// @Failure 400 {object} api.ResponseError
// @Router /users/ [get]
// handlerFindAll - find all users
func (d DeliveryHttpUser) handlerFindAll(ctx *gin.Context) {
	users := d.UserUC.FindAll()
	ctx.JSON(http.StatusCreated, users)
}

// @Summary Обновить пользователя
// @Tags 	users
// @Accept 	json
// @Produce json
// @Param 	input 	  body   dto.UpdateUserDTO true "credentials"
// @Param 	user_id	  path	 string false "ID пользователя"
// @Success 200 {object} user.Model
// @Failure 400 {object} api.ResponseError
// @Router /users/{user_id} [put]
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

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

// handlerDelete - delete user
func (d DeliveryHttpUser) handlerDelete(ctx *gin.Context) {
	userId := ctx.Param("id")

	err := d.UserUC.DeleteById(userId)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}