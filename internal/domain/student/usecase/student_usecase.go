package usecase

import (
	"errors"

	"electronic_diary/internal/domain/student"
	"electronic_diary/internal/domain/student/dto"
	"electronic_diary/internal/domain/user"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type Student struct {
	db     *gorm.DB
	userUC user.UseCase
}

func New(db *gorm.DB, userUC user.UseCase) student.UseCase {
	return &Student{db: db, userUC: userUC}
}

func (s Student) Create(dto dto.CreateStudentDTO) (*student.Model, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	newUser, err := s.userUC.Create(dto.User)
	if err != nil {
		return nil, err
	}

	newStudent := student.Model{
		UserID:   dto.UserID,
		ParentID: dto.ParentID,
	}

	if err := s.db.Create(&newUser).Preload("User").Preload("Parent").Error; err != nil {
		return nil, err
	}

	return &newStudent, nil
}

func (s Student) FindByID(id string) (*student.Model, error) {
	idParse, _ := uuid.Parse(id)
	var candidate student.Model

	err := s.db.First(&candidate, idParse).Error
	if err != nil {
		return nil, errors.New("student not found")
	}

	return &candidate, nil
}

func (s Student) FindAll() []*student.Model {
	users := make([]*student.Model, 0)
	s.db.Find(&users)

	return users
}

func (s Student) Delete(id string) error {
	candidate, err := s.FindByID(id)

	if err != nil {
		return err
	}

	return s.db.Delete(&candidate).Error
}

func (s Student) UpdateById(id string, dto dto.UpdateStudentDTO) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	candidate, err := s.FindByID(id)
	if err != nil {
		return err
	}

	if err := mapstructure.Decode(dto, &candidate); err != nil {
		return err
	}

	if err := s.db.Updates(&candidate).Error; err != nil {
		return err
	}

	return nil
}
