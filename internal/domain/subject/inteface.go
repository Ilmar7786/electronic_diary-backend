package subject

import "electronic_diary/internal/domain/subject/dto"

type UseCase interface {
	Create(dto dto.CreateSubjectDTO) (*Model, error)
	FindByID(id string) (*Model, error)
	FindAll() []*Model
	Delete(id string) error
	UpdateById(id string, dto dto.UpdateSubjectDTO) error
}
