package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile returns the current user's profile (protected)
func GetProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	// Return the user profile directly
	c.JSON(http.StatusOK, user)
}

// UpdateProfile updates the current user's profile (protected)
func UpdateProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update profile endpoint - to be implemented"})
}

// ChangePassword changes the current user's password (protected)
func ChangePassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Change password endpoint - to be implemented"})
}
