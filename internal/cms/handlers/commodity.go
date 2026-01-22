package handlers

import (
	"net/http"
	"strconv"

	"gcx-cms/internal/cms/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetCommodities retrieves all commodities with pagination and search
func GetCommodities(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")
	category := c.Query("category")
	status := c.Query("status")

	offset := (page - 1) * limit

	var commodities []models.Commodity
	var total int64

	query := db.Model(&models.Commodity{})

	// Apply search filter
	if search != "" {
		query = query.Where("name LIKE ? OR code LIKE ? OR description LIKE ? OR category LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Apply category filter
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Apply status filter
	if status != "" {
		query = query.Where("market_status = ?", status)
	}

	// Get total count
	query.Count(&total)

	// Get paginated results
	if err := query.Offset(offset).Limit(limit).Order("name ASC").Find(&commodities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch commodities",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    commodities,
		"pagination": gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

// GetCommodity retrieves a single commodity by ID
func GetCommodity(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var commodity models.Commodity
	if err := db.First(&commodity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Commodity not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    commodity,
	})
}

// CreateCommodity creates a new commodity
func CreateCommodity(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var commodity models.Commodity
	if err := c.ShouldBindJSON(&commodity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid commodity data",
		})
		return
	}

	if err := db.Create(&commodity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create commodity",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    commodity,
	})
}

// UpdateCommodity updates an existing commodity
func UpdateCommodity(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var commodity models.Commodity
	if err := db.First(&commodity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Commodity not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&commodity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid commodity data",
		})
		return
	}

	if err := db.Save(&commodity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update commodity",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    commodity,
	})
}

// DeleteCommodity deletes a commodity
func DeleteCommodity(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.Commodity{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete commodity",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Commodity deleted successfully",
	})
}
