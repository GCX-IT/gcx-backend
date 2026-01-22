package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetNewsItems returns all active news items for the ticker
func GetNewsItems(c *gin.Context) {
	db := database.GetDB()

	var newsItems []models.NewsItem

	// Get query parameters
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit > 100 {
		limit = 20
	}

	// Build query for published news - match blog posts filtering style
	// Simple filtering: just check status and published_at (like blog posts)
	query := db.Where("status = ?", models.NewsStatusPublished).
		Where("published_at IS NOT NULL")
	
	// Optional: Only filter by expires_at if it's set and in the past
	// This allows items without expiration to show indefinitely
	query = query.Where("(expires_at IS NULL OR expires_at > ?)", time.Now())

	// Filter by source if provided
	if source := c.Query("source"); source != "" {
		query = query.Where("source = ?", source)
	}

	// Filter by category if provided
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}

	// Filter breaking news only
	if c.Query("breaking") == "true" {
		query = query.Where("is_breaking = ?", true)
	}

	// Order by newest first (published_at DESC, or created_at if published_at is NULL)
	// Breaking news and priority are still considered but newest items come first
	err = query.Order("COALESCE(published_at, created_at) DESC").
		Limit(limit).
		Find(&newsItems).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news items"})
		return
	}

	// Log for debugging
	log.Printf("📰 GetNewsItems: Found %d published news items (limit: %d)", len(newsItems), limit)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    newsItems,
		"count":   len(newsItems),
	})
}

// GetBreakingNews returns only breaking news items
func GetBreakingNews(c *gin.Context) {
	db := database.GetDB()

	var newsItems []models.NewsItem

	err := db.Where("status = ? AND is_active = ? AND is_breaking = ?",
		models.NewsStatusPublished, true, true).
		Where("(published_at IS NULL OR published_at <= ?)", time.Now()).
		Where("(expires_at IS NULL OR expires_at > ?)", time.Now()).
		Order("priority DESC, published_at DESC").
		Find(&newsItems).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch breaking news"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    newsItems,
		"count":   len(newsItems),
	})
}

// GetNewsItem returns a single news item by ID
func GetNewsItem(c *gin.Context) {
	db := database.GetDB()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news item ID"})
		return
	}

	var newsItem models.NewsItem
	err = db.First(&newsItem, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "News item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    newsItem,
	})
}

// CreateNewsItem creates a new news item (CMS only)
func CreateNewsItem(c *gin.Context) {
	db := database.GetDB()

	var newsItem models.NewsItem
	if err := c.ShouldBindJSON(&newsItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set defaults
	if newsItem.Source == "" {
		newsItem.Source = models.NewsSourceGCX
	}
	if newsItem.Status == "" {
		newsItem.Status = models.NewsStatusDraft
	}

	err := db.Create(&newsItem).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create news item"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    newsItem,
	})
}

// UpdateNewsItem updates an existing news item (CMS only)
func UpdateNewsItem(c *gin.Context) {
	db := database.GetDB()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news item ID"})
		return
	}

	var newsItem models.NewsItem
	err = db.First(&newsItem, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "News item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news item"})
		return
	}

	if err := c.ShouldBindJSON(&newsItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.Save(&newsItem).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    newsItem,
	})
}

// DeleteNewsItem deletes a news item (CMS only)
func DeleteNewsItem(c *gin.Context) {
	db := database.GetDB()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news item ID"})
		return
	}

	var newsItem models.NewsItem
	err = db.First(&newsItem, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "News item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news item"})
		return
	}

	err = db.Delete(&newsItem).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete news item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "News item deleted successfully",
	})
}

// GetAllNewsItems returns all news items for CMS management
func GetAllNewsItems(c *gin.Context) {
	db := database.GetDB()

	var newsItems []models.NewsItem

	// Get query parameters for filtering
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "20")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 20
	}

	offset := (page - 1) * limit

	// Build query
	query := db.Model(&models.NewsItem{})

	// Filter by status
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// Filter by source
	if source := c.Query("source"); source != "" {
		query = query.Where("source = ?", source)
	}

	// Filter by category
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}

	// Filter by breaking news
	if breaking := c.Query("breaking"); breaking != "" {
		query = query.Where("is_breaking = ?", breaking == "true")
	}

	// Get total count
	var total int64
	query.Count(&total)

	// Get items
	err = query.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&newsItems).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    newsItems,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

// GetNewsCategories returns all news categories
func GetNewsCategories(c *gin.Context) {
	db := database.GetDB()

	var categories []models.NewsCategory

	err := db.Where("is_active = ?", true).
		Order("name ASC").
		Find(&categories).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news categories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    categories,
	})
}

// CreateNewsCategory creates a new news category (CMS only)
func CreateNewsCategory(c *gin.Context) {
	db := database.GetDB()

	var category models.NewsCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Create(&category).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create news category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    category,
	})
}

// UpdateNewsCategory updates an existing news category (CMS only)
func UpdateNewsCategory(c *gin.Context) {
	db := database.GetDB()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var category models.NewsCategory
	err = db.First(&category, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "News category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news category"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.Save(&category).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    category,
	})
}

// DeleteNewsCategory deletes a news category (CMS only)
func DeleteNewsCategory(c *gin.Context) {
	db := database.GetDB()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var category models.NewsCategory
	err = db.First(&category, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "News category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news category"})
		return
	}

	err = db.Delete(&category).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete news category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "News category deleted successfully",
	})
}

// PublishNewsItem publishes a news item (CMS only)
func PublishNewsItem(c *gin.Context) {
	db := database.GetDB()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news item ID"})
		return
	}

	var newsItem models.NewsItem
	err = db.First(&newsItem, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "News item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news item"})
		return
	}

	newsItem.Publish()

	err = db.Save(&newsItem).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish news item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    newsItem,
		"message": "News item published successfully",
	})
}

// ArchiveNewsItem archives a news item (CMS only)
func ArchiveNewsItem(c *gin.Context) {
	db := database.GetDB()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news item ID"})
		return
	}

	var newsItem models.NewsItem
	err = db.First(&newsItem, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "News item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news item"})
		return
	}

	newsItem.Archive()

	err = db.Save(&newsItem).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to archive news item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    newsItem,
		"message": "News item archived successfully",
	})
}

// SetBreakingNews sets a news item as breaking news (CMS only)
func SetBreakingNews(c *gin.Context) {
	db := database.GetDB()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news item ID"})
		return
	}

	var newsItem models.NewsItem
	err = db.First(&newsItem, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "News item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news item"})
		return
	}

	var request struct {
		IsBreaking bool `json:"is_breaking"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newsItem.SetBreaking(request.IsBreaking)

	err = db.Save(&newsItem).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update breaking news status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    newsItem,
		"message": "Breaking news status updated successfully",
	})
}
