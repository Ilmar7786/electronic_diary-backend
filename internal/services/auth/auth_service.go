package auth

import (
	"errors"

	"electronic_diary/internal/config"
	"electronic_diary/internal/domain/user"
	"electronic_diary/internal/services/auth/dto"
	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

const keyUserStorageCtx = "userId"

type Auth struct {
	userUC         user.UseCase
	authMiddleware gin.HandlerFunc

	cfg config.App
}

func New(userUC user.UseCase, cfg config.App) Service {
	authMiddleware := newMiddleware(cfg.Jwt.AccessTokenPrivateKey)

	return &Auth{
		userUC:         userUC,
		authMiddleware: authMiddleware,
		cfg:            cfg,
	}
}

func (a Auth) SignIn(dto dto.SignInDTO) (*ResponseAuth, error) {
	currentUser, err := a.userUC.FindByEmail(dto.Email)

	if err != nil || currentUser.ComparePassword(dto.Password) != nil {
		return nil, errors.New("not validation email or password")
	}

	tokens, err := a.generateTokens(currentUser.ID.String(), dto.Remember)
	if err != nil {
		return nil, err
	}

	response := &ResponseAuth{
		User:   currentUser,
		Tokens: tokens,
	}

	return response, nil
}

func (a Auth) RefreshToken(token string) (*Tokens, error) {
	sub, err := api.ValidateToken(token, a.cfg.Jwt.RefreshTokenPrivateKey)
	if err != nil {
		return nil, err
	}

	var approveType PayloadToken
	mapstructure.Decode(&sub, &approveType)

	tokens, err := a.generateTokens(approveType.UserID, approveType.IsRemember)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (a Auth) Middleware() gin.HandlerFunc {
	return a.authMiddleware
}

func (a Auth) GetUserID(ctx *gin.Context) string {
	return getMustGetAuth(ctx)
}
