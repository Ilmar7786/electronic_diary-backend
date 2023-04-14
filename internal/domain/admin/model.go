package admin

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Login     string    `gorm:"unique;not null;size:30" json:"login"`
	Password  string    `gorm:"unique;not null;size:30" json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m Model) TableName() string {
	return "administrators"
}
