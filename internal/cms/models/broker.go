package models

import "time"

// Broker represents a broker in the system
type Broker struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"not null"`
	Company         string    `json:"company" gorm:"size:255"`
	LicenseNumber   string    `json:"license_number" gorm:"size:100"`
	PhoneNo         string    `json:"phone_no" gorm:"size:50"`
	Email           string    `json:"email" gorm:"size:255"`
	Address         string    `json:"address" gorm:"type:text"`
	Specialization  string    `json:"specialization" gorm:"size:255"`
	ExperienceYears int       `json:"experience_years" gorm:"default:0"`
	Status          string    `json:"status" gorm:"type:enum('Active','Inactive','Suspended');default:'Active'"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
