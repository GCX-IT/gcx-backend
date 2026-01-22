package handlers

import (
	"net/http"
	"time"

	"gcx-cms/internal/marketdata/models"
	"gcx-cms/internal/shared/config"

	"github.com/gin-gonic/gin"
)

// GetRealTimeData returns real-time market data
func GetRealTimeData(c *gin.Context) {
	commodity := c.Query("commodity")

	var data []models.MarketData

	query := config.DB.Where("market_date = CURRENT_DATE")
	if commodity != "" {
		query = query.Where("commodity = ?", commodity)
	}

	if err := query.Order("created_at DESC").Limit(100).Find(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch real-time data",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"commodity": commodity,
		"data":      data,
		"count":     len(data),
		"timestamp": time.Now(),
	})
}

// GetDataStream returns a stream of real-time data (placeholder for WebSocket implementation)
func GetDataStream(c *gin.Context) {
	// This is a placeholder for WebSocket implementation
	// In a real implementation, this would upgrade the connection to WebSocket
	// and stream real-time market data updates

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data stream endpoint - WebSocket implementation required",
		"note":    "This endpoint will be implemented with WebSocket for real-time streaming",
	})
}

// AdminCreatePrice allows admins to create new price records
func AdminCreatePrice(c *gin.Context) {
	var price models.MarketData

	if err := c.ShouldBindJSON(&price); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	// Set market date to today if not provided
	if price.MarketDate.IsZero() {
		price.MarketDate = time.Now()
	}

	if err := config.DB.Create(&price).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create price record",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Price record created successfully",
		"data":    price,
	})
}

// AdminUpdatePrice allows admins to update price records
func AdminUpdatePrice(c *gin.Context) {
	priceID := c.Param("id")
	if priceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Price ID is required",
		})
		return
	}

	var price models.MarketData
	if err := config.DB.First(&price, priceID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Price record not found",
		})
		return
	}

	var req models.MarketData
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	// Update fields
	if err := config.DB.Model(&price).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update price record",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Price record updated successfully",
	})
}

// AdminDeletePrice allows admins to delete price records
func AdminDeletePrice(c *gin.Context) {
	priceID := c.Param("id")
	if priceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Price ID is required",
		})
		return
	}

	if err := config.DB.Delete(&models.MarketData{}, priceID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete price record",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Price record deleted successfully",
	})
}

// AdminCreatePlan allows admins to create subscription plans
func AdminCreatePlan(c *gin.Context) {
	var plan models.SubscriptionPlan

	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	if err := config.DB.Create(&plan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create subscription plan",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Subscription plan created successfully",
		"data":    plan,
	})
}

// AdminUpdatePlan allows admins to update subscription plans
func AdminUpdatePlan(c *gin.Context) {
	planID := c.Param("id")
	if planID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Plan ID is required",
		})
		return
	}

	var plan models.SubscriptionPlan
	if err := config.DB.First(&plan, planID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Subscription plan not found",
		})
		return
	}

	var req models.SubscriptionPlan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	if err := config.DB.Model(&plan).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update subscription plan",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Subscription plan updated successfully",
	})
}

// AdminDeletePlan allows admins to delete subscription plans
func AdminDeletePlan(c *gin.Context) {
	planID := c.Param("id")
	if planID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Plan ID is required",
		})
		return
	}

	if err := config.DB.Delete(&models.SubscriptionPlan{}, planID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete subscription plan",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Subscription plan deleted successfully",
	})
}

// AdminCreateCommodity allows admins to create commodity information
func AdminCreateCommodity(c *gin.Context) {
	var commodity models.CommodityInfo

	if err := c.ShouldBindJSON(&commodity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	if err := config.DB.Create(&commodity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create commodity",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Commodity created successfully",
		"data":    commodity,
	})
}

// AdminUpdateCommodity allows admins to update commodity information
func AdminUpdateCommodity(c *gin.Context) {
	commodityID := c.Param("id")
	if commodityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Commodity ID is required",
		})
		return
	}

	var commodity models.CommodityInfo
	if err := config.DB.First(&commodity, commodityID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Commodity not found",
		})
		return
	}

	var req models.CommodityInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	if err := config.DB.Model(&commodity).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update commodity",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Commodity updated successfully",
	})
}

// AdminDeleteCommodity allows admins to delete commodity information
func AdminDeleteCommodity(c *gin.Context) {
	commodityID := c.Param("id")
	if commodityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Commodity ID is required",
		})
		return
	}

	if err := config.DB.Delete(&models.CommodityInfo{}, commodityID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete commodity",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Commodity deleted successfully",
	})
}
