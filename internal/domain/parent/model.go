package parent

import (
	"time"

	"electronic_diary/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID  `json:"id"`
	Guardian  string     `json:"guardian"`
	UserID    uuid.UUID  `json:"userId"`
	User      user.Model `json:"user"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

func (m *Model) TableName() string {
	return "parents"
}

func (m *Model) BeforeCreate(ctx *gorm.DB) (err error) {
	m.ID = uuid.New()

	return nil
}
