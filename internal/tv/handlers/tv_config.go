package tv_handlers

import (
	"encoding/json"
	"net/http"

	"gcx-cms/internal/services"
	"gcx-cms/internal/shared/database"
	tv_models "gcx-cms/internal/tv/models"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// getOrCreateConfig fetches the single TV config row, creating it with defaults if absent.
func getOrCreateConfig(db *gorm.DB) (tv_models.TVConfig, error) {
	var cfg tv_models.TVConfig
	err := db.First(&cfg).Error
	if err != nil {
		cfg = tv_models.TVConfig{
			AutoNext:           true,
			Loop:               true,
			Playlist:           datatypes.JSON([]byte("[]")),
			Images:             datatypes.JSON([]byte("[]")),
			VideoDuration:      60,
			MarketDataDuration: 10,
			ImageDuration:      120,
			EnableRotation:     false,
		}
		if createErr := db.Create(&cfg).Error; createErr != nil {
			return cfg, createErr
		}
	}
	// Ensure JSON arrays are never null
	if cfg.Playlist == nil {
		cfg.Playlist = datatypes.JSON([]byte("[]"))
	}
	if cfg.Images == nil {
		cfg.Images = datatypes.JSON([]byte("[]"))
	}
	return cfg, nil
}

// GetTVConfig returns the current TV config (public).
// GET /api/tv/config
func GetTVConfig(c *gin.Context) {
	db := database.GetDB()
	cfg, err := getOrCreateConfig(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get TV config"})
		return
	}
	c.JSON(http.StatusOK, cfg)
}

// SaveTVConfig accepts a partial TVConfig JSON body and merges it into the stored config.
// POST /api/tv/config
func SaveTVConfig(c *gin.Context) {
	db := database.GetDB()

	current, err := getOrCreateConfig(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load TV config"})
		return
	}

	// Parse request as a raw field map so we can detect which keys were sent
	var patch map[string]json.RawMessage
	if err := c.ShouldBindJSON(&patch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if v, ok := patch["nowPlaying"]; ok {
		if string(v) == "null" {
			current.NowPlaying = nil
		} else {
			var s string
			if json.Unmarshal(v, &s) == nil {
				current.NowPlaying = &s
			}
		}
	}
	if v, ok := patch["nowPlayingId"]; ok {
		if string(v) == "null" {
			current.NowPlayingId = nil
		} else {
			var s string
			if json.Unmarshal(v, &s) == nil {
				current.NowPlayingId = &s
			}
		}
	}
	if v, ok := patch["autoNext"]; ok {
		var b bool
		if json.Unmarshal(v, &b) == nil {
			current.AutoNext = b
		}
	}
	if v, ok := patch["loop"]; ok {
		var b bool
		if json.Unmarshal(v, &b) == nil {
			current.Loop = b
		}
	}
	if v, ok := patch["enableRotation"]; ok {
		var b bool
		if json.Unmarshal(v, &b) == nil {
			current.EnableRotation = b
		}
	}
	if v, ok := patch["videoDuration"]; ok {
		var n int
		if json.Unmarshal(v, &n) == nil {
			current.VideoDuration = n
		}
	}
	if v, ok := patch["marketDataDuration"]; ok {
		var n int
		if json.Unmarshal(v, &n) == nil {
			current.MarketDataDuration = n
		}
	}
	if v, ok := patch["imageDuration"]; ok {
		var n int
		if json.Unmarshal(v, &n) == nil {
			current.ImageDuration = n
		}
	}
	if v, ok := patch["playlist"]; ok {
		current.Playlist = datatypes.JSON(v)
	}
	if v, ok := patch["images"]; ok {
		current.Images = datatypes.JSON(v)
	}

	// Save selects all columns including zero-value booleans and null pointers
	if err := db.Save(&current).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save TV config"})
		return
	}

	c.JSON(http.StatusOK, current)
}

// UploadTVImage uploads an image file to S3 under gcx-tv/images/ and returns the public URL.
// POST /api/tv/upload/image
func UploadTVImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}
	defer file.Close()

	s3Svc, err := services.NewS3Service()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "S3 service unavailable: " + err.Error()})
		return
	}

	url, err := s3Svc.UploadFile(file, header, "gcx-tv/images")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload failed: " + err.Error()})
		return
	}

	name := c.PostForm("name")
	if name == "" {
		name = header.Filename
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"url":     url,
		"name":    name,
	})
}

// UploadTVVideo uploads a video file to S3 under gcx-tv/videos/ and returns the public URL.
// POST /api/tv/upload/video
func UploadTVVideo(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}
	defer file.Close()

	s3Svc, err := services.NewS3Service()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "S3 service unavailable: " + err.Error()})
		return
	}

	url, err := s3Svc.UploadFile(file, header, "gcx-tv/videos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload failed: " + err.Error()})
		return
	}

	name := c.PostForm("name")
	if name == "" {
		name = header.Filename
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"url":     url,
		"name":    name,
		"filename": header.Filename,
	})
}
