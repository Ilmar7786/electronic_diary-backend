package admin

import "electronic_diary/internal/domain/admin/dto"

type UseCase interface {
	Create(dto dto.CreateAdminDTO) (*Model, error)
	FindAll() []*Model
}
