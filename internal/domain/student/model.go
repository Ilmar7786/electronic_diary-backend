package student

import (
	"time"

	"electronic_diary/internal/domain/parent"
	"electronic_diary/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID                 uuid.UUID    `json:"id"`
	ResidentialAddress string       `json:"residentialAddress"`
	UserID             uuid.UUID    `json:"userId"`
	ParentID           uuid.UUID    `json:"parentId"`
	User               user.Model   `json:"user"`
	Parent             parent.Model `json:"parent"`
	CreatedAt          time.Time    `json:"createdAt"`
	UpdatedAt          time.Time    `json:"updatedAt"`
}

func (m *Model) TableName() string {
	return "students"
}

func (m *Model) BeforeCreate(ctx *gorm.DB) (err error) {
	m.ID = uuid.New()

	return nil
}
