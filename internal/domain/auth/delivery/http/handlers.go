package http

import (
	"net/http"

	"electronic_diary/internal/domain/auth/dto"
	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
)

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

	ctx.JSON(http.StatusOK, user)
}
