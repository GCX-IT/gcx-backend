package models

import (
	"time"

	"gorm.io/gorm"
)

type NewsStatus string
type NewsSource string

const (
	NewsStatusDraft     NewsStatus = "draft"
	NewsStatusPublished NewsStatus = "published"
	NewsStatusArchived  NewsStatus = "archived"
)

const (
	NewsSourceGCX      NewsSource = "gcx"
	NewsSourcePartner  NewsSource = "partner"
	NewsSourceExternal NewsSource = "external"
	NewsSourceAPI      NewsSource = "api"
	NewsSourceFirebase NewsSource = "firebase"
)

// NewsItem represents a news item for the ticker
type NewsItem struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Content     string         `json:"content" gorm:"type:text"`
	Source      NewsSource     `json:"source" gorm:"not null;default:gcx"`
	SourceName  *string        `json:"source_name"`               // For external sources
	SourceURL   *string        `json:"source_url"`                // Link to original source
	Category    *string        `json:"category"`                  // e.g., "market", "announcement", "event"
	Priority    int            `json:"priority" gorm:"default:0"` // Higher numbers = higher priority
	Status      NewsStatus     `json:"status" gorm:"default:draft"`
	IsBreaking  bool           `json:"is_breaking" gorm:"default:false"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	PublishedAt *time.Time     `json:"published_at"`
	ExpiresAt   *time.Time     `json:"expires_at"` // Auto-hide after this time
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// External API fields
	ExternalID   *string    `json:"external_id"`   // ID from external system
	ExternalData *string    `json:"external_data"` // JSON data from external source
	LastSyncAt   *time.Time `json:"last_sync_at"`
}

// NewsCategory represents news categories
type NewsCategory struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Slug        string    `json:"slug" gorm:"type:varchar(191);uniqueIndex;not null"`
	Description *string   `json:"description"`
	Color       *string   `json:"color"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewsSourceConfig represents configuration for external news sources
type NewsSourceConfig struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	SourceName      string     `json:"source_name" gorm:"not null;uniqueIndex"`
	SourceType      NewsSource `json:"source_type" gorm:"not null"`
	IsActive        bool       `json:"is_active" gorm:"default:true"`
	APIEndpoint     *string    `json:"api_endpoint"`
	APIKey          *string    `json:"api_key"`
	RefreshInterval int        `json:"refresh_interval" gorm:"default:300"` // seconds
	LastSyncAt      *time.Time `json:"last_sync_at"`
	Config          *string    `json:"config" gorm:"type:json"` // Additional config as JSON
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

// IsPublished checks if the news item is published and active
func (n *NewsItem) IsPublished() bool {
	if n.Status != NewsStatusPublished || !n.IsActive {
		return false
	}

	// Check if published
	if n.PublishedAt == nil || n.PublishedAt.After(time.Now()) {
		return false
	}

	// Check if expired
	if n.ExpiresAt != nil && n.ExpiresAt.Before(time.Now()) {
		return false
	}

	return true
}

// Publish sets the news item as published
func (n *NewsItem) Publish() {
	n.Status = NewsStatusPublished
	n.IsActive = true
	now := time.Now()
	n.PublishedAt = &now
}

// Unpublish sets the news item as draft
func (n *NewsItem) Unpublish() {
	n.Status = NewsStatusDraft
	n.IsActive = false
}

// Archive sets the news item as archived
func (n *NewsItem) Archive() {
	n.Status = NewsStatusArchived
	n.IsActive = false
}

// SetBreaking sets the news item as breaking news
func (n *NewsItem) SetBreaking(isBreaking bool) {
	n.IsBreaking = isBreaking
	if isBreaking {
		n.Priority = 10 // High priority for breaking news
	}
}

// BeforeCreate hook to set published_at if status is published
func (n *NewsItem) BeforeCreate(tx *gorm.DB) error {
	if n.Status == NewsStatusPublished && n.PublishedAt == nil {
		now := time.Now()
		n.PublishedAt = &now
	}
	return nil
}

// TableName returns the table name for NewsItem model
func (NewsItem) TableName() string {
	return "news_items"
}

// TableName returns the table name for NewsCategory model
func (NewsCategory) TableName() string {
	return "news_categories"
}

// TableName returns the table name for NewsSourceConfig model
func (NewsSourceConfig) TableName() string {
	return "news_source_configs"
}
