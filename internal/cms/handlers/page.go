package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/config"
	shared_models "gcx-cms/internal/shared/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetPages returns all pages with optional filtering
func GetPages(c *gin.Context) {
	var pages []models.Page
	query := config.DB.Preload("Author").Preload("Parent").Preload("Children")

	// Filter by status
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// Filter by author
	if authorID := c.Query("author_id"); authorID != "" {
		query = query.Where("author_id = ?", authorID)
	}

	// Filter by parent (for hierarchical pages)
	if parentID := c.Query("parent_id"); parentID != "" {
		if parentID == "null" {
			query = query.Where("parent_id IS NULL")
		} else {
			query = query.Where("parent_id = ?", parentID)
		}
	}

	// Search by title or content
	if search := c.Query("search"); search != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Order by
	orderBy := c.DefaultQuery("order_by", "created_at")
	orderDir := c.DefaultQuery("order_dir", "desc")
	query = query.Order(orderBy + " " + orderDir)

	if err := query.Find(&pages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pages": pages})
}

// GetPage returns a single page by ID
func GetPage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}

	var page models.Page
	if err := config.DB.Preload("Author").Preload("Parent").Preload("Children").First(&page, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch page"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"page": page})
}

// GetPageBySlug returns a single page by slug (for public access)
func GetPageBySlug(c *gin.Context) {
	slug := c.Param("slug")

	var page models.Page
	if err := config.DB.Preload("Author").Preload("Parent").Preload("Children").
		Where("slug = ? AND status = ?", slug, models.PageStatusPublished).
		First(&page).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch page"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"page": page})
}

// CreatePage creates a new page
func CreatePage(c *gin.Context) {
	var page models.Page
	if err := c.ShouldBindJSON(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get current user from context (set by auth middleware)
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	page.AuthorID = user.(*shared_models.User).ID

	// Generate slug if not provided
	if page.Slug == "" {
		page.Slug = generateSlug(page.Title)
	}

	// Ensure slug is unique
	page.Slug = ensureUniqueSlug(page.Slug, 0)

	if err := config.DB.Create(&page).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create page"})
		return
	}

	// Reload with relationships
	config.DB.Preload("Author").Preload("Parent").Preload("Children").First(&page, page.ID)

	c.JSON(http.StatusCreated, gin.H{"page": page})
}

// UpdatePage updates an existing page
func UpdatePage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}

	var page models.Page
	if err := config.DB.First(&page, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch page"})
		}
		return
	}

	var updateData models.Page
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	page.Title = updateData.Title
	page.Content = updateData.Content
	page.Excerpt = updateData.Excerpt
	page.Template = updateData.Template
	page.Status = updateData.Status
	page.FeaturedImage = updateData.FeaturedImage
	page.MetaTitle = updateData.MetaTitle
	page.MetaDescription = updateData.MetaDescription
	page.MetaKeywords = updateData.MetaKeywords
	page.ParentID = updateData.ParentID
	page.SortOrder = updateData.SortOrder

	// Handle slug update
	if updateData.Slug != "" && updateData.Slug != page.Slug {
		page.Slug = ensureUniqueSlug(updateData.Slug, page.ID)
	}

	// Handle status changes
	if updateData.Status == models.PageStatusPublished && page.PublishedAt == nil {
		page.Publish()
	} else if updateData.Status == models.PageStatusDraft {
		page.Unpublish()
	}

	if err := config.DB.Save(&page).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update page"})
		return
	}

	// Reload with relationships
	config.DB.Preload("Author").Preload("Parent").Preload("Children").First(&page, page.ID)

	c.JSON(http.StatusOK, gin.H{"page": page})
}

// DeletePage deletes a page
func DeletePage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}

	var page models.Page
	if err := config.DB.First(&page, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch page"})
		}
		return
	}

	if err := config.DB.Delete(&page).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Page deleted successfully"})
}

// Helper functions

// generateSlug creates a URL-friendly slug from a title
func generateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")
	// Remove special characters (keep only alphanumeric and hyphens)
	var result strings.Builder
	for _, char := range slug {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' {
			result.WriteRune(char)
		}
	}
	return result.String()
}

// ensureUniqueSlug ensures the slug is unique by appending a number if needed
func ensureUniqueSlug(slug string, excludeID uint) string {
	baseSlug := slug
	counter := 1

	for {
		var count int64
		query := config.DB.Model(&models.Page{}).Where("slug = ?", slug)
		if excludeID > 0 {
			query = query.Where("id != ?", excludeID)
		}
		query.Count(&count)

		if count == 0 {
			return slug
		}

		slug = baseSlug + "-" + strconv.Itoa(counter)
		counter++
	}
}
