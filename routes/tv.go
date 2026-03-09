package routes

import (
	tv_handlers "gcx-cms/internal/tv/handlers"
	"gcx-cms/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

// SetupTVRoutes registers all GCX TV endpoints under /api/tv/
func SetupTVRoutes(r *gin.Engine) {
	tv := r.Group("/api/tv")
	tv.Use(middleware.DatabaseMiddleware())

	// Public: reading the config and playing video requires no auth
	tv.GET("/config", tv_handlers.GetTVConfig)

	// Protected: writing config and uploading assets requires auth
	protected := tv.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/config", tv_handlers.SaveTVConfig)
		protected.POST("/upload/image", tv_handlers.UploadTVImage)
		protected.POST("/upload/video", tv_handlers.UploadTVVideo)
	}
}
