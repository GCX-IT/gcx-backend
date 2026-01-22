package models

import "time"

// Trader represents a trader/member in the membership list
type Trader struct {
	ID               uint       `json:"id" gorm:"primaryKey"`
	Name             string     `json:"name" gorm:"not null"`
	Industry         string     `json:"industry" gorm:"size:255"`
	MemberType       string     `json:"member_type" gorm:"type:enum('Associates','Full Members','Brokers','Warehouse Operators');default:'Associates'"`
	PhoneNo          string     `json:"phone_no" gorm:"size:50"`
	Email            string     `json:"email" gorm:"size:255"`
	Address          string     `json:"address" gorm:"type:text"`
	RegistrationDate *time.Time `json:"registration_date"`
	Status           string     `json:"status" gorm:"type:enum('Active','Inactive','Suspended');default:'Active'"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}
