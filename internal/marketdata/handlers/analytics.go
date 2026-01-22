package handlers

import (
	"net/http"
	"time"

	"gcx-cms/internal/marketdata/models"
	"gcx-cms/internal/shared/config"

	"github.com/gin-gonic/gin"
)

// GetMarketAnalytics returns market analytics data
func GetMarketAnalytics(c *gin.Context) {
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

	var analytics []models.MarketAnalytics

	query := config.DB.Where("commodity = ? AND date BETWEEN ? AND ?",
		commodity, start, end)

	if err := query.Order("date ASC").Find(&analytics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch market analytics",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"commodity":  commodity,
		"start_date": start,
		"end_date":   end,
		"data":       analytics,
		"count":      len(analytics),
	})
}

// GetPriceAlerts returns user's price alerts
func GetPriceAlerts(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	var alerts []models.PriceAlert

	if err := config.DB.Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&alerts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch price alerts",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    alerts,
		"count":   len(alerts),
	})
}

// CreatePriceAlert creates a new price alert
func CreatePriceAlert(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	var req struct {
		Commodity   string  `json:"commodity" binding:"required"`
		TargetPrice float64 `json:"target_price" binding:"required"`
		Condition   string  `json:"condition" binding:"required"` // above, below
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	// Validate condition
	if req.Condition != "above" && req.Condition != "below" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Condition must be 'above' or 'below'",
		})
		return
	}

	alert := models.PriceAlert{
		UserID:      userID.(uint),
		Commodity:   req.Commodity,
		TargetPrice: req.TargetPrice,
		Condition:   req.Condition,
		IsActive:    true,
	}

	if err := config.DB.Create(&alert).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create price alert",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Price alert created successfully",
		"data":    alert,
	})
}

// UpdatePriceAlert updates a price alert
func UpdatePriceAlert(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	alertID := c.Param("id")
	if alertID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Alert ID is required",
		})
		return
	}

	var req struct {
		TargetPrice *float64 `json:"target_price"`
		Condition   *string  `json:"condition"`
		IsActive    *bool    `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	// Get the alert
	var alert models.PriceAlert
	if err := config.DB.Where("id = ? AND user_id = ?", alertID, userID).
		First(&alert).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Price alert not found",
		})
		return
	}

	// Update fields if provided
	updates := make(map[string]interface{})
	if req.TargetPrice != nil {
		updates["target_price"] = *req.TargetPrice
	}
	if req.Condition != nil {
		if *req.Condition != "above" && *req.Condition != "below" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Condition must be 'above' or 'below'",
			})
			return
		}
		updates["condition"] = *req.Condition
	}
	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}

	if len(updates) > 0 {
		if err := config.DB.Model(&alert).Updates(updates).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to update price alert",
				"details": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Price alert updated successfully",
	})
}

// DeletePriceAlert deletes a price alert
func DeletePriceAlert(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	alertID := c.Param("id")
	if alertID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Alert ID is required",
		})
		return
	}

	// Delete the alert
	if err := config.DB.Where("id = ? AND user_id = ?", alertID, userID).
		Delete(&models.PriceAlert{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete price alert",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Price alert deleted successfully",
	})
}
