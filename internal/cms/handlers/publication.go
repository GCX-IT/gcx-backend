package handlers

import (
	"net/http"
	"strconv"

	"gcx-cms/internal/cms/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetPublications retrieves all publications with pagination and search
func GetPublications(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")
	category := c.Query("category")

	offset := (page - 1) * limit

	var publications []models.Publication
	var total int64

	query := db.Model(&models.Publication{})

	// Apply search filter
	if search != "" {
		query = query.Where("title LIKE ? OR description LIKE ? OR author LIKE ? OR tags LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Apply category filter
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Get total count
	query.Count(&total)

	// Get paginated results
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&publications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch publications",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    publications,
		"pagination": gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

// GetPublication retrieves a single publication by ID
func GetPublication(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var publication models.Publication
	if err := db.First(&publication, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Publication not found",
		})
		return
	}

	// Increment download count
	db.Model(&publication).Update("download_count", publication.DownloadCount+1)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    publication,
	})
}

// CreatePublication creates a new publication
func CreatePublication(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var publication models.Publication
	if err := c.ShouldBindJSON(&publication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid publication data",
		})
		return
	}

	if err := db.Create(&publication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create publication",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    publication,
	})
}

// UpdatePublication updates an existing publication
func UpdatePublication(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var publication models.Publication
	if err := db.First(&publication, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Publication not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&publication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid publication data",
		})
		return
	}

	if err := db.Save(&publication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update publication",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    publication,
	})
}

// DeletePublication deletes a publication
func DeletePublication(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.Publication{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete publication",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Publication deleted successfully",
	})
}
