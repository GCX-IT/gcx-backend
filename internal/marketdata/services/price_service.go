package services

import (
	"gcx-cms/internal/marketdata/models"
	"gcx-cms/internal/shared/config"
	"time"
)

// PriceService handles business logic for market prices
type PriceService struct{}

// NewPriceService creates a new price service instance
func NewPriceService() *PriceService {
	return &PriceService{}
}

// GetCurrentPrices returns current market prices for all commodities
func (ps *PriceService) GetCurrentPrices() ([]models.MarketData, error) {
	var prices []models.MarketData

	// Get latest price for each commodity
	if err := config.DB.Raw(`
		SELECT DISTINCT ON (commodity) * 
		FROM market_data 
		WHERE market_date >= CURRENT_DATE - INTERVAL '1 day'
		ORDER BY commodity, market_date DESC, created_at DESC
	`).Scan(&prices).Error; err != nil {
		return nil, err
	}

	return prices, nil
}

// GetCommodityPrices returns prices for a specific commodity
func (ps *PriceService) GetCommodityPrices(commodity string, days int) ([]models.MarketData, error) {
	var prices []models.MarketData

	if days <= 0 {
		days = 30 // Default to 30 days
	}

	if err := config.DB.Where("commodity = ? AND market_date >= ?",
		commodity, time.Now().AddDate(0, 0, -days)).
		Order("market_date DESC").
		Find(&prices).Error; err != nil {
		return nil, err
	}

	return prices, nil
}

// GetHistoricalPrices returns historical prices with date range
func (ps *PriceService) GetHistoricalPrices(commodity string, startDate, endDate time.Time) ([]models.MarketData, error) {
	var prices []models.MarketData

	query := config.DB.Where("commodity = ? AND market_date BETWEEN ? AND ?",
		commodity, startDate, endDate)

	if err := query.Order("market_date ASC").Find(&prices).Error; err != nil {
		return nil, err
	}

	return prices, nil
}

// GetPriceSummary returns price summary for a commodity
func (ps *PriceService) GetPriceSummary(commodity string) (*models.MarketData, error) {
	var summary models.MarketData

	if err := config.DB.Where("commodity = ? AND market_date = CURRENT_DATE", commodity).
		Order("created_at DESC").
		First(&summary).Error; err != nil {
		return nil, err
	}

	return &summary, nil
}

// UpdatePrice updates or creates a new price record
func (ps *PriceService) UpdatePrice(price *models.MarketData) error {
	if price.ID == 0 {
		// Create new price record
		return config.DB.Create(price).Error
	}

	// Update existing price record
	return config.DB.Save(price).Error
}
