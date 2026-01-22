package routes

import (
	auth_handlers "gcx-cms/internal/shared/auth"
	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes configures all authentication related routes
func SetupAuthRoutes(r *gin.Engine) {
	// Auth API group
	auth := r.Group("/api/auth")

	// Public auth routes (no authentication required)
	{
		auth.POST("/login", auth_handlers.LoginHandler)
		auth.POST("/register", auth_handlers.RegisterHandler)
	}
}
