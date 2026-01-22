package models

import (
	"time"

	shared_models "gcx-cms/internal/shared/models"

	"gorm.io/gorm"
)

type PageStatus string

const (
	PageStatusDraft     PageStatus = "draft"
	PageStatusPublished PageStatus = "published"
	PageStatusPrivate   PageStatus = "private"
	PageStatusArchived  PageStatus = "archived"
)

// Page represents a CMS-managed page
type Page struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Title           string         `json:"title" gorm:"not null"`
	Slug            string         `json:"slug" gorm:"uniqueIndex;not null"`
	Content         string         `json:"content" gorm:"type:longtext"`
	Excerpt         string         `json:"excerpt" gorm:"type:text"`
	Template        string         `json:"template" gorm:"default:default"`
	Status          PageStatus     `json:"status" gorm:"default:draft"`
	FeaturedImage   *string        `json:"featured_image"`
	MetaTitle       string         `json:"meta_title"`
	MetaDescription string         `json:"meta_description"`
	MetaKeywords    string         `json:"meta_keywords"`
	ParentID        *uint          `json:"parent_id"`
	SortOrder       int            `json:"sort_order" gorm:"default:0"`
	AuthorID        uint           `json:"author_id"`
	PublishedAt     *time.Time     `json:"published_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Author       shared_models.User `json:"author" gorm:"foreignKey:AuthorID"`
	Parent       *Page              `json:"parent" gorm:"foreignKey:ParentID"`
	Children     []Page             `json:"children" gorm:"foreignKey:ParentID"`
	Translations []Translation      `json:"translations" gorm:"foreignKey:ResourceID;constraint:OnDelete:CASCADE"`
}

// IsPublished checks if the page is published
func (p *Page) IsPublished() bool {
	return p.Status == PageStatusPublished && p.PublishedAt != nil && p.PublishedAt.Before(time.Now())
}

// Publish sets the page as published
func (p *Page) Publish() {
	p.Status = PageStatusPublished
	now := time.Now()
	p.PublishedAt = &now
}

// Unpublish sets the page as draft
func (p *Page) Unpublish() {
	p.Status = PageStatusDraft
	p.PublishedAt = nil
}

// Archive sets the page as archived
func (p *Page) Archive() {
	p.Status = PageStatusArchived
}

// BeforeCreate hook to set published_at if status is published
func (p *Page) BeforeCreate(tx *gorm.DB) error {
	if p.Status == PageStatusPublished && p.PublishedAt == nil {
		now := time.Now()
		p.PublishedAt = &now
	}
	return nil
}

// TableName returns the table name for Page model
func (Page) TableName() string {
	return "pages"
}
