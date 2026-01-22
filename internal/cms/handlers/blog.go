package handlers

import (
	"encoding/json"
	"fmt"
	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/config"
	shared_models "gcx-cms/internal/shared/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// GetPublicPosts returns all published blog posts for public viewing
func GetPublicPosts(c *gin.Context) {
	var posts []models.BlogPost

	// Get only published posts, ordered by publish date
	if err := config.DB.Preload("Author").
		Where("status = ? AND published_at IS NOT NULL", models.StatusPublished).
		Order("published_at DESC").
		Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	// Filter out posts with missing required fields
	var validPosts []models.BlogPost
	for _, post := range posts {
		// Check if post has required fields
		if post.Title != "" && post.Content != "" && post.AuthorID != 0 {
			validPosts = append(validPosts, post)
		} else {
			// Log invalid posts for debugging
			log.Printf("⚠️ Skipping invalid post ID %d: Title='%s', Content length=%d, AuthorID=%d",
				post.ID, post.Title, len(post.Content), post.AuthorID)
		}
	}

	// Transform posts to be more frontend-friendly
	var frontendPosts []gin.H
	for _, post := range validPosts {
		frontendPost := gin.H{
			"id":             post.ID,
			"title":          post.Title,
			"slug":           post.Slug,
			"content":        post.Content,
			"excerpt":        post.Excerpt,
			"featured_image": post.FeaturedImage,
			"status":         post.Status,
			"published_at":   post.PublishedAt,
			"created_at":     post.CreatedAt,
			"updated_at":     post.UpdatedAt,
			"author":         post.Author.Name, // Just the author name, not the full object
			"author_id":      post.AuthorID,
		}
		frontendPosts = append(frontendPosts, frontendPost)
	}

	// Return transformed posts
	c.JSON(http.StatusOK, frontendPosts)
}

// GetPublicPost returns a single published blog post by slug
func GetPublicPost(c *gin.Context) {
	slug := c.Param("slug")
	c.JSON(http.StatusOK, gin.H{"message": "Public post endpoint", "slug": slug, "status": "to be implemented"})
}

// GetAllPosts returns all blog posts for CMS management (protected)
func GetAllPosts(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	currentUser := user.(*shared_models.User)

	var posts []models.BlogPost
	query := config.DB.Preload("Author").Order("created_at DESC")

	// Non-admin users can only see their own posts
	if currentUser.Role != shared_models.RoleAdmin {
		query = query.Where("author_id = ?", currentUser.ID)
	}

	if err := query.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	// Return the posts array directly (not wrapped in a message)
	c.JSON(http.StatusOK, posts)
}

// GetPost returns a single blog post by ID for editing (protected)
func GetPost(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get post endpoint", "id": id, "status": "to be implemented"})
}

// CreatePost creates a new blog post (protected)
func CreatePost(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	currentUser := user.(*shared_models.User)

	var req struct {
		Title         string   `json:"title" binding:"required"`
		Content       string   `json:"content" binding:"required"`
		Excerpt       string   `json:"excerpt"`
		FeaturedImage *string  `json:"featured_image"`
		Tags          []string `json:"tags"`
		Status        string   `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// Validate required fields
	if req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	if req.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content is required"})
		return
	}
	
	// Auto-generate excerpt if empty
	if req.Excerpt == "" {
		req.Excerpt = generateExcerptFromContent(req.Content)
	}

	// Generate slug from title
	slug := generateBlogSlug(req.Title)

	// Check if slug already exists
	var existingPost models.BlogPost
	if err := config.DB.Where("slug = ?", slug).First(&existingPost).Error; err == nil {
		// Make slug unique by appending timestamp
		slug = slug + "-" + fmt.Sprintf("%d", time.Now().Unix())
	}

	// Set status
	status := models.StatusDraft
	if req.Status != "" {
		switch req.Status {
		case string(models.StatusPublished), string(models.StatusDraft), string(models.StatusPrivate):
			status = models.BlogStatus(req.Status)
		}
	}

	post := models.BlogPost{
		Title:         req.Title,
		Slug:          slug,
		Content:       req.Content,
		Excerpt:       req.Excerpt,
		FeaturedImage: req.FeaturedImage,
		Status:        status,
		AuthorID:      currentUser.ID,
	}

	// Convert tags to JSON
	if len(req.Tags) > 0 {
		if tagsJSON, err := json.Marshal(req.Tags); err == nil {
			post.Tags = datatypes.JSON(tagsJSON)
		}
	}

	if status == models.StatusPublished {
		post.Publish()
	}

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	// Load author for response
	config.DB.Preload("Author").First(&post, post.ID)

	c.JSON(http.StatusCreated, post)
}

// UpdatePost updates an existing blog post (protected)
func UpdatePost(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	currentUser := user.(*shared_models.User)
	id := c.Param("id")

	// Parse ID
	postID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Find existing post
	var post models.BlogPost
	if err := config.DB.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch post"})
		}
		return
	}

	// Check if user has permission to update this post
	if post.AuthorID != currentUser.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this post"})
		return
	}

	var req struct {
		Title         string   `json:"title"`
		Content       string   `json:"content"`
		Excerpt       string   `json:"excerpt"`
		FeaturedImage *string  `json:"featured_image"`
		Tags          []string `json:"tags"`
		Status        string   `json:"status"`
		Slug          string   `json:"slug"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// Update fields if provided
	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
		// Auto-generate excerpt if content changed and excerpt is empty
		if req.Excerpt == "" {
			post.Excerpt = generateExcerptFromContent(req.Content)
		} else {
			post.Excerpt = req.Excerpt
		}
	} else if req.Excerpt != "" {
		post.Excerpt = req.Excerpt
	}
	if req.FeaturedImage != nil {
		post.FeaturedImage = req.FeaturedImage
	}

	// Handle slug update
	if req.Slug != "" && req.Slug != post.Slug {
		// Check if new slug is unique
		var existingPost models.BlogPost
		if err := config.DB.Where("slug = ? AND id != ?", req.Slug, post.ID).First(&existingPost).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Slug already exists"})
			return
		}
		post.Slug = req.Slug
	}

	// Handle tags update
	if len(req.Tags) > 0 {
		if tagsJSON, err := json.Marshal(req.Tags); err == nil {
			post.Tags = datatypes.JSON(tagsJSON)
		}
	}

	// Handle status update
	if req.Status != "" {
		switch req.Status {
		case string(models.StatusPublished):
			post.Publish()
		case string(models.StatusDraft):
			post.Unpublish()
		case string(models.StatusPrivate):
			post.Status = models.StatusPrivate
			post.PublishedAt = nil
		}
	}

	// Save changes
	if err := config.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	// Reload with relationships
	config.DB.Preload("Author").First(&post, post.ID)

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "post": post})
}

// DeletePost deletes a blog post (protected)
func DeletePost(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	currentUser := user.(*shared_models.User)
	id := c.Param("id")

	// Parse ID
	postID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Find existing post
	var post models.BlogPost
	if err := config.DB.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch post"})
		}
		return
	}

	// Check if user has permission to delete this post
	if post.AuthorID != currentUser.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this post"})
		return
	}

	// Delete the post
	if err := config.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

// Helper function to generate slug from title
func generateBlogSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	var result strings.Builder
	for _, r := range slug {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// Helper function to generate excerpt from content
// Strips HTML tags and returns first 160 characters
func generateExcerptFromContent(content string) string {
	// Simple HTML tag removal
	excerpt := strings.ReplaceAll(content, "<p>", " ")
	excerpt = strings.ReplaceAll(excerpt, "</p>", " ")
	excerpt = strings.ReplaceAll(excerpt, "<br>", " ")
	excerpt = strings.ReplaceAll(excerpt, "<br/>", " ")
	excerpt = strings.ReplaceAll(excerpt, "<br />", " ")
	excerpt = strings.ReplaceAll(excerpt, "<div>", " ")
	excerpt = strings.ReplaceAll(excerpt, "</div>", " ")
	excerpt = strings.ReplaceAll(excerpt, "<h1>", " ")
	excerpt = strings.ReplaceAll(excerpt, "</h1>", " ")
	excerpt = strings.ReplaceAll(excerpt, "<h2>", " ")
	excerpt = strings.ReplaceAll(excerpt, "</h2>", " ")
	excerpt = strings.ReplaceAll(excerpt, "<h3>", " ")
	excerpt = strings.ReplaceAll(excerpt, "</h3>", " ")
	
	// Remove all remaining HTML tags using a simple regex-like approach
	var result strings.Builder
	inTag := false
	for _, r := range excerpt {
		if r == '<' {
			inTag = true
			continue
		}
		if r == '>' {
			inTag = false
			continue
		}
		if !inTag {
			result.WriteRune(r)
		}
	}
	
	excerpt = result.String()
	// Clean up multiple spaces
	excerpt = strings.ReplaceAll(excerpt, "  ", " ")
	excerpt = strings.TrimSpace(excerpt)
	
	// Limit to 160 characters
	if len(excerpt) > 160 {
		excerpt = excerpt[:160]
		// Try to cut at a word boundary
		if lastSpace := strings.LastIndex(excerpt, " "); lastSpace > 100 {
			excerpt = excerpt[:lastSpace]
		}
		excerpt += "..."
	}
	
	return excerpt
}
