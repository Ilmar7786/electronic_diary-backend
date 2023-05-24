package user

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Model struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Surname    string    `json:"surname" gorm:"not null;size:40"`
	Name       string    `json:"name" gorm:"not null;size:40"`
	Patronymic string    `json:"patronymic" gorm:"not null;size:40"`
	Email      string    `json:"email" gorm:"unique;not null;size:50"`
	IsEmail    bool      `json:"isEmail" gorm:"default:false"`
	Password   string    `json:"-" gorm:"not null;type:text"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (m *Model) TableName() string {
	return "users"
}

// HashPassword substitutes Model.Password with its bcrypt hash
func (m *Model) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	m.Password = string(hash)
	return nil
}

// ComparePassword compares Model.Password hash with raw password
func (m *Model) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
}

// BeforeCreate gorm hook
func (m *Model) BeforeCreate(ctx *gorm.DB) (err error) {
	return m.HashPassword()
}
