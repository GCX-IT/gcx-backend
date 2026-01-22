package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Translation represents a translation for any CMS resource
type Translation struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	ResourceType string         `json:"resource_type"` // page, post, setting, menu_item, etc.
	ResourceID   uint           `json:"resource_id"`
	LanguageCode string         `json:"language_code"` // en, es, fr, etc.
	FieldName    string         `json:"field_name"`    // title, content, description, label, etc.
	Content      string         `json:"content" gorm:"type:longtext"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Indexes for performance
	// gorm:"index:idx_resource,unique" - composite unique index
	// gorm:"index:idx_language" - language index
}

// GetTranslationKey returns a unique key for this translation
func (t *Translation) GetTranslationKey() string {
	return fmt.Sprintf("%s:%d:%s:%s", t.ResourceType, t.ResourceID, t.LanguageCode, t.FieldName)
}

// TableName returns the table name for Translation model
func (Translation) TableName() string {
	return "translations"
}

// BeforeCreate hook to ensure unique translations
func (t *Translation) BeforeCreate(tx *gorm.DB) error {
	// Check if translation already exists
	var existing Translation
	err := tx.Where("resource_type = ? AND resource_id = ? AND language_code = ? AND field_name = ?",
		t.ResourceType, t.ResourceID, t.LanguageCode, t.FieldName).
		First(&existing).Error

	if err == nil {
		// Translation exists, update it instead
		existing.Content = t.Content
		return tx.Save(&existing).Error
	}

	return nil
}
