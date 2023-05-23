package http

import (
	"net/http"

	"electronic_diary/internal/domain/auth/dto"
	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
)

// @Summary Авторизация
// @Tags 	Аунтификация
// @Accept 	json
// @Produce json
// @Param 	input body   dto.SignInDTO true "credentials"
// @Success 200 {object} role.Model
// @Failure 400 {object} api.ResponseError
// @Router /sign-in [post]
// handlerSignIn - authorization
func (d DeliveryHttpAuth) handlerSignIn(ctx *gin.Context) {
	body, err := api.ParseAndValidateJSON[dto.SignInDTO](ctx)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := d.AuthUC.SignIn(body)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary Регистрация
// @Tags 	Аунтификация
// @Accept 	json
// @Produce json
// @Param 	input body   dto.SignUpDTO true "credentials"
// @Success 201 {object} role.Model
// @Failure 400 {object} api.ResponseError
// @Router /sign-up [post]
// handlerSignUp - register
func (d DeliveryHttpAuth) handlerSignUp(ctx *gin.Context) {
	body, err := api.ParseAndValidateJSON[dto.SignUpDTO](ctx)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := d.AuthUC.SignUp(body)
	if err != nil {
		api.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
