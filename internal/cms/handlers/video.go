package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"gcx-cms/internal/cms/models"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

// GetVideoLibraries retrieves all active video libraries (public endpoint)
func GetVideoLibraries(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	category := c.Query("category")
	featured := c.Query("featured")

	var libraries []models.VideoLibrary
	query := db.Where("is_active = ?", true)

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if featured == "true" {
		query = query.Where("is_featured = ?", true)
	}

	if err := query.Order("sort_order ASC, date DESC").Find(&libraries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve video libraries"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": libraries})
}

// GetVideoLibrary retrieves a single video library with videos (public endpoint)
func GetVideoLibrary(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var library models.VideoLibrary
	if err := db.Where("id = ? AND is_active = ?", id, true).
		Preload("Videos", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC, created_at DESC")
		}).
		First(&library).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Video library not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve video library"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": library})
}

// GetVideoLibraryBySlug retrieves a video library by slug (public endpoint)
func GetVideoLibraryBySlug(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	librarySlug := c.Param("slug")

	var library models.VideoLibrary
	if err := db.Where("slug = ? AND is_active = ?", librarySlug, true).
		Preload("Videos", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC, created_at DESC")
		}).
		First(&library).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Video library not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve video library"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": library})
}

// GetAllVideoLibraries retrieves all video libraries for CMS (protected endpoint)
func GetAllVideoLibraries(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var libraries []models.VideoLibrary
	if err := db.Order("created_at DESC").Find(&libraries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve video libraries"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": libraries})
}

// CreateVideoLibrary creates a new video library (protected endpoint)
func CreateVideoLibrary(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var library models.VideoLibrary
	if err := c.ShouldBindJSON(&library); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate slug from title if not provided
	if library.Slug == "" {
		library.Slug = slug.Make(library.Title)
	} else {
		library.Slug = slug.Make(library.Slug)
	}

	// Check if slug already exists
	var existingLibrary models.VideoLibrary
	if err := db.Where("slug = ?", library.Slug).First(&existingLibrary).Error; err == nil {
		// Slug exists, append timestamp
		library.Slug = library.Slug + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	}

	if err := db.Create(&library).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": "Video library with this slug already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create video library"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Video library created successfully",
		"data":    library,
	})
}

// UpdateVideoLibrary updates an existing video library (protected endpoint)
func UpdateVideoLibrary(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var library models.VideoLibrary
	if err := db.First(&library, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Video library not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve video library"})
		return
	}

	var updates models.VideoLibrary
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&library).Updates(&updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update video library"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Video library updated successfully",
		"data":    library,
	})
}

// DeleteVideoLibrary deletes a video library (soft delete) (protected endpoint)
func DeleteVideoLibrary(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var library models.VideoLibrary
	if err := db.First(&library, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Video library not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve video library"})
		return
	}

	if err := db.Delete(&library).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete video library"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video library deleted successfully"})
}

// AddVideoToLibrary adds a video to a library (protected endpoint)
func AddVideoToLibrary(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	libraryID := c.Param("id")

	var video models.LibraryVideo
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify library exists
	var library models.VideoLibrary
	if err := db.First(&library, libraryID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video library not found"})
		return
	}

	libID, _ := strconv.ParseUint(libraryID, 10, 32)
	video.LibraryID = uint(libID)

	if err := db.Create(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add video"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Video added successfully",
		"data":    video,
	})
}

// UpdateLibraryVideo updates a video (protected endpoint)
func UpdateLibraryVideo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var video models.LibraryVideo
	if err := db.First(&video, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve video"})
		return
	}

	var updates models.LibraryVideo
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&video).Updates(&updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update video"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Video updated successfully",
		"data":    video,
	})
}

// DeleteLibraryVideo deletes a video (soft delete) (protected endpoint)
func DeleteLibraryVideo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var video models.LibraryVideo
	if err := db.First(&video, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve video"})
		return
	}

	if err := db.Delete(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete video"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video deleted successfully"})
}

// GetLibraryVideos retrieves all videos in a library (public endpoint)
func GetLibraryVideos(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	libraryID := c.Param("id")

	var videos []models.LibraryVideo
	if err := db.Where("library_id = ?", libraryID).
		Order("sort_order ASC, created_at DESC").
		Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve videos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": videos})
}

// TrackVideoView increments view count for a video (public endpoint)
func TrackVideoView(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var video models.LibraryVideo
	if err := db.First(&video, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve video"})
		return
	}

	if err := video.IncrementViewCount(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to track view"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "View tracked successfully",
		"data":    video,
	})
}
