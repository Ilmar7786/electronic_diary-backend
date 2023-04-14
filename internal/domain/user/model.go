package user

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Surname    string    `json:"surname" gorm:"not null;size:40"`
	Name       string    `json:"name" gorm:"not null;size:40"`
	Patronymic string    `json:"patronymic" gorm:"not null;size:40"`
	Email      string    `json:"email" gorm:"unique;not null;size:50"`
	IsEmail    bool      `json:"isEmail" gorm:"default:false"`
	Password   string    `json:"-" gorm:"not null;size:30"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (m *Model) TableName() string {
	return "users"
}
