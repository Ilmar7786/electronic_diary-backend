package parent

import "electronic_diary/internal/domain/parent/dto"

type UseCase interface {
	Create(dto dto.CreateParentDTO) (*Model, error)
	FindByID(id string) (*Model, error)
	FindAll() []*Model
	Delete(id string) error
	UpdateById(id string, dto dto.UpdateParentDTO) error
}
