package models

import (
	"time"

	"gorm.io/gorm"
)

// Menu represents a navigation menu
type Menu struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Location  string         `json:"location"` // header, footer, sidebar, mobile
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Items []MenuItem `json:"items" gorm:"foreignKey:MenuID;constraint:OnDelete:CASCADE"`
}

// MenuItem represents a single menu item
type MenuItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	MenuID    uint           `json:"menu_id"`
	Label     string         `json:"label" gorm:"not null"`
	URL       string         `json:"url"`
	Target    string         `json:"target" gorm:"default:_self"` // _self, _blank, _parent, _top
	IconClass string         `json:"icon_class"`
	ParentID  *uint          `json:"parent_id"`
	SortOrder int            `json:"sort_order" gorm:"default:0"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Menu         Menu          `json:"menu" gorm:"foreignKey:MenuID"`
	Parent       *MenuItem     `json:"parent" gorm:"foreignKey:ParentID"`
	Children     []MenuItem    `json:"children" gorm:"foreignKey:ParentID"`
	Translations []Translation `json:"translations" gorm:"foreignKey:ResourceID;constraint:OnDelete:CASCADE"`
}

// IsExternal checks if the menu item points to an external URL
func (m *MenuItem) IsExternal() bool {
	return m.Target == "_blank" || m.Target == "_parent" || m.Target == "_top"
}

// HasChildren checks if the menu item has child items
func (m *MenuItem) HasChildren() bool {
	return len(m.Children) > 0
}

// GetActiveChildren returns only active child items
func (m *MenuItem) GetActiveChildren() []MenuItem {
	var activeChildren []MenuItem
	for _, child := range m.Children {
		if child.IsActive {
			activeChildren = append(activeChildren, child)
		}
	}
	return activeChildren
}

// TableName returns the table name for Menu model
func (Menu) TableName() string {
	return "menus"
}

// TableName returns the table name for MenuItem model
func (MenuItem) TableName() string {
	return "menu_items"
}
