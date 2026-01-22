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

// GetGalleries retrieves all active galleries (public endpoint)
func GetGalleries(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	category := c.Query("category")
	featured := c.Query("featured")

	var galleries []models.PhotoGallery
	query := db.Where("is_active = ?", true)

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if featured == "true" {
		query = query.Where("is_featured = ?", true)
	}

	if err := query.Order("sort_order ASC, date DESC").Find(&galleries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve galleries"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": galleries})
}

// GetGallery retrieves a single gallery with photos (public endpoint)
func GetGallery(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var gallery models.PhotoGallery
	if err := db.Where("id = ? AND is_active = ?", id, true).
		Preload("Photos", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC, created_at DESC")
		}).
		First(&gallery).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Gallery not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve gallery"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gallery})
}

// GetGalleryBySlug retrieves a gallery by slug (public endpoint)
func GetGalleryBySlug(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	gallerySlug := c.Param("slug")

	var gallery models.PhotoGallery
	if err := db.Where("slug = ? AND is_active = ?", gallerySlug, true).
		Preload("Photos", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC, created_at DESC")
		}).
		First(&gallery).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Gallery not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve gallery"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gallery})
}

// GetAllGalleries retrieves all galleries for CMS (protected endpoint)
func GetAllGalleries(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var galleries []models.PhotoGallery
	if err := db.Order("created_at DESC").Find(&galleries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve galleries"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": galleries})
}

// CreateGallery creates a new photo gallery (protected endpoint)
func CreateGallery(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var gallery models.PhotoGallery
	if err := c.ShouldBindJSON(&gallery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate slug from title if not provided
	if gallery.Slug == "" {
		gallery.Slug = slug.Make(gallery.Title)
	} else {
		gallery.Slug = slug.Make(gallery.Slug)
	}

	// Check if slug already exists
	var existingGallery models.PhotoGallery
	if err := db.Where("slug = ?", gallery.Slug).First(&existingGallery).Error; err == nil {
		// Slug exists, append timestamp
		gallery.Slug = gallery.Slug + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	}

	if err := db.Create(&gallery).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": "Gallery with this slug already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create gallery"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Gallery created successfully",
		"data":    gallery,
	})
}

// UpdateGallery updates an existing gallery (protected endpoint)
func UpdateGallery(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var gallery models.PhotoGallery
	if err := db.First(&gallery, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Gallery not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve gallery"})
		return
	}

	var updates models.PhotoGallery
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&gallery).Updates(&updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update gallery"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Gallery updated successfully",
		"data":    gallery,
	})
}

// DeleteGallery deletes a gallery (soft delete) (protected endpoint)
func DeleteGallery(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var gallery models.PhotoGallery
	if err := db.First(&gallery, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Gallery not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve gallery"})
		return
	}

	if err := db.Delete(&gallery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete gallery"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gallery deleted successfully"})
}

// AddPhotoToGallery adds a photo to a gallery (protected endpoint)
func AddPhotoToGallery(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	galleryID := c.Param("id")

	var photo models.GalleryPhoto
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify gallery exists
	var gallery models.PhotoGallery
	if err := db.First(&gallery, galleryID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gallery not found"})
		return
	}

	gid, _ := strconv.ParseUint(galleryID, 10, 32)
	photo.GalleryID = uint(gid)

	if err := db.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add photo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Photo added successfully",
		"data":    photo,
	})
}

// UpdateGalleryPhoto updates a photo (protected endpoint)
func UpdateGalleryPhoto(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var photo models.GalleryPhoto
	if err := db.First(&photo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve photo"})
		return
	}

	var updates models.GalleryPhoto
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&photo).Updates(&updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo updated successfully",
		"data":    photo,
	})
}

// DeleteGalleryPhoto deletes a photo (soft delete) (protected endpoint)
func DeleteGalleryPhoto(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var photo models.GalleryPhoto
	if err := db.First(&photo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve photo"})
		return
	}

	if err := db.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}

// GetGalleryPhotos retrieves all photos in a gallery (public endpoint)
func GetGalleryPhotos(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	galleryID := c.Param("id")

	var photos []models.GalleryPhoto
	if err := db.Where("gallery_id = ?", galleryID).
		Order("sort_order ASC, created_at DESC").
		Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve photos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": photos})
}
