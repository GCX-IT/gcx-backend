package handlers

import (
	"net/http"
	"time"

	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/config"
	shared_models "gcx-cms/internal/shared/models"

	"github.com/gin-gonic/gin"
)

// DashboardStats represents the dashboard statistics
type DashboardStats struct {
	TotalPosts     int `json:"total_posts"`
	PublishedPosts int `json:"published_posts"`
	DraftPosts     int `json:"draft_posts"`
	MediaFiles     int `json:"media_files"`
	TotalPages     int `json:"total_pages"`
	PublishedPages int `json:"published_pages"`
	DraftPages     int `json:"draft_pages"`
}

// DashboardActivity represents a recent activity item
type DashboardActivity struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Color     string `json:"color"`
}

// GetDashboardStats returns dashboard statistics from real data
func GetDashboardStats(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	currentUser := user.(*shared_models.User)

	// Initialize stats
	stats := DashboardStats{}

	// Count blog posts
	var totalPosts, publishedPosts, draftPosts int64

	// Base query for posts
	postQuery := config.DB.Model(&models.BlogPost{})

	// Non-admin users can only see their own posts
	if currentUser.Role != shared_models.RoleAdmin {
		postQuery = postQuery.Where("author_id = ?", currentUser.ID)
	}

	// Count total posts
	postQuery.Count(&totalPosts)
	stats.TotalPosts = int(totalPosts)

	// Count published posts
	postQuery.Where("status = ?", models.StatusPublished).Count(&publishedPosts)
	stats.PublishedPosts = int(publishedPosts)

	// Count draft posts
	postQuery.Where("status = ?", models.StatusDraft).Count(&draftPosts)
	stats.DraftPosts = int(draftPosts)

	// Count pages
	var totalPages, publishedPages, draftPages int64

	// Base query for pages
	pageQuery := config.DB.Model(&models.Page{})

	// Non-admin users can only see their own pages
	if currentUser.Role != shared_models.RoleAdmin {
		pageQuery = pageQuery.Where("author_id = ?", currentUser.ID)
	}

	// Count total pages
	pageQuery.Count(&totalPages)
	stats.TotalPages = int(totalPages)

	// Count published pages
	pageQuery.Where("status = ?", models.PageStatusPublished).Count(&publishedPages)
	stats.PublishedPages = int(publishedPages)

	// Count draft pages
	pageQuery.Where("status = ?", models.PageStatusDraft).Count(&draftPages)
	stats.DraftPages = int(draftPages)

	// Count media files (simplified - you may want to enhance this)
	// For now, we'll count files in the uploads directory or use a media table
	stats.MediaFiles = 0 // This would need to be implemented based on your media storage

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    stats,
	})
}

// GetDashboardActivity returns recent activity from real data
func GetDashboardActivity(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	currentUser := user.(*shared_models.User)

	var activities []DashboardActivity

	// Get recent blog posts
	var recentPosts []models.BlogPost
	postQuery := config.DB.Preload("Author").Order("updated_at DESC").Limit(5)

	// Non-admin users can only see their own posts
	if currentUser.Role != shared_models.RoleAdmin {
		postQuery = postQuery.Where("author_id = ?", currentUser.ID)
	}

	if err := postQuery.Find(&recentPosts).Error; err == nil {
		for i, post := range recentPosts {
			activityType := "updated"
			color := "blue"

			if post.Status == models.StatusPublished {
				activityType = "published"
				color = "green"
			} else if post.Status == models.StatusDraft {
				activityType = "draft"
				color = "yellow"
			}

			activities = append(activities, DashboardActivity{
				ID:        i + 1,
				Type:      activityType,
				Message:   "Blog post: " + post.Title,
				Timestamp: post.UpdatedAt.Format(time.RFC3339),
				Color:     color,
			})
		}
	}

	// Get recent pages
	var recentPages []models.Page
	pageQuery := config.DB.Preload("Author").Order("updated_at DESC").Limit(3)

	// Non-admin users can only see their own pages
	if currentUser.Role != shared_models.RoleAdmin {
		pageQuery = pageQuery.Where("author_id = ?", currentUser.ID)
	}

	if err := pageQuery.Find(&recentPages).Error; err == nil {
		for _, page := range recentPages {
			activityType := "updated"
			color := "purple"

			if page.Status == models.PageStatusPublished {
				activityType = "published"
				color = "green"
			} else if page.Status == models.PageStatusDraft {
				activityType = "draft"
				color = "yellow"
			}

			activities = append(activities, DashboardActivity{
				ID:        len(activities) + 1,
				Type:      activityType,
				Message:   "Page: " + page.Title,
				Timestamp: page.UpdatedAt.Format(time.RFC3339),
				Color:     color,
			})
		}
	}

	// Sort activities by timestamp (most recent first)
	// Simple bubble sort for small arrays
	for i := 0; i < len(activities)-1; i++ {
		for j := 0; j < len(activities)-i-1; j++ {
			time1, _ := time.Parse(time.RFC3339, activities[j].Timestamp)
			time2, _ := time.Parse(time.RFC3339, activities[j+1].Timestamp)
			if time1.Before(time2) {
				activities[j], activities[j+1] = activities[j+1], activities[j]
			}
		}
	}

	// Limit to 10 most recent activities
	if len(activities) > 10 {
		activities = activities[:10]
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    activities,
	})
}
