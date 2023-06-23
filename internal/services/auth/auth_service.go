package auth

import (
	"errors"

	"electronic_diary/internal/config"
	"electronic_diary/internal/domain/user"
	"electronic_diary/internal/services/auth/dto"
	"electronic_diary/pkg/api"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

const keyUserStorageCtx = "userId"

type Auth struct {
	userUC user.UseCase

	cfg config.App
	db  *gorm.DB
}

func New(userUC user.UseCase, cfg config.App, db *gorm.DB) Service {
	return &Auth{
		userUC: userUC,
		cfg:    cfg,

		db: db,
	}
}

func (a Auth) SignIn(dto dto.SignInDTO) (*ResponseAuth, error) {
	currentUser, err := a.userUC.FindByEmail(dto.Email)

	if err != nil || currentUser.ComparePassword(dto.Password) != nil {
		return nil, errors.New("wrong login or password")
	}

	tokens, err := a.generateTokens(currentUser.ID.String(), dto.Remember)
	if err != nil {
		return nil, err
	}

	refreshHash, err := getTokenHash(tokens.Refresh, a.cfg.Jwt.RefreshTokenPrivateKey)
	if err != nil {
		return nil, err
	}
	a.db.Where(&Model{UserID: currentUser.ID}).Delete(&Model{})
	a.db.Create(&Model{Hash: refreshHash, UserID: currentUser.ID})

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

	refreshHash, err := getTokenHash(token, a.cfg.Jwt.RefreshTokenPrivateKey)
	if err != nil {
		return nil, err
	}

	isValidateRefresh := a.db.Find(&Model{Hash: refreshHash})
	if isValidateRefresh.RowsAffected == 0 {
		return nil, errors.New("not a valid token")
	}

	var approveType PayloadToken
	mapstructure.Decode(sub, &approveType)

	tokens, err := a.generateTokens(approveType.UserID, approveType.IsRemember)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(approveType.UserID)
	if err != nil {
		return nil, err
	}

	newRefreshHash, err := getTokenHash(tokens.Refresh, a.cfg.Jwt.RefreshTokenPrivateKey)
	if err != nil {
		return nil, err
	}
	a.db.Where(&Model{UserID: userId}).Delete(&Model{})
	a.db.Create(&Model{Hash: newRefreshHash, UserID: userId})

	return tokens, nil
}

func (a Auth) GetUser(ctx *gin.Context) user.Model {
	return ctx.MustGet(keyUserStorageCtx).(user.Model)
}
