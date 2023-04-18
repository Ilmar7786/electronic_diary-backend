package role

import "electronic_diary/internal/domain/role/dto"

type UseCase interface {
	Create(dto dto.CreateRoleDTO) (*Model, error)
	FindAll() []*Model
	FindByID(id string) (*Model, error)
	FindByName(name string) (*Model, error)
	Delete(id string) error
	UpdateById(id string, dto dto.UpdateRoleDTO) error
}
