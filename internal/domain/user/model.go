package user

import (
	"time"

	"electronic_diary/internal/constants"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Model struct {
	ID          uuid.UUID      `json:"id"`
	Surname     string         `json:"surname"`
	Name        string         `json:"name"`
	Patronymic  string         `json:"patronymic"`
	Address     string         `json:"address"`
	Phone       string         `json:"phone"`
	Email       string         `json:"email"`
	Password    string         `json:"-"`
	Role        constants.Role `json:"role" enums:"student,teacher,parent"`
	IsActive    bool           `json:"isActive"`
	IsSuperUser bool           `json:"isSuperUser"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
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

func (m *Model) BeforeCreate(ctx *gorm.DB) (err error) {
	if err = m.HashPassword(); err != nil {
		return err
	}
	m.ID = uuid.New()

	return nil
}
