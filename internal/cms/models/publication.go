package models

import "time"

// Publication represents a publication in the system
type Publication struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	Title           string     `json:"title" gorm:"not null"`
	Description     string     `json:"description" gorm:"type:text"`
	Category        string     `json:"category" gorm:"type:enum('Research Papers','Annual Reports','Policy Documents');not null"`
	ImagePath       string     `json:"image_path" gorm:"size:500"` // Publication cover/thumbnail image
	FilePath        string     `json:"file_path" gorm:"size:500"`  // PDF document file
	FileName        string     `json:"file_name" gorm:"size:255"`
	FileSize        int        `json:"file_size"`
	FileType        string     `json:"file_type" gorm:"size:100"`
	PublicationDate *time.Time `json:"publication_date"`
	Author          string     `json:"author" gorm:"size:255"`
	Tags            string     `json:"tags" gorm:"type:text"`
	Status          string     `json:"status" gorm:"type:enum('Published','Draft','Archived');default:'Published'"`
	DownloadCount   int        `json:"download_count" gorm:"default:0"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
