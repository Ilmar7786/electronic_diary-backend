package group

import (
	"time"

	"electronic_diary/internal/domain/student"
	"electronic_diary/internal/domain/teacher"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID                 uuid.UUID     `json:"id"`
	Title              string        `json:"title"`
	StudentID          uuid.UUID     `json:"studentId"`
	ClassroomTeacherId uuid.UUID     `json:"classroomTeacherId"`
	ClassroomTeacher   teacher.Model `json:"classroomTeacher"`
	Student            student.Model `json:"student"`
	CreatedAt          time.Time     `json:"createdAt"`
	UpdatedAt          time.Time     `json:"updatedAt"`
}

func (m *Model) TableName() string {
	return "groups"
}

func (m *Model) BeforeCreate(ctx *gorm.DB) (err error) {
	m.ID = uuid.New()

	return nil
}
