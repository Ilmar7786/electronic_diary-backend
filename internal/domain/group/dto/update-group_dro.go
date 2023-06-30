package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type UpdateGroupDTO struct {
	Title              *string    `json:"title"`
	StudentID          *uuid.UUID `json:"studentId"`
	ClassroomTeacherId *uuid.UUID `json:"classroomTeacherId"`
}

func (d UpdateGroupDTO) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Title, validation.Length(0, 20)),
		validation.Field(&d.StudentID, is.UUID),
		validation.Field(&d.ClassroomTeacherId, is.UUID),
	)
}
