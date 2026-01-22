package handlers

import (
	"net/http"
	"strconv"

	"gcx-cms/internal/cms/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetTraders retrieves all traders with pagination and search
func GetTraders(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	offset := (page - 1) * limit

	var traders []models.Trader
	var total int64

	query := db.Model(&models.Trader{})

	// Apply search filter
	if search != "" {
		query = query.Where("name LIKE ? OR phone_no LIKE ? OR email LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Get total count
	query.Count(&total)

	// Get paginated results
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&traders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch traders",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    traders,
		"pagination": gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

// GetTrader retrieves a single trader by ID
func GetTrader(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var trader models.Trader
	if err := db.First(&trader, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Trader not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    trader,
	})
}

// CreateTrader creates a new trader
func CreateTrader(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var trader models.Trader
	if err := c.ShouldBindJSON(&trader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid trader data",
		})
		return
	}

	if err := db.Create(&trader).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create trader",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    trader,
	})
}

// UpdateTrader updates an existing trader
func UpdateTrader(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var trader models.Trader
	if err := db.First(&trader, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Trader not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&trader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid trader data",
		})
		return
	}

	if err := db.Save(&trader).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update trader",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    trader,
	})
}

// DeleteTrader deletes a trader
func DeleteTrader(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.Trader{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete trader",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Trader deleted successfully",
	})
}
