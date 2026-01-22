package handlers

import (
"fmt"
"io"
"net/http"
"os"
"path/filepath"
"strconv"
"strings"
"time"

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
