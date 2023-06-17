package user

import (
	"errors"
	"fmt"
	"time"

	"electronic_diary/internal/constants"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Model struct {
	ID          uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Surname     string         `json:"surname" gorm:"not null;size:40"`
	Name        string         `json:"name" gorm:"not null;size:40"`
	Patronymic  string         `json:"patronymic" gorm:"not null;size:40"`
	Address     string         `json:"address" gorm:"not null"`
	Phone       string         `json:"phone" gorm:"not null"`
	Email       string         `json:"email" gorm:"unique;not null;size:50"`
	Password    string         `json:"-" gorm:"not null;type:text"`
	Role        constants.Role `json:"role" gorm:"type:role_enum;not null" enums:"STUDENT,TEACHER,PARENT"`
	IsActive    bool           `json:"isActive" gorm:"default:false"`
	IsSuperUser bool           `json:"isSuperUser" gorm:"not null"`
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
	err = validation.Validate(m.Role, validation.In(
		constants.TeacherRole,
		constants.ParentRole,
		constants.StudentRole,
	))

	if err != nil {
		return errors.New(fmt.Sprintf("role: invalid value for enum: %s", m.Role))
	}

	if err = m.HashPassword(); err != nil {
		return err
	}

	return nil
}
