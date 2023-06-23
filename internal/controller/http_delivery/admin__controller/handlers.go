package adminController

import (
	"net/http"

	subjectDto "electronic_diary/internal/domain/subject/dto"
	userDto "electronic_diary/internal/domain/user/dto"
	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
)

// @Tags 	 	Администратор
// @Summary 	Создать предмет
// @Security 	ApiKeyAuth
// @Accept 	 	json
// @Produce  	json
// @Param	 	request body dto.CreateSubjectDTO true " "
// @Success  	200 {object} subject.Model
// @Failure  	400,401 {object} api.ResponseError
// @Router 		/admin/subject [post]
func (d DeliveryHttpAdmin) handlerSubjectCreate(ctx *gin.Context) {
	body, err := api.ParseAndValidateJSON[subjectDto.CreateSubjectDTO](ctx)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	subject, err := d.subjectUC.Create(body)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, subject)
}

// @Tags 	 	Администратор
// @Summary  	Список предметов
// @Security 	ApiKeyAuth
// @Accept 	 	json
// @Produce  	json
// @Success  	200 {object} []subject.Model
// @Failure	 	400,401 {object} api.ResponseError
// @Router 		/admin/subject [get]
func (d DeliveryHttpAdmin) handlerSubjectFindAll(ctx *gin.Context) {
	subjects := d.subjectUC.FindAll()
	ctx.JSON(http.StatusOK, subjects)
}

// @Tags 	 	Администратор
// @Summary  	Получить предмет
// @Security 	ApiKeyAuth
// @Accept 	 	json
// @Produce  	json
// @Success  	200 {object} subject.Model
// @Failure	 	400,401,404 {object} api.ResponseError
// @Param  	 	id path string true "Индефикатор предмета"
// @Router 		/admin/subject/{id} [get]
func (d DeliveryHttpAdmin) handlerSubjectById(ctx *gin.Context) {
	id := ctx.Param("id")
	candidate, err := d.subjectUC.FindByID(id)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, candidate)
}

// @Tags 	 	Администратор
// @Summary  	Обновить предмет
// @Security 	ApiKeyAuth
// @Accept 	 	json
// @Produce  	json
// @Param  	 	request body dto.UpdateSubjectDTO true " "
// @Success  	200 {object} api.Response
// @Failure	 	400,401,404 {object}  api.ResponseError
// @Param  	 	id path string true "Индефикатор предмета"
// @Router 		/admin/subject/{id} [patch]
func (d DeliveryHttpAdmin) handlerSubjectUpdateByID(ctx *gin.Context) {
	id := ctx.Param("id")
	body, err := api.ParseAndValidateJSON[subjectDto.UpdateSubjectDTO](ctx)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := d.subjectUC.UpdateById(id, body); err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "ok"})
}

// @Tags 	 	Администратор
// @Summary  	Удалить предмет
// @Security 	ApiKeyAuth
// @Accept 	 	json
// @Produce  	json
// @Success  	200 {object} api.Response
// @Failure 	400,401,404 {object}  api.ResponseError
// @Param  	 	id path string true "Индефикатор предмета"
// @Router 		/admin/subject/{id} [delete]
func (d DeliveryHttpAdmin) handlerSubjectDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := d.subjectUC.Delete(id)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "ok"})
}

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
	body, err := api.ParseAndValidateJSON[userDto.CreateUserDTO](ctx)
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
	body, err := api.ParseAndValidateJSON[userDto.UpdateUserDTO](ctx)
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
