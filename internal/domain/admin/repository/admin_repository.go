package repository

import (
	"electronic_diary/internal/domain/admin"

	"gorm.io/gorm"
)

type Admin struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) admin.UseCase {
	return &Admin{db: db}
}

func (a Admin) Create() (*admin.Model, error) {
	//TODO implement me
	panic("implement me")
}

func (a Admin) FindById(id string) (*admin.Model, error) {
	//TODO implement me
	panic("implement me")
}

func (a Admin) FindAll() []*admin.Model {
	//TODO implement me
	panic("implement me")
}

func (a Admin) Update() (*admin.Model, error) {
	//TODO implement me
	panic("implement me")
}

func (a Admin) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
