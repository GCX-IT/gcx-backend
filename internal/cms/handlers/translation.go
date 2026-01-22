package handlers

import (
	"net/http"

	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetTranslations returns all translations
func GetTranslations(c *gin.Context) {
	var translations []models.Translation
	if err := config.DB.Find(&translations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch translations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"translations": translations})
}

// GetTranslation returns a single translation by ID
func GetTranslation(c *gin.Context) {
	id := c.Param("id")
	var translation models.Translation
	if err := config.DB.First(&translation, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Translation not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch translation"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"translation": translation})
}

// CreateTranslation creates a new translation
func CreateTranslation(c *gin.Context) {
	var translation models.Translation
	if err := c.ShouldBindJSON(&translation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&translation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create translation"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"translation": translation})
}

// UpdateTranslation updates an existing translation
func UpdateTranslation(c *gin.Context) {
	id := c.Param("id")
	var translation models.Translation
	if err := config.DB.First(&translation, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Translation not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch translation"})
		}
		return
	}

	var updateData models.Translation
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	translation.ResourceType = updateData.ResourceType
	translation.ResourceID = updateData.ResourceID
	translation.LanguageCode = updateData.LanguageCode
	translation.FieldName = updateData.FieldName
	translation.Content = updateData.Content

	if err := config.DB.Save(&translation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update translation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"translation": translation})
}

// DeleteTranslation deletes a translation
func DeleteTranslation(c *gin.Context) {
	id := c.Param("id")
	var translation models.Translation
	if err := config.DB.First(&translation, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Translation not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch translation"})
		}
		return
	}

	if err := config.DB.Delete(&translation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete translation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Translation deleted successfully"})
}
