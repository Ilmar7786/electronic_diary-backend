package subjectsController

import (
	"net/http"

	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
)

// @Tags 	 	Предметы
// @Summary  	Список предметов
// @Accept 	 	json
// @Produce  	json
// @Success  	200 {object} []subject.Model
// @Router 		/subjects [get]
func (d DeliveryHttp) handlerSubjectFindAll(ctx *gin.Context) {
	subjects := d.subjectUC.FindAll()
	ctx.JSON(http.StatusOK, subjects)
}

// @Tags 	 	Предметы
// @Summary  	Список предметов
// @Accept 	 	json
// @Produce  	json
// @Success  	200 {object} []subject.Model
// @Failure	 	400,404 {object}  api.ResponseError
// @Param  	 	id path string true "Индефикатор предмета"
// @Router 		/subjects/{id} [get]
func (d DeliveryHttp) handlerSubjectById(ctx *gin.Context) {
	id := ctx.Param("id")
	candidate, err := d.subjectUC.FindByID(id)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, candidate)
}
