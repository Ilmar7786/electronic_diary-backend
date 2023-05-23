package role

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"unique;not null;size:50"`
	Key       string    `json:"key" gorm:"unique;not null;size:50"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m Model) TableName() string {
	return "roles"
}
