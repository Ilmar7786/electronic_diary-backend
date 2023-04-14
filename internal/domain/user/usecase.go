package user

import "electronic_diary/internal/domain/user/dto"

type UseCase interface {
	Create(dto dto.CreateUserDTO) (*Model, error)
	FindById(id string) (*Model, error)
	FindByEmail(email string) (*Model, error)
	FindAll() []*Model
	DeleteById(id string) error
	UpdateById(id string, dto dto.UpdateUserDTO) error
}
