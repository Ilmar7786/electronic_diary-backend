package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type CreateGroupDTO struct {
	Title              string    `json:"title"`
	StudentID          uuid.UUID `json:"studentId"`
	ClassroomTeacherId uuid.UUID `json:"classroomTeacherId"`
}

func (d CreateGroupDTO) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Title, validation.Required, validation.Length(0, 20)),
		validation.Field(&d.StudentID, validation.Required, is.UUID),
		validation.Field(&d.ClassroomTeacherId, validation.Required, is.UUID),
	)
}
