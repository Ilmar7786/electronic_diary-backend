package http

import (
	"net/http"

	"electronic_diary/internal/domain/admin/dto"
	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
)

func (a DeliveryHttpAdmin) handlerCreate(ctx *gin.Context) {
	body, err := api.ParseAndValidateJSON[dto.CreateAdminDTO](ctx)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := a.AdminUC.Create(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusCreated, user)
}

func (a DeliveryHttpAdmin) handlerFindAll(ctx *gin.Context) {
	users := a.AdminUC.FindAll()
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
