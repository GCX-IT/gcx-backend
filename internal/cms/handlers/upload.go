package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"gcx-cms/internal/services"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	s3Service *services.S3Service
}

func NewUploadHandler() (*UploadHandler, error) {
	s3Service, err := services.NewS3Service()
	if err != nil {
		return nil, err
	}

	return &UploadHandler{
		s3Service: s3Service,
	}, nil
}

// UploadImage handles image uploads for blog posts and user avatars
func (h *UploadHandler) UploadImage(c *gin.Context) {
	// Check if user is authenticated
	_, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	// Parse multipart form
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large or invalid form"})
		return
	}

	// Get file from form
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	// Validate file type
	if !isValidImageType(header.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only JPG, PNG, GIF, and WebP are allowed"})
		return
	}

	// Validate file size (5MB max)
	if header.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large. Maximum size is 5MB"})
		return
	}

	// Get the upload type from form (optional) - this will be the S3 folder
	uploadType := c.PostForm("type") // "avatar", "featured_image", "content", "partners", "team", etc.
	if uploadType == "" {
		uploadType = "cms"
	}

	// Upload to S3
	fileURL, err := h.s3Service.UploadFile(file, header, uploadType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to S3: " + err.Error()})
		return
	}

	// Extract filename from URL for response
	filename := filepath.Base(fileURL)

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "File uploaded successfully",
		"url":      fileURL,
		"filename": filename,
		"type":     uploadType,
		"size":     header.Size,
	})
}

// UploadMultipleImages handles multiple image uploads
func (h *UploadHandler) UploadMultipleImages(c *gin.Context) {
	// Check if user is authenticated
	_, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	// Parse multipart form
	err := c.Request.ParseMultipartForm(50 << 20) // 50MB max for multiple files
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Form too large or invalid"})
		return
	}

	form := c.Request.MultipartForm
	files := form.File["images"]

	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
		return
	}

	if len(files) > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Maximum 10 files allowed"})
		return
	}

	// Get the upload type from form (optional) - this will be the S3 folder
	uploadType := c.PostForm("type") // "avatar", "featured_image", "content", "partners", "team", etc.
	if uploadType == "" {
		uploadType = "cms"
	}

	var uploadedFiles []gin.H
	var errors []string

	for _, fileHeader := range files {
		// Validate file type
		if !isValidImageType(fileHeader.Filename) {
			errors = append(errors, fmt.Sprintf("%s: Invalid file type", fileHeader.Filename))
			continue
		}

		// Validate file size (5MB max per file)
		if fileHeader.Size > 5*1024*1024 {
			errors = append(errors, fmt.Sprintf("%s: File too large", fileHeader.Filename))
			continue
		}

		// Open uploaded file
		file, err := fileHeader.Open()
		if err != nil {
			errors = append(errors, fmt.Sprintf("%s: Failed to open file", fileHeader.Filename))
			continue
		}

		// Upload to S3
		fileURL, err := h.s3Service.UploadFile(file, fileHeader, uploadType)
		file.Close() // Close the file after upload

		if err != nil {
			errors = append(errors, fmt.Sprintf("%s: Failed to upload to S3: %s", fileHeader.Filename, err.Error()))
			continue
		}

		// Add to successful uploads
		filename := filepath.Base(fileURL)
		uploadedFiles = append(uploadedFiles, gin.H{
			"original_name": fileHeader.Filename,
			"filename":      filename,
			"url":           fileURL,
			"size":          fileHeader.Size,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success":        len(uploadedFiles) > 0,
		"message":        fmt.Sprintf("Uploaded %d of %d files", len(uploadedFiles), len(files)),
		"uploaded_files": uploadedFiles,
		"errors":         errors,
	})
}

// DeleteImage deletes an uploaded image from S3
func (h *UploadHandler) DeleteImage(c *gin.Context) {
	// Check if user is authenticated
	_, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	// Get S3 key from query parameter or path
	s3Key := c.Query("key")
	if s3Key == "" {
		// Fallback to filename for backward compatibility
		filename := c.Param("filename")
		if filename == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "S3 key or filename is required"})
			return
		}
		s3Key = filename
	}

	// Delete file from S3
	err := h.s3Service.DeleteFile(s3Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file from S3: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "File deleted successfully from S3",
	})
}

// isValidImageType checks if the file extension is valid for images
func isValidImageType(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	validExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}
	return validExts[ext]
}

// UploadDocument handles document uploads (PDF, DOC, DOCX)
func (h *UploadHandler) UploadDocument(c *gin.Context) {
	// Check if user is authenticated
	_, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	// Parse multipart form
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large or invalid form"})
		return
	}

	// Get file from form
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	// Validate file type
	if !isValidDocumentType(header.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only PDF, DOC, and DOCX are allowed"})
		return
	}

	// Validate file size (10MB max)
	if header.Size > 10*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large. Maximum size is 10MB"})
		return
	}

	// Get the upload type from form
	uploadType := c.PostForm("type")
	if uploadType == "" {
		uploadType = "documents"
	}

	// Upload to S3
	fileURL, err := h.s3Service.UploadFile(file, header, uploadType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to S3: " + err.Error()})
		return
	}

	// Extract filename from URL for response
	filename := filepath.Base(fileURL)

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "Document uploaded successfully",
		"url":      fileURL,
		"filename": filename,
		"type":     uploadType,
		"size":     header.Size,
	})
}

// UploadVideo handles video uploads
func (h *UploadHandler) UploadVideo(c *gin.Context) {
	// Check if user is authenticated
	_, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	// Parse multipart form
	err := c.Request.ParseMultipartForm(100 << 20) // 100MB max for videos
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large or invalid form"})
		return
	}

	// Get file from form
	file, header, err := c.Request.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	// Validate file type
	if !isValidVideoType(header.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only MP4 and WebM are allowed"})
		return
	}

	// Validate file size (100MB max)
	if header.Size > 100*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large. Maximum size is 100MB"})
		return
	}

	// Get the upload type from form
	uploadType := c.PostForm("type")
	if uploadType == "" {
		uploadType = "videos"
	}

	// Upload to S3
	fileURL, err := h.s3Service.UploadFile(file, header, uploadType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to S3: " + err.Error()})
		return
	}

	// Extract filename from URL for response
	filename := filepath.Base(fileURL)

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "Video uploaded successfully",
		"url":      fileURL,
		"filename": filename,
		"type":     uploadType,
		"size":     header.Size,
	})
}

// DeleteFile deletes a file from S3 (alias for DeleteImage for consistency)
func (h *UploadHandler) DeleteFile(c *gin.Context) {
	h.DeleteImage(c)
}

// ListFiles lists files in S3 bucket
func (h *UploadHandler) ListFiles(c *gin.Context) {
	// Check if user is authenticated
	_, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	// Get prefix from query parameter
	prefix := c.Query("prefix")
	if prefix == "" {
		prefix = "cms" // Default prefix
	}

	// List files from S3
	files, err := h.s3Service.ListFiles(prefix)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list files from S3: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"files":   files,
		"prefix":  prefix,
		"count":   len(files),
	})
}

// isValidDocumentType checks if the file extension is valid for documents
func isValidDocumentType(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	validExts := map[string]bool{
		".pdf":  true,
		".doc":  true,
		".docx": true,
	}
	return validExts[ext]
}

// isValidVideoType checks if the file extension is valid for videos
func isValidVideoType(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	validExts := map[string]bool{
		".mp4":  true,
		".webm": true,
	}
	return validExts[ext]
}

// generateRandomString generates a random string for filename uniqueness
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}
