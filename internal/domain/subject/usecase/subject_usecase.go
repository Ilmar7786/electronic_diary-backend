package usecase

import (
	"errors"

	"electronic_diary/internal/domain/subject"
	"electronic_diary/internal/domain/subject/dto"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type Subject struct {
	db *gorm.DB
}

func New(db *gorm.DB) subject.UseCase {
	return &Subject{db: db}
}

func (s Subject) Create(dto dto.CreateSubjectDTO) (*subject.Model, error) {
	var instance subject.Model

	if err := mapstructure.Decode(dto, &instance); err != nil {
		return nil, err
	}

	affected := s.db.Where(subject.Model{Title: dto.Title}).RowsAffected
	if affected != 0 {
		return nil, errors.New("subject with this title already exists")
	}

	if err := s.db.Create(&instance).Error; err != nil {
		return nil, err
	}

	return &instance, nil
}

func (s Subject) FindByID(id string) (*subject.Model, error) {
	idParse, _ := uuid.Parse(id)
	var instance subject.Model

	err := s.db.First(&instance, idParse).Error
	if err != nil {
		return nil, errors.New("subject not found")
	}

	return &instance, nil
}

func (s Subject) FindAll() []*subject.Model {
	subjects := make([]*subject.Model, 0)
	s.db.Find(&subjects)

	return subjects
}

func (s Subject) Delete(id string) error {
	candidate, err := s.FindByID(id)

	if err != nil {
		return err
	}

	return s.db.Delete(&candidate).Error
}

func (s Subject) UpdateById(id string, dto dto.UpdateSubjectDTO) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	candidate, err := s.FindByID(id)
	if err != nil {
		return err
	}

	affected := s.db.Where("title = ?", dto.Title).RowsAffected
	if affected != 0 {
		return errors.New("subject with this title exists")
	}

	if err = mapstructure.Decode(dto, &candidate); err != nil {
		return err
	}

	if err = s.db.Updates(&candidate).Error; err != nil {
		return err
	}

	return nil
}
