package http

import (
	"net/http"

	"electronic_diary/internal/domain/role/dto"
	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
)

// @Summary Создать
// @Tags 	Роль
// @Accept 	json
// @Produce json
// @Param 	input body   dto.CreateRoleDTO true "credentials"
// @Success 201 {object} role.Model
// @Failure 400 {object} api.ResponseError
// @Router /roles [post]
// handlerCreate - create new role
func (d DeliveryHttpRole) handlerCreate(ctx *gin.Context) {
	body, err := api.ParseAndValidateJSON[dto.CreateRoleDTO](ctx)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	create, err := d.RoleUC.Create(body)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, create)
}

// @Summary Список
// @Tags 	Роль
// @Accept 	json
// @Produce json
// @Success 200 {object} []role.Model
// @Failure 400 {object} api.ResponseError
// @Router /roles [get]
// handlerFindAll - find all roles
func (d DeliveryHttpRole) handlerFindAll(ctx *gin.Context) {
	roles := d.RoleUC.FindAll()
	ctx.JSON(http.StatusOK, roles)
}

// @Summary Получить одну роль
// @Tags 	Роль
// @Accept 	json
// @Produce json
// @Param 	id	  path	 string false "ID роли"
// @Success 200 {object} role.Model
// @Failure 400 {object} api.ResponseError
// @Router /roles/{id} [get]
// handlerFindByID - find role by id
func (d DeliveryHttpRole) handlerFindByID(ctx *gin.Context) {
	roleId := ctx.Param("id")
	if roleId == "" {
		api.NewErrorResponse(ctx, http.StatusBadRequest, "params - id not found")
		return
	}

	role, err := d.RoleUC.FindByID(roleId)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, role)
}

// @Summary Обновить
// @Tags 	Роль
// @Accept 	json
// @Produce json
// @Param 	input 	  body   dto.UpdateRoleDTO true "credentials"
// @Param 	id	  path	 string false "ID роли"
// @Success 200 {bool}   true
// @Failure 400 {object} api.ResponseError
// @Router /roles/{id} [patch]
// handlerUpdate - update role
func (d DeliveryHttpRole) handlerUpdate(ctx *gin.Context) {
	roleId := ctx.Param("id")
	body, err := api.ParseAndValidateJSON[dto.UpdateRoleDTO](ctx)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = d.RoleUC.UpdateById(roleId, body)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, true)
}

// @Summary Удалить
// @Tags 	Роль
// @Accept 	json
// @Produce json
// @Param 	id	  path	 string false "ID роли"
// @Success 200 {bool} true
// @Failure 400 {bool} api.ResponseError
// @Router /roles/{id} [delete]
// handlerDelete - delete role
func (d DeliveryHttpRole) handlerDelete(ctx *gin.Context) {
	roleId := ctx.Param("id")

	err := d.RoleUC.Delete(roleId)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, true)
}
