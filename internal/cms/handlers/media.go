package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gcx-cms/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// MediaFile represents a media file record for API responses
type MediaFile struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Size      int64     `json:"size"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// MediaFileRecord represents the database model for media files
type MediaFileRecord struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	OriginalName string     `gorm:"type:longtext;not null" json:"original_name"`
	Filename     string     `gorm:"type:varchar(191);not null;uniqueIndex" json:"filename"`
	URL          string     `gorm:"type:longtext;not null" json:"url"`
	ThumbnailURL *string    `gorm:"type:longtext" json:"thumbnail_url"`
	MimeType     string     `gorm:"type:longtext;not null" json:"mime_type"`
	Size         int64      `gorm:"type:bigint;not null" json:"size"`
	UploadedBy   uint       `gorm:"not null" json:"uploaded_by"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `gorm:"index" json:"deleted_at"`
}

// TableName specifies the table name for MediaFileRecord
func (MediaFileRecord) TableName() string {
	return "media_files"
}

// GetMedia returns all media files (protected)
func GetMedia(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mediaRecords []MediaFileRecord
	if err := db.Where("deleted_at IS NULL").Order("created_at DESC").Find(&mediaRecords).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch media files: " + err.Error(),
		})
		return
	}

	// Convert to API response format
	files := make([]MediaFile, len(mediaRecords))
	for i, record := range mediaRecords {
		files[i] = MediaFile{
			ID:        record.Filename,
			Name:      record.OriginalName,
			URL:       record.URL,
			Size:      record.Size,
			Type:      record.MimeType,
			CreatedAt: record.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    files,
		"count":   len(files),
	})
}

// UploadFile handles file uploads (protected)
func UploadFile(c *gin.Context) {
	// Get the file from the request
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "No file provided",
		})
		return
	}
	defer file.Close()

	// Validate file type
	contentType := header.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Only image files are allowed",
		})
		return
	}

	// Validate file size (10MB limit)
	if header.Size > 10*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "File size too large (max 10MB)",
		})
		return
	}

	// Get the upload type from form (optional) - this will be the S3 folder
	uploadType := c.PostForm("type")
	if uploadType == "" {
		uploadType = "cms"
	}

	// Initialize S3 service
	s3Service, err := services.NewS3Service()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to initialize S3 service: " + err.Error(),
		})
		return
	}

	// Upload to S3
	fileURL, err := s3Service.UploadFile(file, header, uploadType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to upload file to S3: " + err.Error(),
		})
		return
	}

	// Generate filename from URL
	filename := filepath.Base(fileURL)

	// Get the current user ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "User ID not found in context",
		})
		return
	}

	// Save to database
	db := c.MustGet("db").(*gorm.DB)
	mediaRecord := MediaFileRecord{
		OriginalName: header.Filename,
		Filename:     filename,
		URL:          fileURL,
		MimeType:     contentType,
		Size:         header.Size,
		UploadedBy:   userID.(uint),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := db.Create(&mediaRecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to save file record to database: " + err.Error(),
		})
		return
	}

	// Return the file info
	mediaFile := MediaFile{
		ID:        filename,
		Name:      header.Filename,
		URL:       fileURL,
		Size:      header.Size,
		Type:      contentType,
		CreatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    []MediaFile{mediaFile},
		"message": "File uploaded successfully to S3",
	})
}

// GetMediaFile returns a single media file by ID (protected)
func GetMediaFile(c *gin.Context) {
	id := c.Param("id")

	// Check if file exists
	filePath := filepath.Join("./uploads/images", id)
	if info, err := os.Stat(filePath); err == nil && !info.IsDir() {
		// Return file info
		ext := strings.ToLower(filepath.Ext(id))
		mediaFile := MediaFile{
			ID:        id,
			Name:      id,
			URL:       "/uploads/images/" + id,
			Size:      info.Size(),
			Type:      "image/" + strings.TrimPrefix(ext, "."),
			CreatedAt: info.ModTime(),
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    mediaFile,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "File not found",
		})
	}
}

// GetDocuments returns all document files (protected)
func GetDocuments(c *gin.Context) {
	files := []MediaFile{}

	// Scan uploads/documents directory
	uploadPath := "./uploads/documents"
	if entries, err := os.ReadDir(uploadPath); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() {
				info, err := entry.Info()
				if err != nil {
					continue
				}

				// Include document files (PDF, DOC, DOCX, etc.)
				ext := strings.ToLower(filepath.Ext(info.Name()))
				if ext == ".pdf" || ext == ".doc" || ext == ".docx" || ext == ".txt" || ext == ".xlsx" || ext == ".xls" {
					files = append(files, MediaFile{
						ID:        info.Name(),
						Name:      info.Name(),
						URL:       "/uploads/documents/" + info.Name(),
						Size:      info.Size(),
						Type:      getMimeType(ext),
						CreatedAt: info.ModTime(),
					})
				}
			}
		}
	}

	// Also scan the uploads/publications directory for existing files
	publicationsPath := "./uploads/publications"
	if entries, err := os.ReadDir(publicationsPath); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() {
				info, err := entry.Info()
				if err != nil {
					continue
				}

				// Include document files (PDF, DOC, DOCX, etc.)
				ext := strings.ToLower(filepath.Ext(info.Name()))
				if ext == ".pdf" || ext == ".doc" || ext == ".docx" || ext == ".txt" || ext == ".xlsx" || ext == ".xls" {
					files = append(files, MediaFile{
						ID:        info.Name(),
						Name:      info.Name(),
						URL:       "/uploads/publications/" + info.Name(),
						Size:      info.Size(),
						Type:      getMimeType(ext),
						CreatedAt: info.ModTime(),
					})
				}
			}
		}
	}

	// Also scan the uploads/careers directory for existing files
	careersPath := "./uploads/careers"
	if entries, err := os.ReadDir(careersPath); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() {
				info, err := entry.Info()
				if err != nil {
					continue
				}

				// Include document files (PDF, DOC, DOCX, etc.)
				ext := strings.ToLower(filepath.Ext(info.Name()))
				if ext == ".pdf" || ext == ".doc" || ext == ".docx" || ext == ".txt" || ext == ".xlsx" || ext == ".xls" {
					files = append(files, MediaFile{
						ID:        info.Name(),
						Name:      info.Name(),
						URL:       "/uploads/careers/" + info.Name(),
						Size:      info.Size(),
						Type:      getMimeType(ext),
						CreatedAt: info.ModTime(),
					})
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    files,
	})
}

// UploadDocument handles document file uploads (protected)
func UploadDocument(c *gin.Context) {
	// Get the file from the request
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "No file provided",
		})
		return
	}
	defer file.Close()

	// Get folder from form data
	folder := c.PostForm("folder")
	if folder == "" {
		folder = "publications"
	}

	// Ensure folder is under uploads
	if !strings.HasPrefix(folder, "uploads/") {
		folder = "uploads/" + folder
	}

	// Validate file type
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowedExts := []string{".pdf", ".doc", ".docx", ".txt", ".xlsx", ".xls"}
	allowed := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			allowed = true
			break
		}
	}

	if !allowed {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "File type not allowed. Only PDF, DOC, DOCX, TXT, XLS, XLSX files are allowed",
		})
		return
	}

	// Validate file size (10MB limit)
	if header.Size > 10*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "File size too large (max 10MB)",
		})
		return
	}

	// Initialize S3 service
	s3Service, err := services.NewS3Service()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to initialize S3 service: " + err.Error(),
		})
		return
	}

	// Upload to S3
	fileURL, err := s3Service.UploadFile(file, header, folder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to upload file to S3: " + err.Error(),
		})
		return
	}

	// Generate filename from URL
	filename := filepath.Base(fileURL)

	// Return the file info
	mediaFile := MediaFile{
		ID:        filename,
		Name:      header.Filename,
		URL:       fileURL,
		Size:      header.Size,
		Type:      getMimeType(ext),
		CreatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    mediaFile,
		"message": "File uploaded successfully",
	})
}

// getMimeType returns the MIME type for a file extension
func getMimeType(ext string) string {
	switch ext {
	case ".pdf":
		return "application/pdf"
	case ".doc":
		return "application/msword"
	case ".docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case ".txt":
		return "text/plain"
	case ".xls":
		return "application/vnd.ms-excel"
	case ".xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	default:
		return "application/octet-stream"
	}
}

// DeleteMedia deletes a media file (protected)
func DeleteMedia(c *gin.Context) {
	id := c.Param("id")

	// Validate filename (prevent directory traversal)
	if strings.Contains(id, "..") || strings.Contains(id, "/") || strings.Contains(id, "\\") {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid filename",
		})
		return
	}

	// Get database connection
	db := c.MustGet("db").(*gorm.DB)

	// Find the media record in database
	var mediaRecord MediaFileRecord
	if err := db.Where("filename = ? AND deleted_at IS NULL", id).First(&mediaRecord).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "File not found in database",
		})
		return
	}

	// Check if it's an S3 URL or local URL
	if strings.HasPrefix(mediaRecord.URL, "https://") {
		// Delete from S3
		s3Service, err := services.NewS3Service()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to initialize S3 service: " + err.Error(),
			})
			return
		}

		// Extract S3 key from URL
		// URL format: https://bucket.s3.region.amazonaws.com/key
		urlParts := strings.Split(mediaRecord.URL, "/")
		if len(urlParts) < 4 {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid S3 URL format",
			})
			return
		}
		s3Key := strings.Join(urlParts[3:], "/")

		if err := s3Service.DeleteFile(s3Key); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to delete file from S3: " + err.Error(),
			})
			return
		}
	} else {
		// Delete from local filesystem
		filePath := filepath.Join("./uploads/images", id)
		if err := os.Remove(filePath); err != nil {
			// Don't fail if local file doesn't exist (might have been moved to S3)
			fmt.Printf("Warning: Could not delete local file %s: %v\n", filePath, err)
		}
	}

	// Soft delete from database
	if err := db.Model(&mediaRecord).Update("deleted_at", time.Now()).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete file record from database: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "File deleted successfully",
	})
}
