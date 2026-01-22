package models

import (
	"time"

	shared_models "gcx-cms/internal/shared/models"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type BlogStatus string

const (
	StatusDraft     BlogStatus = "draft"
	StatusPublished BlogStatus = "published"
	StatusPrivate   BlogStatus = "private"
)

type BlogPost struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Title         string         `json:"title" gorm:"not null"`
	Slug          string         `json:"slug" gorm:"uniqueIndex;not null"`
	Content       string         `json:"content" gorm:"type:longtext"`
	Excerpt       string         `json:"excerpt" gorm:"type:text"`
	FeaturedImage *string        `json:"featured_image"`
	Tags          datatypes.JSON `json:"tags" gorm:"type:json"`
	Status        BlogStatus     `json:"status" gorm:"default:draft"`
	AuthorID      uint           `json:"author_id"`
	PublishedAt   *time.Time     `json:"published_at"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`

	// Relationships
	Author shared_models.User `json:"author" gorm:"foreignKey:AuthorID"`
}

// BlogCategory represents blog categories
type BlogCategory struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Slug        string    `json:"slug" gorm:"uniqueIndex;not null"`
	Description *string   `json:"description"`
	Color       *string   `json:"color"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// IsPublished checks if the blog post is published
func (b *BlogPost) IsPublished() bool {
	return b.Status == StatusPublished && b.PublishedAt != nil && b.PublishedAt.Before(time.Now())
}

// Publish sets the blog post as published
func (b *BlogPost) Publish() {
	b.Status = StatusPublished
	now := time.Now()
	b.PublishedAt = &now
}

// Unpublish sets the blog post as draft
func (b *BlogPost) Unpublish() {
	b.Status = StatusDraft
	b.PublishedAt = nil
}

// BeforeCreate hook to set published_at if status is published
func (b *BlogPost) BeforeCreate(tx *gorm.DB) error {
	if b.Status == StatusPublished && b.PublishedAt == nil {
		now := time.Now()
		b.PublishedAt = &now
	}
	return nil
}

// TableName returns the table name for BlogPost model
func (BlogPost) TableName() string {
	return "blog_posts"
}

// TableName returns the table name for BlogCategory model
func (BlogCategory) TableName() string {
	return "blog_categories"
}
