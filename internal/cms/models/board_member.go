package models

import (
	"time"
	"gorm.io/gorm"
)

// BoardMember represents a board member record
type BoardMember struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Position    string         `json:"position" gorm:"not null"`
	Image       string         `json:"image"`
	Description string         `json:"description" gorm:"type:text;not null"`
	LinkedIn    string         `json:"linkedin"`
	Facebook    string         `json:"facebook"`
	Instagram   string         `json:"instagram"`
	OrderIndex  int            `json:"order_index" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName returns the table name for BoardMember
func (BoardMember) TableName() string {
	return "board_members"
}
