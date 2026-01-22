package models

import "time"

// Career represents a career opportunity in the system
type Career struct {
	ID                  uint       `json:"id" gorm:"primaryKey"`
	Title               string     `json:"title" gorm:"not null"`
	Description         string     `json:"description" gorm:"type:text"`
	Category            string     `json:"category" gorm:"type:enum('Job Openings','Internship','Job Functional Areas');not null"`
	Department          string     `json:"department" gorm:"size:255"`
	Location            string     `json:"location" gorm:"size:255"`
	EmploymentType      string     `json:"employment_type" gorm:"type:enum('Full-time','Part-time','Contract','Internship');not null"`
	ExperienceLevel     string     `json:"experience_level" gorm:"type:enum('Entry Level','Mid Level','Senior Level','Executive');not null"`
	Requirements        string     `json:"requirements" gorm:"type:text"`
	Responsibilities    string     `json:"responsibilities" gorm:"type:text"`
	Benefits            string     `json:"benefits" gorm:"type:text"`
	SalaryRange         string     `json:"salary_range" gorm:"size:100"`
	ApplicationDeadline *time.Time `json:"application_deadline"`
	StartDate           *time.Time `json:"start_date"`
	Status              string     `json:"status" gorm:"type:enum('Open','Closed','On Hold');default:'Open'"`
	ApplicationCount    int        `json:"application_count" gorm:"default:0"`
	ImagePath           string     `json:"image_path" gorm:"size:500"` // Career flyer image
	ContactEmail        string     `json:"contact_email" gorm:"size:255"` // Contact email for applications
	FilePath            string     `json:"file_path" gorm:"size:500"`
	FileName            string     `json:"file_name" gorm:"size:255"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}
