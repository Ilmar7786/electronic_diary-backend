package parentController

import (
	"net/http"

	"electronic_diary/internal/domain/parent/dto"
	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
)

// @Tags 	 	Родитель
// @Summary  	Создать
// @Accept 	 	json
// @Produce  	json
// @Param	 	request body dto.CreateParentDTO true " "
// @Success  	200 {object} parent.Model
// @Failure 	400,401,404 {object}  api.ResponseError
// @Router 		/parent [post]
func (d DeliveryHttp) handlerCreate(ctx *gin.Context) {
	body, err := api.ParseAndValidateJSON[dto.CreateParentDTO](ctx)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	parent, err := d.parentUC.Create(body)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, parent)
}
