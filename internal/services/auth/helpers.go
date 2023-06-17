package auth

import (
	"errors"
	"net/http"
	"strings"

	"electronic_diary/pkg/api"
	"electronic_diary/pkg/core"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type PayloadToken struct {
	UserID     string `json:"userId"`
	IsRemember bool   `json:"isRemember"`
}

func (a Auth) generateTokens(userId string, remember bool) (*Tokens, error) {
	accessToken, err := api.CreateToken(
		core.Ternary(remember, a.cfg.Jwt.AccessTokenExpiredIn, a.cfg.Jwt.AccessTokenExpiredInNotRemember),
		a.cfg.Jwt.AccessTokenPrivateKey,
		PayloadToken{
			UserID:     userId,
			IsRemember: remember,
		},
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := api.CreateToken(
		core.Ternary(remember, a.cfg.Jwt.RefreshTokenExpiredIn, a.cfg.Jwt.RefreshTokenExpiredInNotRemember),
		a.cfg.Jwt.RefreshTokenPrivateKey,
		PayloadToken{
			UserID:     userId,
			IsRemember: remember,
		},
	)
	if err != nil {
		return nil, err
	}

	tokens := &Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}

	return tokens, nil
}

func parseToken(c *gin.Context, privateKey string) (PayloadToken, error) {
	var accessToken string

	authorizationHeader := c.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	}

	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.NewError("you are not logged in"))
		return PayloadToken{}, errors.New("access token empty")
	}

	sub, err := api.ValidateToken(accessToken, privateKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.NewError(err.Error()))
		return PayloadToken{}, err
	}

	var approveType PayloadToken
	err = mapstructure.Decode(sub, &approveType)
	if err != nil {
		// Обработка ошибки преобразования
		c.AbortWithStatusJSON(http.StatusBadRequest, api.NewError(err.Error()))
		return PayloadToken{}, err
	}

	return approveType, nil
}
