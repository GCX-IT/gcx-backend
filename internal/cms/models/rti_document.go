package models

import (
	"time"

	"gorm.io/gorm"
)

// RTIDocument represents a downloadable RTI resource
type RTIDocument struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Title         string         `json:"title" gorm:"not null"`
	Description   *string        `json:"description"`
	Category      string         `json:"category" gorm:"not null"` // manual, form, guide, policy, other
	FilePath      string         `json:"file_path" gorm:"not null"`
	FileName      *string        `json:"file_name"`
	FileSize      *int64         `json:"file_size"`
	Icon          string         `json:"icon" gorm:"default:'pi-file-pdf'"`
	DownloadCount int            `json:"download_count" gorm:"default:0"`
	IsFeatured    bool           `json:"is_featured" gorm:"default:false"`
	IsActive      bool           `json:"is_active" gorm:"default:true"`
	SortOrder     int            `json:"sort_order" gorm:"default:0"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

// IncrementDownloadCount increments the download counter
func (d *RTIDocument) IncrementDownloadCount(db *gorm.DB) error {
	return db.Model(d).UpdateColumn("download_count", gorm.Expr("download_count + ?", 1)).Error
}

// TableName returns the table name for RTIDocument model
func (RTIDocument) TableName() string {
	return "rti_documents"
}
