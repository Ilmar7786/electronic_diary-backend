package authController

import (
	"net/http"

	"electronic_diary/internal/services/auth/dto"
	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
)

// @Tags 	Аунтификация
// @Summary Авторизация
// @Accept 	json
// @Produce json
// @Param 	input body   dto.SignInDTO true "credentials"
// @Success 200 {object} user.Model
// @Failure 400 {object} api.ResponseErrors
// @Router /auth/sign-in [post]
func (d DeliveryHttpAuth) handlerSignIn(ctx *gin.Context) {
	body, err := api.ParseAndValidateJSON[dto.SignInDTO](ctx)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := d.AuthService.SignIn(body)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Tags 	Аунтификация
// @Summary Обновление токена
// @Accept 	json
// @Produce json
// @Param 	input body   dto.RefreshTokenDTO true "credentials"
// @Success 200 {object} auth.Tokens
// @Failure 400 {object} api.ResponseErrors
// @Router /auth/refresh [post]
func (d DeliveryHttpAuth) handlerRefresh(ctx *gin.Context) {
	body, err := api.ParseAndValidateJSON[dto.RefreshTokenDTO](ctx)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := d.AuthService.RefreshToken(body.Token)
	if err != nil {
		api.NewErrorsResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}
