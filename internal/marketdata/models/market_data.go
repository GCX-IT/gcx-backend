package models

import (
	"time"
	"gorm.io/datatypes"
)

// MarketData represents commodity market data
type MarketData struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Commodity   string         `json:"commodity" gorm:"not null"` // maize, soybean, etc.
	Price       float64        `json:"price" gorm:"not null"`
	Currency    string         `json:"currency" gorm:"default:GHS"`
	Unit        string         `json:"unit" gorm:"default:metric_ton"`
	Change      float64        `json:"change"` // Price change from previous
	ChangePercent float64      `json:"change_percent"`
	Volume      *float64       `json:"volume"`
	High        *float64       `json:"high"` // Day's high
	Low         *float64       `json:"low"`  // Day's low
	Open        *float64       `json:"open"` // Day's opening price
	Close       *float64       `json:"close"` // Day's closing price
	MarketDate  time.Time      `json:"market_date"`
	Source      string         `json:"source"` // GCX, external API, etc.
	Metadata    datatypes.JSON `json:"metadata" gorm:"type:json"` // Additional data
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// CommodityInfo represents static commodity information
type CommodityInfo struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Name          string         `json:"name" gorm:"not null"`
	Code          string         `json:"code" gorm:"uniqueIndex;not null"` // MAIZE, SOYBEAN
	Description   string         `json:"description" gorm:"type:text"`
	Specifications datatypes.JSON `json:"specifications" gorm:"type:json"`
	TradingInfo   datatypes.JSON `json:"trading_info" gorm:"type:json"`
	IsActive      bool           `json:"is_active" gorm:"default:true"`
	SortOrder     int            `json:"sort_order" gorm:"default:0"`
	Image         *string        `json:"image"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

// TradingSession represents trading session information
type TradingSession struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Date        time.Time `json:"date" gorm:"uniqueIndex"`
	OpenTime    time.Time `json:"open_time"`
	CloseTime   time.Time `json:"close_time"`
	IsOpen      bool      `json:"is_open" gorm:"default:false"`
	Status      string    `json:"status"` // open, closed, pre_market, post_market
	Volume      float64   `json:"volume"`
	Transactions int      `json:"transactions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PriceAlert represents user price alerts
type PriceAlert struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id"`
	Commodity  string    `json:"commodity"`
	TargetPrice float64  `json:"target_price"`
	Condition  string    `json:"condition"` // above, below
	IsActive   bool      `json:"is_active" gorm:"default:true"`
	TriggeredAt *time.Time `json:"triggered_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// MarketAnalytics represents aggregated market analytics
type MarketAnalytics struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Commodity       string    `json:"commodity" gorm:"not null"`
	Date            time.Time `json:"date" gorm:"not null"`
	TotalVolume     float64   `json:"total_volume"`
	AveragePrice    float64   `json:"average_price"`
	PriceChange     float64   `json:"price_change"`
	PriceChangePercent float64 `json:"price_change_percent"`
	HighPrice       float64   `json:"high_price"`
	LowPrice        float64   `json:"low_price"`
	OpenPrice       float64   `json:"open_price"`
	ClosePrice      float64   `json:"close_price"`
	TransactionCount int      `json:"transaction_count"`
	MarketSentiment string    `json:"market_sentiment"` // bullish, bearish, neutral
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// TableName returns the table name for MarketData model
func (MarketData) TableName() string {
	return "market_data"
}

// TableName returns the table name for CommodityInfo model
func (CommodityInfo) TableName() string {
	return "commodity_info"
}

// TableName returns the table name for TradingSession model
func (TradingSession) TableName() string {
	return "trading_sessions"
}

// TableName returns the table name for PriceAlert model
func (PriceAlert) TableName() string {
	return "price_alerts"
}

// TableName returns the table name for MarketAnalytics model
func (MarketAnalytics) TableName() string {
	return "market_analytics"
}
