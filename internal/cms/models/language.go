package models

import (
	"time"

	"gorm.io/gorm"
)

// Language represents a supported language in the CMS
type Language struct {
	Code       string         `json:"code" gorm:"primaryKey"` // en, es, fr, etc.
	Name       string         `json:"name"`                   // English, Spanish, French
	NativeName string         `json:"native_name"`            // English, Español, Français
	Flag       string         `json:"flag"`                   // Country flag emoji or URL
	IsActive   bool           `json:"is_active" gorm:"default:true"`
	IsRTL      bool           `json:"is_rtl" gorm:"default:false"` // Right-to-left language
	SortOrder  int            `json:"sort_order" gorm:"default:0"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// IsDefault checks if this is the default language
func (l *Language) IsDefault() bool {
	return l.Code == "en" // English is default
}

// GetDisplayName returns the native name for display
func (l *Language) GetDisplayName() string {
	if l.NativeName != "" {
		return l.NativeName
	}
	return l.Name
}

// GetFlagEmoji returns the flag emoji or a default one
func (l *Language) GetFlagEmoji() string {
	if l.Flag != "" {
		return l.Flag
	}

	// Default flags based on language code
	flagMap := map[string]string{
		"en": "🇺🇸",
		"es": "🇪🇸",
		"fr": "🇫🇷",
		"pt": "🇵🇹",
		"de": "🇩🇪",
		"zh": "🇨🇳",
		"ar": "🇸🇦",
	}

	if flag, exists := flagMap[l.Code]; exists {
		return flag
	}

	return "🌐" // Default globe emoji
}

// TableName returns the table name for Language model
func (Language) TableName() string {
	return "languages"
}
