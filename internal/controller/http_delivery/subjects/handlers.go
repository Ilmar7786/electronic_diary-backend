package subjectController

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
// @Router 		/admin/subjects [post]
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
// @Failure	 	400,401,404 {object} api.ResponseError
// @Router 		/admin/subjects [get]
func (d DeliveryHttpAdmin) handlerSubjectFindAll(ctx *gin.Context) {
	subjects := d.subjectUC.FindAll()
	ctx.JSON(http.StatusOK, subjects)
}

// @Tags 	 	Администратор
// @Summary  	Обновить предмет
// @Security 	ApiKeyAuth
// @Accept 	 	json
// @Produce  	json
// @Param  	 	subjectId path string true "Индефикатор предмета"
// @Param  	 	request body dto.UpdateSubjectDTO true " "
// @Success  	200 {object} api.Response
// @Failure	 	400,401,404 {object}  api.ResponseError
// @Router 		/admin/subjects/{subjectId} [patch]
func (d DeliveryHttpAdmin) handlerSubjectUpdateByID(ctx *gin.Context) {
	id := ctx.Param("subjectId")
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
// @Param  	 	subjectId path string true "Индефикатор предмета"
// @Success  	200 {object} api.Response
// @Failure 	400,401,404 {object}  api.ResponseError
// @Router 		/admin/subjects/{subjectId} [delete]
func (d DeliveryHttpAdmin) handlerSubjectDelete(ctx *gin.Context) {
	id := ctx.Param("subjectId")
	err := d.subjectUC.Delete(id)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, api.Response{Message: "ok"})
}
