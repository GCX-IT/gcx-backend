package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// PhotoGallery represents a photo album/collection
type PhotoGallery struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Slug        string         `json:"slug" gorm:"uniqueIndex;not null"`
	Description *string        `json:"description"`
	Category    string         `json:"category" gorm:"not null"`
	CoverImage  *string        `json:"cover_image"`
	Date        *time.Time     `json:"date"`
	Location    *string        `json:"location"`
	Tags        datatypes.JSON `json:"tags" gorm:"type:json"`
	PhotoCount  int            `json:"photo_count" gorm:"default:0"`
	IsFeatured  bool           `json:"is_featured" gorm:"default:false"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	SortOrder   int            `json:"sort_order" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Photos []GalleryPhoto `json:"photos,omitempty" gorm:"foreignKey:GalleryID"`
}

// GalleryPhoto represents an individual photo in a gallery
type GalleryPhoto struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	GalleryID    uint           `json:"gallery_id" gorm:"not null"`
	Title        *string        `json:"title"`
	Description  *string        `json:"description"`
	ImageURL     string         `json:"image_url" gorm:"not null"`
	ThumbnailURL *string        `json:"thumbnail_url"`
	Photographer *string        `json:"photographer"`
	Caption      *string        `json:"caption"`
	Tags         datatypes.JSON `json:"tags" gorm:"type:json"`
	SortOrder    int            `json:"sort_order" gorm:"default:0"`
	IsCover      bool           `json:"is_cover" gorm:"default:false"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Gallery PhotoGallery `json:"gallery,omitempty" gorm:"foreignKey:GalleryID"`
}

// BeforeCreate hook to initialize JSON fields
func (g *PhotoGallery) BeforeCreate(tx *gorm.DB) error {
	if g.Tags == nil {
		g.Tags = datatypes.JSON([]byte("[]"))
	}
	return nil
}

// BeforeCreate hook for GalleryPhoto
func (p *GalleryPhoto) BeforeCreate(tx *gorm.DB) error {
	if p.Tags == nil {
		p.Tags = datatypes.JSON([]byte("[]"))
	}
	return nil
}

// AfterCreate hook to update gallery photo count
func (p *GalleryPhoto) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&PhotoGallery{}).
		Where("id = ?", p.GalleryID).
		UpdateColumn("photo_count", gorm.Expr("photo_count + ?", 1)).Error
}

// AfterDelete hook to update gallery photo count
func (p *GalleryPhoto) AfterDelete(tx *gorm.DB) error {
	return tx.Model(&PhotoGallery{}).
		Where("id = ?", p.GalleryID).
		UpdateColumn("photo_count", gorm.Expr("photo_count - ?", 1)).Error
}

// TableName returns the table name for PhotoGallery model
func (PhotoGallery) TableName() string {
	return "photo_galleries"
}

// TableName returns the table name for GalleryPhoto model
func (GalleryPhoto) TableName() string {
	return "gallery_photos"
}
