package subject

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m *Model) TableName() string {
	return "subjects"
}

func (m *Model) BeforeCreate(ctx *gorm.DB) (err error) {
	m.ID = uuid.New()

	return nil
}
