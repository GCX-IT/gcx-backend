package routes

import (
	"gcx-cms/internal/handlers"
	"gcx-cms/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

// SetupUploadRoutes configures file upload routes
func SetupUploadRoutes(r *gin.Engine) {
	// Initialize upload handler
	uploadHandler := handlers.NewUploadHandler()

	// Upload routes group with authentication
	upload := r.Group("/api/upload")
	upload.Use(middleware.AuthMiddleware())
	{
		// General file upload
		upload.POST("/file", uploadHandler.UploadFile)

		// File management
		upload.DELETE("/file/:filename", uploadHandler.DeleteFile)
	}

	// Public download route (no authentication required)
	r.GET("/api/files/download", uploadHandler.DownloadFile) // GET /api/files/download?key=path/to/file&action=view|download

	// Note: Public file serving is handled by r.Static("/uploads", "./uploads") in main.go
	// Removing the conflicting route to avoid panic
}
