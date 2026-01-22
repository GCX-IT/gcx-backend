package models

import "time"

// CommodityContractType represents different contract types for a commodity
// Note: Price data (current_price, price_change, trading_volume) comes from Firebase API
// CMS only handles static content (descriptions, images, specifications)
type CommodityContractType struct {
	ID                  uint      `json:"id" gorm:"primaryKey"`
	CommodityID         uint      `json:"commodity_id" gorm:"not null;index"`
	Name                string    `json:"name" gorm:"not null"` // e.g., "Paddy Rice", "Milled Rice", "Parboiled Rice"
	Code                string    `json:"code" gorm:"not null"` // e.g., "PADDY", "MILLED", "PARBOILED"
	Description         string    `json:"description" gorm:"type:text"`
	FullDescription     string    `json:"full_description" gorm:"type:text"`
	Specifications      string    `json:"specifications" gorm:"type:text"`
	TradingHours        string    `json:"trading_hours" gorm:"size:255"`
	ContractSize        string    `json:"contract_size" gorm:"size:100"`
	PriceUnit           string    `json:"price_unit" gorm:"size:50"`
	ImagePath           string    `json:"image_path" gorm:"size:500"`
	ContractFile        string    `json:"contract_file" gorm:"size:500"`
	DeliveryMonths      string    `json:"delivery_months" gorm:"size:255"`
	StorageRequirements string    `json:"storage_requirements" gorm:"type:text"`
	QualityStandards    string    `json:"quality_standards" gorm:"type:text"`
	IsActive            bool      `json:"is_active" gorm:"default:true"`
	SortOrder           int       `json:"sort_order" gorm:"default:0"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`

	// Relationship
	Commodity Commodity `json:"commodity" gorm:"foreignKey:CommodityID"`
}

// TableName specifies the table name for CommodityContractType
func (CommodityContractType) TableName() string {
	return "commodity_contract_types"
}
