package repository

import "github.com/google/uuid"

type GormAdmin struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Login    string    `gorm:"unique;not null;size:30"`
	Password string    `gorm:"unique;not null;size:30"`
}
