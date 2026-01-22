package handlers

import (
	"net/http"
	"time"

	"gcx-cms/internal/marketdata/models"
	"gcx-cms/internal/shared/config"

	"github.com/gin-gonic/gin"
)

// GetCurrentPrices returns current market prices for all commodities
func GetCurrentPrices(c *gin.Context) {
	var prices []models.MarketData

	// Get latest price for each commodity
	if err := config.DB.Raw(`
		SELECT DISTINCT ON (commodity) * 
		FROM market_data 
		WHERE market_date >= CURRENT_DATE - INTERVAL '1 day'
		ORDER BY commodity, market_date DESC, created_at DESC
	`).Scan(&prices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch current prices",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"data":      prices,
		"timestamp": time.Now(),
	})
}

// GetCommodityPrices returns prices for a specific commodity
func GetCommodityPrices(c *gin.Context) {
	commodity := c.Param("commodity")
	if commodity == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Commodity parameter is required",
		})
		return
	}

	var prices []models.MarketData

	// Get prices for the last 30 days
	if err := config.DB.Where("commodity = ? AND market_date >= ?",
		commodity, time.Now().AddDate(0, 0, -30)).
		Order("market_date DESC").
		Find(&prices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch commodity prices",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"commodity": commodity,
		"data":      prices,
		"count":     len(prices),
	})
}

// GetHistoricalPrices returns historical prices with date range
func GetHistoricalPrices(c *gin.Context) {
	commodity := c.Query("commodity")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if commodity == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Commodity parameter is required",
		})
		return
	}

	// Parse date parameters
	var start, end time.Time
	var err error

	if startDate != "" {
		start, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid start_date format. Use YYYY-MM-DD",
			})
			return
		}
	} else {
		start = time.Now().AddDate(0, 0, -30) // Default to 30 days ago
	}

	if endDate != "" {
		end, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid end_date format. Use YYYY-MM-DD",
			})
			return
		}
	} else {
		end = time.Now() // Default to today
	}

	var prices []models.MarketData

	query := config.DB.Where("commodity = ? AND market_date BETWEEN ? AND ?",
		commodity, start, end)

	if err := query.Order("market_date ASC").Find(&prices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch historical prices",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"commodity":  commodity,
		"start_date": start,
		"end_date":   end,
		"data":       prices,
		"count":      len(prices),
	})
}

// GetPriceSummary returns price summary for a commodity
func GetPriceSummary(c *gin.Context) {
	commodity := c.Param("commodity")
	if commodity == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Commodity parameter is required",
		})
		return
	}

	// Get today's trading summary
	var summary struct {
		Commodity     string    `json:"commodity"`
		CurrentPrice  float64   `json:"current_price"`
		Change        float64   `json:"change"`
		ChangePercent float64   `json:"change_percent"`
		High          float64   `json:"high"`
		Low           float64   `json:"low"`
		Open          float64   `json:"open"`
		Volume        float64   `json:"volume"`
		LastUpdated   time.Time `json:"last_updated"`
	}

	if err := config.DB.Raw(`
		SELECT 
			commodity,
			price as current_price,
			change,
			change_percent,
			COALESCE(high, 0) as high,
			COALESCE(low, 0) as low,
			COALESCE(open, 0) as open,
			COALESCE(volume, 0) as volume,
			updated_at as last_updated
		FROM market_data 
		WHERE commodity = ? AND market_date = CURRENT_DATE
		ORDER BY created_at DESC 
		LIMIT 1
	`, commodity).Scan(&summary).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch price summary",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    summary,
	})
}
