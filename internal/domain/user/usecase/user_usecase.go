package usecase

import (
	"errors"

	"electronic_diary/internal/domain/user"
	"electronic_diary/internal/domain/user/dto"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UseCase {
	return &User{db: db}
}

func (u User) Create(dto dto.CreateUserDTO) (*user.Model, error) {
	candidate, _ := u.FindByEmail(dto.Email)
	if candidate != nil {
		return nil, errors.New(userExistsEmailError)
	}

	var newUser user.Model

	if err := mapstructure.Decode(dto, &newUser); err != nil {
		return nil, err
	}

	if err := u.db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (u User) FindByID(id string) (*user.Model, error) {
	idParse, _ := uuid.Parse(id)
	var candidate user.Model

	err := u.db.First(&candidate, idParse).Error
	if err != nil {
		return nil, errors.New(userNotFoundError)
	}

	return &candidate, nil
}
func (u User) FindByEmail(email string) (*user.Model, error) {
	var candidate *user.Model
	err := u.db.Where("email = ?", email).First(&candidate).Error

	if err != nil {
		return nil, errors.New(userNotFoundError)
	}

	return candidate, nil
}

func (u User) FindAll() []*user.Model {
	users := make([]*user.Model, 0)
	u.db.Find(&users)

	return users
}

func (u User) Delete(id string) error {
	candidate, err := u.FindByID(id)

	if err != nil {
		return err
	}

	return u.db.Delete(&candidate).Error
}

func (u User) UpdateById(id string, dto dto.UpdateUserDTO) error {
	candidate, err := u.FindByID(id)
	if err != nil {
		return err
	}

	if dto.Email != nil {
		exist, _ := u.FindByEmail(*dto.Email)
		if exist != nil {
			if candidate.Email != *dto.Email {
				return errors.New(userExistsEmailError)
			}
		}
	}

	if err := mapstructure.Decode(dto, &candidate); err != nil {
		return err
	}

	if err := u.db.Updates(&candidate).Error; err != nil {
		return err
	}

	return nil
}
