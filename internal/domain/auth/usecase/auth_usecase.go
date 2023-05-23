package usecase

import (
	"errors"

	"electronic_diary/internal/domain/auth"
	"electronic_diary/internal/domain/auth/dto"
	"electronic_diary/internal/domain/user"
	userDTO "electronic_diary/internal/domain/user/dto"

	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type Auth struct {
	db     *gorm.DB
	userUC user.UseCase
}

func NewAuth(db *gorm.DB, userUC user.UseCase) auth.UseCase {
	return &Auth{db, userUC}
}

func (a Auth) SignIn(dto dto.SignInDTO) (*user.Model, error) {
	candidateUser, err := a.userUC.FindByEmail(dto.Email)

	if err != nil || candidateUser.Password != dto.Password {
		return nil, errors.New("not validation email or password")
	}

	return candidateUser, nil
}

func (a Auth) SignUp(dto dto.SignUpDTO) (*user.Model, error) {
	var authDto userDTO.CreateUserDTO

	if err := mapstructure.Decode(dto, &authDto); err != nil {
		return nil, err
	}

	newUser, err := a.userUC.Create(authDto)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (a Auth) LogOut() {
	//TODO implement me
	panic("implement me")
}
