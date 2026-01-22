package models

import (
	shared_models "gcx-cms/internal/shared/models"
	"time"

	"gorm.io/gorm"
)

type MediaFile struct {
	ID           uint               `json:"id" gorm:"primaryKey"`
	OriginalName string             `json:"original_name" gorm:"not null"`
	Filename     string             `json:"filename" gorm:"not null;unique"`
	URL          string             `json:"url" gorm:"not null"`
	ThumbnailURL *string            `json:"thumbnail_url"`
	MimeType     string             `json:"mime_type" gorm:"not null"`
	Size         int64              `json:"size" gorm:"not null"`
	UploadedBy   uint               `json:"uploaded_by" gorm:"not null"`
	User         shared_models.User `json:"user" gorm:"foreignKey:UploadedBy"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	DeletedAt    gorm.DeletedAt     `json:"-" gorm:"index"`
}

// TableName specifies the table name for the MediaFile model
func (MediaFile) TableName() string {
	return "media_files"
}
