package handlers

import (
"fmt"
"io"
"net/http"
"net/url"
"os"
"path/filepath"
"strconv"
"strings"
"time"

"gcx-cms/internal/services"

"github.com/gin-gonic/gin"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
return &UploadHandler{}
}

// UploadFile handles file uploads
func (h *UploadHandler) UploadFile(c *gin.Context) {
// Parse multipart form with max memory of 32MB
err := c.Request.ParseMultipartForm(32 << 20)
if err != nil {
c.JSON(http.StatusBadRequest, gin.H{
"success": false,
"error":   "Failed to parse multipart form",
})
return
}

file, handler, err := c.Request.FormFile("file")
if err != nil {
c.JSON(http.StatusBadRequest, gin.H{
"success": false,
"error":   "No file provided",
})
return
}
defer file.Close()

// Validate file size (max 10MB)
maxSize := int64(10 << 20) // 10MB
if handler.Size > maxSize {
c.JSON(http.StatusBadRequest, gin.H{
"success": false,
"error":   "File size exceeds maximum limit of 10MB",
})
return
}

// Create uploads directory if it doesn't exist
uploadsDir := "uploads"
if err := os.MkdirAll(uploadsDir, 0755); err != nil {
c.JSON(http.StatusInternalServerError, gin.H{
"success": false,
"error":   "Failed to create uploads directory",
})
return
}

// Generate unique filename
timestamp := strconv.FormatInt(time.Now().Unix(), 10)
ext := filepath.Ext(handler.Filename)
filename := fmt.Sprintf("%s_%s%s", timestamp, sanitizeFilename(strings.TrimSuffix(handler.Filename, ext)), ext)
filepath := filepath.Join(uploadsDir, filename)

// Create the file
dst, err := os.Create(filepath)
if err != nil {
c.JSON(http.StatusInternalServerError, gin.H{
"success": false,
"error":   "Failed to create file",
})
return
}
defer dst.Close()

// Copy file content
if _, err := io.Copy(dst, file); err != nil {
c.JSON(http.StatusInternalServerError, gin.H{
"success": false,
"error":   "Failed to save file",
})
return
}

// Return success response
c.JSON(http.StatusOK, gin.H{
"success": true,
"message": "File uploaded successfully",
"data": gin.H{
"filename":     filename,
"original_name": handler.Filename,
"size":        handler.Size,
"url":         fmt.Sprintf("/uploads/%s", filename),
},
})
}

// ServeFile serves uploaded files
func (h *UploadHandler) ServeFile(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Filename is required",
		})
		return
	}

	filepath := filepath.Join("uploads", filename)
	
	// Check if file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "File not found",
		})
		return
	}

	c.File(filepath)
}

// DeleteFile removes an uploaded file
func (h *UploadHandler) DeleteFile(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Filename is required",
		})
		return
	}

	filepath := filepath.Join("uploads", filename)
	
	// Check if file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "File not found",
		})
		return
	}

	// Delete the file
	if err := os.Remove(filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "File deleted successfully",
	})
}

// DownloadFile streams a file from S3 or local filesystem
func (h *UploadHandler) DownloadFile(c *gin.Context) {
	// Get the file key from query parameter (URL encoded)
	fileKey := c.Query("key")
	if fileKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "File key is required",
		})
		return
	}

	// Decode URL-encoded filename
	decodedKey, err := url.QueryUnescape(fileKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid file key",
		})
		return
	}

	// Check if it's a local path first
	var fileContent []byte
	var contentType string
	var filename string
	var localPath string

	if strings.HasPrefix(decodedKey, "/uploads/") {
		// Local file path - try to serve from filesystem
		// First try: ./uploads/uploads/contracts/... (nested path)
		localPath = "." + decodedKey
		fileContent, err = os.ReadFile(localPath)
		
		// If not found, try alternative path with double /uploads
		if err != nil {
			localPath = "./uploads" + decodedKey
			fileContent, err = os.ReadFile(localPath)
		}
		
		if err == nil {
			// Successfully read local file
			filename = filepath.Base(localPath)
			contentType = getContentType(filename)
		} else {
			// Local file not found, try S3
			fileContent = nil
		}
	}

	// If local file not found, try S3
	if fileContent == nil {
		s3Key := decodedKey
		if strings.Contains(decodedKey, "amazonaws.com/") {
			// Extract key after domain
			parts := strings.Split(decodedKey, "amazonaws.com/")
			if len(parts) == 2 {
				s3Key = parts[1]
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   "Invalid S3 URL format",
				})
				return
			}
		} else if strings.HasPrefix(decodedKey, "/uploads/") {
			// Convert local path to S3 key
			// /uploads/contracts/filename.pdf -> contracts/filename.pdf
			filename = filepath.Base(decodedKey)
			s3Key = "contracts/" + filename
		}

		// Initialize S3 service
		s3Service, err := services.NewS3Service()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "S3 service not available",
			})
			return
		}

		// Get file from S3
		fileContent, contentType, err = s3Service.GetFile(s3Key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to retrieve file: " + err.Error(),
			})
			return
		}

		// Extract filename from key
		filename = filepath.Base(s3Key)
	}

	// Set headers based on whether user wants to download or view
	action := c.DefaultQuery("action", "download")
	if action == "view" {
		// Display inline (browser will display if it can)
		c.Header("Content-Disposition", "inline; filename="+filename)
	} else {
		// Force download
		c.Header("Content-Disposition", "attachment; filename="+filename)
	}

	c.Header("Content-Type", contentType)
	c.Data(http.StatusOK, contentType, fileContent)
}

// getContentType returns the appropriate content type for a file
func getContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".pdf":
		return "application/pdf"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".txt":
		return "text/plain"
	case ".csv":
		return "text/csv"
	case ".zip":
		return "application/zip"
	default:
		return "application/octet-stream"
	}
}

// sanitizeFilename removes unsafe characters from filename
func sanitizeFilename(filename string) string {
// Remove or replace unsafe characters
unsafe := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|", " "}
sanitized := filename

for _, char := range unsafe {
sanitized = strings.ReplaceAll(sanitized, char, "_")
}

return sanitized
}
