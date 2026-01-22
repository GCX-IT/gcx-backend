package repository

import (
	"gcx-cms/internal/marketdata/models"
	"gcx-cms/internal/shared/config"
	"time"
)

// PriceRepository handles data access for market prices
type PriceRepository struct{}

// NewPriceRepository creates a new price repository instance
func NewPriceRepository() *PriceRepository {
	return &PriceRepository{}
}

// Create creates a new price record
func (pr *PriceRepository) Create(price *models.MarketData) error {
	return config.DB.Create(price).Error
}

// Update updates an existing price record
func (pr *PriceRepository) Update(price *models.MarketData) error {
	return config.DB.Save(price).Error
}

// Delete deletes a price record
func (pr *PriceRepository) Delete(id uint) error {
	return config.DB.Delete(&models.MarketData{}, id).Error
}

// FindByID finds a price record by ID
func (pr *PriceRepository) FindByID(id uint) (*models.MarketData, error) {
	var price models.MarketData
	err := config.DB.First(&price, id).Error
	if err != nil {
		return nil, err
	}
	return &price, nil
}

// FindByCommodity finds price records by commodity
func (pr *PriceRepository) FindByCommodity(commodity string) ([]models.MarketData, error) {
	var prices []models.MarketData
	err := config.DB.Where("commodity = ?", commodity).
		Order("market_date DESC").
		Find(&prices).Error
	return prices, err
}

// FindLatestByCommodity finds the latest price for each commodity
func (pr *PriceRepository) FindLatestByCommodity() ([]models.MarketData, error) {
	var prices []models.MarketData

	err := config.DB.Raw(`
		SELECT DISTINCT ON (commodity) * 
		FROM market_data 
		WHERE market_date >= CURRENT_DATE - INTERVAL '1 day'
		ORDER BY commodity, market_date DESC, created_at DESC
	`).Scan(&prices).Error

	return prices, err
}

// FindByDateRange finds price records within a date range
func (pr *PriceRepository) FindByDateRange(commodity string, startDate, endDate time.Time) ([]models.MarketData, error) {
	var prices []models.MarketData

	query := config.DB.Where("commodity = ? AND market_date BETWEEN ? AND ?",
		commodity, startDate, endDate)

	err := query.Order("market_date ASC").Find(&prices).Error
	return prices, err
}

// FindTodaySummary finds today's trading summary for a commodity
func (pr *PriceRepository) FindTodaySummary(commodity string) (*models.MarketData, error) {
	var summary models.MarketData

	err := config.DB.Where("commodity = ? AND market_date = CURRENT_DATE", commodity).
		Order("created_at DESC").
		First(&summary).Error

	if err != nil {
		return nil, err
	}

	return &summary, nil
}

// BulkCreate creates multiple price records
func (pr *PriceRepository) BulkCreate(prices []models.MarketData) error {
	return config.DB.CreateInBatches(prices, 100).Error
}

// GetCommodityList returns list of all available commodities
func (pr *PriceRepository) GetCommodityList() ([]string, error) {
	var commodities []string

	err := config.DB.Model(&models.MarketData{}).
		Distinct("commodity").
		Pluck("commodity", &commodities).Error

	return commodities, err
}
