package handlers

import (
	"log"
	"net/http"
	"strings"

	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetSettings returns all settings with optional filtering
func GetSettings(c *gin.Context) {
	var settings []models.Setting
	query := config.DB

	// Filter by group
	if group := c.Query("group"); group != "" {
		query = query.Where("`group` = ?", group)
	}

	// Filter by public/private
	if isPublic := c.Query("is_public"); isPublic != "" {
		if isPublic == "true" {
			query = query.Where("is_public = ?", true)
		} else if isPublic == "false" {
			query = query.Where("is_public = ?", false)
		}
	}

	// Order by group and sort_order
	query = query.Order("`group` ASC, sort_order ASC, `key` ASC")

	if err := query.Find(&settings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch settings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"settings": settings})
}

// GetSetting returns a single setting by key
func GetSetting(c *gin.Context) {
	key := c.Param("key")

	var setting models.Setting
	if err := config.DB.Where("`key` = ?", key).First(&setting).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Setting not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch setting"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"setting": setting})
}

// GetSettingsByGroup returns all settings in a specific group
func GetSettingsByGroup(c *gin.Context) {
	group := c.Param("group")

	var settings []models.Setting
	if err := config.DB.Where("`group` = ?", group).
		Order("sort_order ASC, `key` ASC").
		Find(&settings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch settings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"settings": settings})
}

// CreateSetting creates a new setting
func CreateSetting(c *gin.Context) {
	var setting models.Setting
	if err := c.ShouldBindJSON(&setting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if setting with this key already exists
	var existingSetting models.Setting
	if err := config.DB.Where("`key` = ?", setting.Key).First(&existingSetting).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Setting with this key already exists"})
		return
	}

	if err := config.DB.Create(&setting).Error; err != nil {
		log.Printf("Failed to create setting: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create setting",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"setting": setting})
}

// UpdateSetting updates an existing setting
func UpdateSetting(c *gin.Context) {
	key := c.Param("key")

	var setting models.Setting
	if err := config.DB.Where("`key` = ?", key).First(&setting).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Setting not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch setting"})
		}
		return
	}

	var updateData models.Setting
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields (key cannot be changed)
	setting.Value = updateData.Value
	setting.Type = updateData.Type
	setting.Group = updateData.Group
	setting.Label = updateData.Label
	setting.Description = updateData.Description
	setting.IsPublic = updateData.IsPublic
	setting.SortOrder = updateData.SortOrder

	if err := config.DB.Save(&setting).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update setting"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"setting": setting})
}

// UpdateSettingsBatch updates multiple settings at once
func UpdateSettingsBatch(c *gin.Context) {
	var settings []models.Setting
	if err := c.ShouldBindJSON(&settings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Start a transaction
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, setting := range settings {
		if err := tx.Model(&models.Setting{}).
			Where("`key` = ?", setting.Key).
			Updates(map[string]interface{}{
				"value":       setting.Value,
				"type":        setting.Type,
				"group":       setting.Group,
				"label":       setting.Label,
				"description": setting.Description,
				"is_public":   setting.IsPublic,
				"sort_order":  setting.SortOrder,
			}).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update settings"})
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit settings updates"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Settings updated successfully"})
}

// DeleteSetting deletes a setting
func DeleteSetting(c *gin.Context) {
	key := c.Param("key")

	var setting models.Setting
	if err := config.DB.Where("`key` = ?", key).First(&setting).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Setting not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch setting"})
		}
		return
	}

	if err := config.DB.Delete(&setting).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete setting"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Setting deleted successfully"})
}

// GetPublicSettings returns only public settings (for frontend use)
func GetPublicSettings(c *gin.Context) {
	var settings []models.Setting

	// First, check if the table exists by trying to count records
	var count int64
	if err := config.DB.Model(&models.Setting{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Database table error",
			"details": err.Error(),
			"table":   "settings",
		})
		return
	}

	// If no settings exist, return empty map
	if count == 0 {
		c.JSON(http.StatusOK, gin.H{"settings": make(map[string]interface{})})
		return
	}

	// Fetch public settings - try a simpler query first
	if err := config.DB.Where("is_public = ?", true).Find(&settings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch public settings",
			"details": err.Error(),
		})
		return
	}

	// Convert to key-value map for easier frontend consumption
	settingsMap := make(map[string]interface{})
	for _, setting := range settings {
		settingsMap[setting.Key] = setting.Value
	}

	c.JSON(http.StatusOK, gin.H{"settings": settingsMap})
}

// GetSettingsByGroupPublic returns public settings for a specific group
func GetSettingsByGroupPublic(c *gin.Context) {
	// Extract group from URL path
	group := c.Param("group")
	if group == "" {
		// Try to extract from the last part of the path
		path := c.Request.URL.Path
		if path != "" {
			parts := strings.Split(path, "/")
			if len(parts) > 0 {
				group = parts[len(parts)-1]
			}
		}
	}

	if group == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group parameter is required"})
		return
	}

	var settings []models.Setting
	if err := config.DB.Where("`group` = ? AND is_public = ?", group, true).Find(&settings).Error; err != nil {
		log.Printf("Error fetching public settings for group %s: %v", group, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch settings"})
		return
	}

	// Convert to map for easier frontend consumption
	settingsMap := make(map[string]interface{})
	for _, setting := range settings {
		settingsMap[setting.Key] = setting.Value
	}

	c.JSON(http.StatusOK, gin.H{"settings": settingsMap})
}
