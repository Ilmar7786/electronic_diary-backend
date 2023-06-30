package usecase

import (
	"errors"

	"electronic_diary/internal/domain/group"
	"electronic_diary/internal/domain/group/dto"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type Group struct {
	db *gorm.DB
}

func New(db *gorm.DB) group.UseCase {
	return &Group{db: db}
}

func (p Group) Create(dto dto.CreateGroupDTO) (*group.Model, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	var newGroup group.Model
	if err := mapstructure.Decode(dto, &newGroup); err != nil {
		return nil, err
	}

	err := p.db.Create(&newGroup).Preload("ClassroomTeacher").Preload("Student").Find(&newGroup).Error
	if err != nil {
		return nil, err
	}

	return &newGroup, nil
}

func (p Group) FindByID(id string) (*group.Model, error) {
	idParse, _ := uuid.Parse(id)
	var candidate group.Model

	err := p.db.First(&candidate, idParse).Error
	if err != nil {
		return nil, errors.New("nothing found")
	}

	return &candidate, nil
}

func (p Group) FindAll() []*group.Model {
	parents := make([]*group.Model, 0)
	p.db.Find(&parents)

	return parents
}

func (p Group) Delete(id string) error {
	candidate, err := p.FindByID(id)

	if err != nil {
		return err
	}

	return p.db.Delete(&candidate).Error
}

func (p Group) UpdateById(id string, dto dto.UpdateGroupDTO) error {
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
