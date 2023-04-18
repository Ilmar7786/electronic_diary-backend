package usecase

import (
	"errors"

	"electronic_diary/internal/domain/role"
	"electronic_diary/internal/domain/role/dto"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type Role struct {
	db *gorm.DB
}

func NewRoleUseCase(db *gorm.DB) role.UseCase {
	return &Role{db: db}
}

func (r Role) Create(dto dto.CreateRoleDTO) (*role.Model, error) {
	candidate, _ := r.FindByName(dto.Name)
	if candidate != nil {
		return nil, errors.New(roleExistsEmailError)
	}

	var newRole *role.Model

	if err := mapstructure.Decode(dto, &newRole); err != nil {
		return nil, err
	}

	if err := r.db.Create(&newRole).Error; err != nil {
		return nil, err
	}

	return newRole, nil
}

func (r Role) FindAll() []*role.Model {
	roles := make([]*role.Model, 0)
	r.db.Find(&roles)

	return roles
}

func (r Role) FindByID(id string) (*role.Model, error) {
	idParse, _ := uuid.Parse(id)
	var candidate role.Model

	err := r.db.First(&candidate, idParse).Error
	if err != nil {
		return nil, errors.New(roleNotFoundError)
	}

	return &candidate, nil
}

func (r Role) FindByName(name string) (*role.Model, error) {
	var candidate *role.Model

	err := r.db.Where("name = ?", name).First(&candidate).Error
	if err != nil {
		return nil, errors.New(roleNotFoundError)
	}

	return candidate, nil
}

func (r Role) Delete(id string) error {
	candidate, err := r.FindByID(id)

	if err != nil {
		return err
	}

	return r.db.Delete(&candidate).Error
}

func (r Role) UpdateById(id string, dto dto.UpdateRoleDTO) error {
	candidate, err := r.FindByID(id)
	if err != nil {
		return err
	}

	exist, _ := r.FindByName(dto.Name)
	if exist != nil {
		if candidate.Name != dto.Name {
			return errors.New(roleExistsEmailError)
		}
	}

	if err := mapstructure.Decode(dto, &candidate); err != nil {
		return err
	}

	if err := r.db.Save(&candidate).Error; err != nil {
		return err
	}

	return nil
}
