package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// VideoLibrary represents a video category/collection
type VideoLibrary struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Slug        string         `json:"slug" gorm:"uniqueIndex;not null"`
	Description *string        `json:"description"`
	Category    string         `json:"category" gorm:"not null"`
	CoverImage  *string        `json:"cover_image"`
	Date        *time.Time     `json:"date"`
	Location    *string        `json:"location"`
	Tags        datatypes.JSON `json:"tags" gorm:"type:json"`
	VideoCount  int            `json:"video_count" gorm:"default:0"`
	IsFeatured  bool           `json:"is_featured" gorm:"default:false"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	SortOrder   int            `json:"sort_order" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Videos []LibraryVideo `json:"videos,omitempty" gorm:"foreignKey:LibraryID"`
}

// LibraryVideo represents an individual video in a library
type LibraryVideo struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	LibraryID    uint           `json:"library_id" gorm:"not null"`
	Title        string         `json:"title" gorm:"not null"`
	Description  *string        `json:"description"`
	VideoURL     string         `json:"video_url" gorm:"not null"`
	ThumbnailURL *string        `json:"thumbnail_url"`
	Duration     *string        `json:"duration"`
	FileSize     *int64         `json:"file_size"`
	VideoType    string         `json:"video_type" gorm:"default:mp4"`
	Resolution   *string        `json:"resolution"`
	ViewCount    int            `json:"view_count" gorm:"default:0"`
	IsFeatured   bool           `json:"is_featured" gorm:"default:false"`
	IsCover      bool           `json:"is_cover" gorm:"default:false"`
	SortOrder    int            `json:"sort_order" gorm:"default:0"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Library VideoLibrary `json:"library,omitempty" gorm:"foreignKey:LibraryID"`
}

// BeforeCreate hook to initialize JSON fields
func (v *VideoLibrary) BeforeCreate(tx *gorm.DB) error {
	if v.Tags == nil {
		v.Tags = datatypes.JSON([]byte("[]"))
	}
	return nil
}

// AfterCreate hook to update library video count
func (v *LibraryVideo) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&VideoLibrary{}).
		Where("id = ?", v.LibraryID).
		UpdateColumn("video_count", gorm.Expr("video_count + ?", 1)).Error
}

// AfterDelete hook to update library video count
func (v *LibraryVideo) AfterDelete(tx *gorm.DB) error {
	return tx.Model(&VideoLibrary{}).
		Where("id = ?", v.LibraryID).
		UpdateColumn("video_count", gorm.Expr("video_count - ?", 1)).Error
}

// IncrementViewCount increments the view count for a video
func (v *LibraryVideo) IncrementViewCount(db *gorm.DB) error {
	return db.Model(v).UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error
}

// TableName returns the table name for VideoLibrary model
func (VideoLibrary) TableName() string {
	return "video_libraries"
}

// TableName returns the table name for LibraryVideo model
func (LibraryVideo) TableName() string {
	return "library_videos"
}
