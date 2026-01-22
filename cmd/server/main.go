package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gcx-cms/internal/shared/config"
	"gcx-cms/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	config.InitDB()

	// Create upload directories
	uploadDirs := []string{"./uploads", "./uploads/images", "./uploads/videos", "./uploads/documents"}
	for _, dir := range uploadDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Printf("Warning: Failed to create upload directory %s: %v", dir, err)
		}
	}

	// Initialize Gin router
	r := gin.Default()

	// Set max multipart memory to 10MB (for file uploads)
	r.MaxMultipartMemory = 10 << 20 // 10 MB

	// CORS middleware
	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// Serve static files (uploads)
	r.Static("/uploads", "./uploads")

	// Setup all application routes
	routes.SetupAllRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 GCX Market Data Platform API Server starting on port %s", port)
	log.Printf("🔐 JWT Authentication enabled")
	log.Printf("📊 Connected to database")
	log.Printf("✅ Modular architecture ready")
	log.Printf("📈 Market Data Platform routes active")
	log.Printf("📝 CMS routes active")

	r.Run(":" + port)
}
