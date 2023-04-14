package usecase

import (
	"electronic_diary/internal/domain/admin"
	"electronic_diary/internal/domain/admin/dto"

	"gorm.io/gorm"
)

type Admin struct {
	db *gorm.DB
}

func NewAdminUseCase(db *gorm.DB) admin.UseCase {
	return &Admin{
		db: db,
	}
}

func (a Admin) Create(dto dto.CreateAdminDTO) (*admin.Model, error) {
	_admin := &admin.Model{
		Login:    dto.Login,
		Password: dto.Password,
	}

	result := a.db.Create(&_admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return _admin, nil
}

func (a Admin) FindAll() []*admin.Model {
	var users []*admin.Model

	a.db.Find(&users)

	return users
}
