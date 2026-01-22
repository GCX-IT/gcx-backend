package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupAllRoutes configures all application routes
func SetupAllRoutes(r *gin.Engine) {
	// Setup different route modules
	SetupAuthRoutes(r)
	SetupCMSRoutes(r)
	SetupMarketDataRoutes(r)
	SetupUploadRoutes(r)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "service": "GCX Market Data Platform API"})
	})
}
