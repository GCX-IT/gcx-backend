package models

import "time"

// Commodity represents a commodity in the system
type Commodity struct {
	ID                  uint      `json:"id" gorm:"primaryKey"`
	Name                string    `json:"name" gorm:"not null"`
	Code                string    `json:"code" gorm:"uniqueIndex;not null"`
	Description         string    `json:"description" gorm:"type:text"`
	FullDescription     string    `json:"full_description" gorm:"type:text"`
	Specifications      string    `json:"specifications" gorm:"type:text"`
	TradingHours        string    `json:"trading_hours" gorm:"size:255"`
	ContractSize        string    `json:"contract_size" gorm:"size:100"`
	PriceUnit           string    `json:"price_unit" gorm:"size:50"`
	MinimumPrice        float64   `json:"minimum_price" gorm:"type:decimal(10,2)"`
	MaximumPrice        float64   `json:"maximum_price" gorm:"type:decimal(10,2)"`
	CurrentPrice        float64   `json:"current_price" gorm:"type:decimal(10,2)"`
	PriceChange         float64   `json:"price_change" gorm:"type:decimal(10,2)"`
	PriceChangePercent  float64   `json:"price_change_percent" gorm:"type:decimal(5,2)"`
	TradingVolume       int64     `json:"trading_volume" gorm:"default:0"`
	MarketStatus        string    `json:"market_status" gorm:"type:enum('Open','Closed','Suspended');default:'Open'"`
	ImagePath           string    `json:"image_path" gorm:"size:500"`
	ContractFile        string    `json:"contract_file" gorm:"size:500"`
	Category            string    `json:"category" gorm:"size:100"`
	OriginCountry       string    `json:"origin_country" gorm:"size:100"`
	HarvestSeason       string    `json:"harvest_season" gorm:"size:100"`
	DeliveryMonths      string    `json:"delivery_months" gorm:"size:255"`
	StorageRequirements string    `json:"storage_requirements" gorm:"type:text"`
	QualityStandards    string    `json:"quality_standards" gorm:"type:text"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`

	// Relationships
	ContractTypes []CommodityContractType `json:"contract_types" gorm:"foreignKey:CommodityID"`
}
