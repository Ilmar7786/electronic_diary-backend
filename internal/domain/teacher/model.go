package teacher

import (
	"time"

	"electronic_diary/internal/domain/subject"
	"electronic_diary/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID        `json:"id"`
	UserID    uuid.UUID        `json:"userId"`
	User      user.Model       `json:"user"`
	Subject   []*subject.Model `json:"subject"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}

func (m *Model) TableName() string {
	return "teachers"
}

func (m *Model) BeforeCreate(ctx *gorm.DB) (err error) {
	m.ID = uuid.New()

	return nil
}
