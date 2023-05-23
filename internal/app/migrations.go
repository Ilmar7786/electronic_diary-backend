package app

import (
	"electronic_diary/internal/domain/role"
	"electronic_diary/internal/domain/user"

	"gorm.io/gorm"
)

// runAutoMigrate - AutoMigrate run auto migration for given models
func runAutoMigrate(db *gorm.DB) {
	models := []interface{}{
		&role.Model{},
		&user.Model{},
	}

	db.AutoMigrate(models...)
}
