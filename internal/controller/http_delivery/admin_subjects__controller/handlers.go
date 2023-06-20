package adminSubjectController

import (
	"net/http"

	"electronic_diary/internal/domain/subject/dto"
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
	body, err := api.ParseAndValidateJSON[dto.CreateSubjectDTO](ctx)
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
	body, err := api.ParseAndValidateJSON[dto.UpdateSubjectDTO](ctx)
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
