package student

import "electronic_diary/internal/domain/student/dto"

type UseCase interface {
	Create(dto dto.CreateStudentDTO) (*Model, error)
	FindByID(id string) (*Model, error)
	FindAll() []*Model
	Delete(id string) error
	UpdateById(id string, dto dto.UpdateStudentDTO) error
}
