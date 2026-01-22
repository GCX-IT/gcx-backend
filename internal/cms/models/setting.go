package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Setting represents a configuration setting for the CMS
type Setting struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Key         string         `json:"key" gorm:"uniqueIndex;not null;size:100" binding:"required"`
	Value       string         `json:"value" gorm:"type:text"`
	Type        string         `json:"type" gorm:"size:50;default:'string'"` // string, number, boolean, json, text
	Group       string         `json:"group" gorm:"size:50;index"`           // general, social, seo, etc.
	Label       string         `json:"label" gorm:"size:200"`
	Description string         `json:"description" gorm:"type:text"`
	IsPublic    bool           `json:"is_public" gorm:"default:false"` // Whether this setting can be accessed publicly
	SortOrder   int            `json:"sort_order" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// SettingGroup represents a group of related settings
type SettingGroup struct {
	Name        string    `json:"name"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	Settings    []Setting `json:"settings"`
}

// Common setting types
const (
	SettingTypeString  = "string"
	SettingTypeNumber  = "number"
	SettingTypeBoolean = "boolean"
	SettingTypeJSON    = "json"
	SettingTypeText    = "text"
	SettingTypeURL     = "url"
	SettingTypeEmail   = "email"
	SettingTypeColor   = "color"
	SettingTypeImage   = "image"
)

// Common setting groups
const (
	SettingGroupGeneral = "general"
	SettingGroupSocial  = "social"
	SettingGroupSEO     = "seo"
	SettingGroupContact = "contact"
	SettingGroupFooter  = "footer"
	SettingGroupPages   = "pages"
	SettingGroupBlog    = "blog"
	SettingGroupTheme   = "theme"
)

// TableName returns the table name for the Setting model
func (Setting) TableName() string {
	return "settings"
}

// GetValue returns the setting value with proper type conversion
func (s *Setting) GetValue() interface{} {
	switch s.Type {
	case SettingTypeBoolean:
		return s.Value == "true" || s.Value == "1"
	case SettingTypeNumber:
		// You might want to add proper number parsing here
		return s.Value
	default:
		return s.Value
	}
}

// SetValue sets the setting value with proper type conversion
func (s *Setting) SetValue(value interface{}) {
	switch v := value.(type) {
	case bool:
		if v {
			s.Value = "true"
		} else {
			s.Value = "false"
		}
		s.Type = SettingTypeBoolean
	case string:
		s.Value = v
		if s.Type == "" {
			s.Type = SettingTypeString
		}
	default:
		s.Value = fmt.Sprintf("%v", v)
		if s.Type == "" {
			s.Type = SettingTypeString
		}
	}
}
