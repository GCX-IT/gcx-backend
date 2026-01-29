package handlers

import (
	"net/http"

	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetRTIDocuments retrieves all active RTI documents (public endpoint)
func GetRTIDocuments(c *gin.Context) {
	db := database.GetDB()
	category := c.Query("category")

	var documents []models.RTIDocument
	query := db.Where("is_active = ?", true)

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if err := query.Order("sort_order ASC, created_at DESC").Find(&documents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve RTI documents"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": documents})
}

// GetRTIDocument retrieves a single RTI document (public endpoint)
func GetRTIDocument(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var document models.RTIDocument
	if err := db.Where("id = ? AND is_active = ?", id, true).First(&document).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "RTI document not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve RTI document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": document})
}

// DownloadRTIDocument handles document download and increments counter (public endpoint)
func DownloadRTIDocument(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var document models.RTIDocument
	if err := db.Where("id = ? AND is_active = ?", id, true).First(&document).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "RTI document not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve RTI document"})
		return
	}

	// Increment download count
	document.IncrementDownloadCount(db)

	c.JSON(http.StatusOK, gin.H{
		"data":         document,
		"download_url": document.FilePath,
	})
}

// GetAllRTIDocuments retrieves all RTI documents for CMS (protected endpoint)
func GetAllRTIDocuments(c *gin.Context) {
	db := database.GetDB()

	var documents []models.RTIDocument
	if err := db.Order("sort_order ASC, created_at DESC").Find(&documents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve RTI documents"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": documents})
}

// CreateRTIDocument creates a new RTI document (protected endpoint)
func CreateRTIDocument(c *gin.Context) {
	db := database.GetDB()

	var document models.RTIDocument
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&document).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create RTI document"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "RTI document created successfully",
		"data":    document,
	})
}

// UpdateRTIDocument updates an existing RTI document (protected endpoint)
func UpdateRTIDocument(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var document models.RTIDocument
	if err := db.First(&document, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "RTI document not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve RTI document"})
		return
	}

	var updates models.RTIDocument
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&document).Updates(&updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update RTI document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "RTI document updated successfully",
		"data":    document,
	})
}

// DeleteRTIDocument deletes an RTI document (soft delete) (protected endpoint)
func DeleteRTIDocument(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var document models.RTIDocument
	if err := db.First(&document, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "RTI document not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve RTI document"})
		return
	}

	if err := db.Delete(&document).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete RTI document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "RTI document deleted successfully"})
}
