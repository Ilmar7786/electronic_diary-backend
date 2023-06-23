package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"electronic_diary/internal/domain/user"
)

type Model struct {
	ID        uuid.UUID
	Hash      string
	UserID    uuid.UUID
	User      user.Model
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Model) TableName() string {
	return "auth_refresh_tokens"
}

func (m *Model) BeforeCreate(ctx *gorm.DB) (err error) {
	m.ID = uuid.New()

	return nil
}
