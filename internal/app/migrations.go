package app

import (
	"electronic_diary/internal/domain/admin"

	"gorm.io/gorm"
)

// runAutoMigrate - AutoMigrate run auto migration for given models
func runAutoMigrate(db *gorm.DB) {
	db.AutoMigrate(admin.Model{})
}
