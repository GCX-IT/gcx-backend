package models

import (
	"time"

	"gorm.io/gorm"
)

type Partner struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Category    string         `json:"category" gorm:"not null"` // financial, donor, government, ngo, private, tender
	Logo        string         `json:"logo"`                     // URL to logo image
	Website     string         `json:"website"`
	Email       string         `json:"email"`
	Phone       string         `json:"phone"`
	Address     string         `json:"address"`
	Status      string         `json:"status" gorm:"default:'active'"` // active, inactive, suspended
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// TableName returns the table name for the Partner model
func (Partner) TableName() string {
	return "partners"
}
