package teacher

import "electronic_diary/internal/domain/teacher/dto"

type UseCase interface {
	Create(dto dto.CreateTeacherDTO) (*Model, error)
	FindByID(id string) (*Model, error)
	FindAll() []*Model
	Delete(id string) error
	UpdateById(id string, dto dto.UpdateTeacherDTO) error
}
