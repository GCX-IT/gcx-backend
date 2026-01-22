package models

import (
	"time"
)

// TeamMember represents a team member in the database
type TeamMember struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"type:text"`
	Image       string `json:"image"`
	Type        string `json:"type" gorm:"not null;index"` // 'board', 'executive', 'functional'
	OrderIndex  int    `json:"order_index" gorm:"default:0"`
	// Social Media Handles
	LinkedInURL  string    `json:"linkedin_url" gorm:"size:500"`
	TwitterURL   string    `json:"twitter_url" gorm:"size:500"`
	FacebookURL  string    `json:"facebook_url" gorm:"size:500"`
	InstagramURL string    `json:"instagram_url" gorm:"size:500"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// TableName returns the table name for TeamMember
func (TeamMember) TableName() string {
	return "team_members"
}
