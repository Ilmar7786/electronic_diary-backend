package usecase

import (
	"errors"

	"electronic_diary/internal/domain/teacher"
	"electronic_diary/internal/domain/teacher/dto"
	"electronic_diary/internal/domain/user"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type Teacher struct {
	db     *gorm.DB
	userUC user.UseCase
}

func New(db *gorm.DB, userUC user.UseCase) teacher.UseCase {
	return &Teacher{db: db, userUC: userUC}
}

func (t Teacher) Create(dto dto.CreateTeacherDTO) (*teacher.Model, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	newUser, err := t.userUC.Create(dto.User)
	if err != nil {
		return nil, err
	}

	newParent := teacher.Model{
		UserID:  newUser.ID,
		Subject: dto.Subject,
	}

	err = t.db.Create(&newParent).Preload("User").Preload("Subject").Find(&newParent).Error
	if err != nil {
		return nil, err
	}

	return &newParent, nil
}

func (t Teacher) FindByID(id string) (*teacher.Model, error) {
	idParse, _ := uuid.Parse(id)
	var candidate teacher.Model

	err := t.db.First(&candidate, idParse).Error
	if err != nil {
		return nil, errors.New("nothing found")
	}

	return &candidate, nil
}

func (t Teacher) FindAll() []*teacher.Model {
	teachers := make([]*teacher.Model, 0)
	t.db.Find(&teachers)

	return teachers
}

func (t Teacher) Delete(id string) error {
	candidate, err := t.FindByID(id)

	if err != nil {
		return err
	}

	return t.db.Delete(&candidate).Error
}

func (t Teacher) UpdateById(id string, dto dto.UpdateTeacherDTO) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	candidate, err := t.FindByID(id)
	if err != nil {
		return err
	}

	if err := mapstructure.Decode(dto, &candidate); err != nil {
		return err
	}

	if err := t.db.Updates(&candidate).Error; err != nil {
		return err
	}

	return nil
}
