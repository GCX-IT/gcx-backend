package handlers

import (
	"net/http"
	"strconv"

	"gcx-cms/internal/cms/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetBrokers retrieves all brokers with pagination and search
func GetBrokers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	offset := (page - 1) * limit

	var brokers []models.Broker
	var total int64

	query := db.Model(&models.Broker{})

	// Apply search filter
	if search != "" {
		query = query.Where("name LIKE ? OR company LIKE ? OR phone_no LIKE ? OR email LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Get total count
	query.Count(&total)

	// Get paginated results
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&brokers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch brokers",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    brokers,
		"pagination": gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

// GetBroker retrieves a single broker by ID
func GetBroker(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var broker models.Broker
	if err := db.First(&broker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Broker not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    broker,
	})
}

// CreateBroker creates a new broker
func CreateBroker(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var broker models.Broker
	if err := c.ShouldBindJSON(&broker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid broker data",
		})
		return
	}

	if err := db.Create(&broker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create broker",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    broker,
	})
}

// UpdateBroker updates an existing broker
func UpdateBroker(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var broker models.Broker
	if err := db.First(&broker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Broker not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&broker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid broker data",
		})
		return
	}

	if err := db.Save(&broker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update broker",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    broker,
	})
}

// DeleteBroker deletes a broker
func DeleteBroker(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.Broker{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete broker",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Broker deleted successfully",
	})
}
