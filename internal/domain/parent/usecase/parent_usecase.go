package usecase

import (
	"errors"

	"electronic_diary/internal/domain/parent"
	"electronic_diary/internal/domain/parent/dto"
	"electronic_diary/internal/domain/user"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type Parent struct {
	db     *gorm.DB
	userUC user.UseCase
}

func New(db *gorm.DB, userUC user.UseCase) parent.UseCase {
	return &Parent{db: db, userUC: userUC}
}

func (p Parent) Create(dto dto.CreateParentDTO) (*parent.Model, error) {
	newUser, err := p.userUC.Create(dto.User)
	if err != nil {
		return nil, err
	}

	newParent := parent.Model{
		Guardian: dto.Guardian,
		UserID:   newUser.ID,
	}

	err = p.db.Create(&newParent).Preload("User").Find(&newParent).Error
	if err != nil {
		return nil, err
	}

	return &newParent, nil
}

func (p Parent) FindByID(id string) (*parent.Model, error) {
	idParse, _ := uuid.Parse(id)
	var candidate parent.Model

	err := p.db.First(&candidate, idParse).Error
	if err != nil {
		return nil, errors.New("nothing found")
	}

	return &candidate, nil
}

func (p Parent) FindAll() []*parent.Model {
	parents := make([]*parent.Model, 0)
	p.db.Find(&parents)

	return parents
}

func (p Parent) Delete(id string) error {
	candidate, err := p.FindByID(id)

	if err != nil {
		return err
	}

	return p.db.Delete(&candidate).Error
}

func (p Parent) UpdateById(id string, dto dto.UpdateParentDTO) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	candidate, err := p.FindByID(id)
	if err != nil {
		return err
	}

	if err := mapstructure.Decode(dto, &candidate); err != nil {
		return err
	}

	if err := p.db.Updates(&candidate).Error; err != nil {
		return err
	}

	return nil
}
