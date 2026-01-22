package handlers

import (
	"net/http"
	"strconv"

	"gcx-cms/internal/cms/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetCareers retrieves all careers with pagination and search
func GetCareers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")
	category := c.Query("category")
	status := c.Query("status")

	offset := (page - 1) * limit

	var careers []models.Career
	var total int64

	query := db.Model(&models.Career{})

	// Apply search filter
	if search != "" {
		query = query.Where("title LIKE ? OR description LIKE ? OR department LIKE ? OR location LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Apply category filter
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Apply status filter
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Get total count
	query.Count(&total)

	// Get paginated results
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&careers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch careers",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    careers,
		"pagination": gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

// GetCareer retrieves a single career by ID
func GetCareer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var career models.Career
	if err := db.First(&career, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Career not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    career,
	})
}

// CreateCareer creates a new career
func CreateCareer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var career models.Career
	if err := c.ShouldBindJSON(&career); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid career data",
		})
		return
	}

	if err := db.Create(&career).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create career",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    career,
	})
}

// UpdateCareer updates an existing career
func UpdateCareer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var career models.Career
	if err := db.First(&career, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Career not found",
		})
		return
	}

	// Create a map to bind the JSON data
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid career data",
		})
		return
	}

	// Update the career with the new data
	if err := db.Model(&career).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update career",
		})
		return
	}

	// Reload the career to get the updated data
	if err := db.First(&career, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to reload career",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    career,
	})
}

// DeleteCareer deletes a career
func DeleteCareer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.Career{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete career",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Career deleted successfully",
	})
}
