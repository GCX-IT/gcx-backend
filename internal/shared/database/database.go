package database

import (
	"gcx-cms/internal/shared/config"

	"gorm.io/gorm"
)

// GetDB returns the global database instance
func GetDB() *gorm.DB {
	return config.DB
}
