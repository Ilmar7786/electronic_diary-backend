package group

import "electronic_diary/internal/domain/group/dto"

type UseCase interface {
	Create(dto dto.CreateGroupDTO) (*Model, error)
	FindByID(id string) (*Model, error)
	FindAll() []*Model
	Delete(id string) error
	UpdateById(id string, dto dto.UpdateGroupDTO) error
}
