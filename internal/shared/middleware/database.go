package middleware

import (
	"gcx-cms/internal/shared/config"

	"github.com/gin-gonic/gin"
)

// DatabaseMiddleware injects the database connection into the Gin context
func DatabaseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", config.DB)
		c.Next()
	}
}
