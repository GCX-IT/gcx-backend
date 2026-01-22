package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/config"
	"gorm.io/gorm"
)

// GetBoardMembers returns all board members (protected)
func GetBoardMembers(c *gin.Context) {
	var boardMembers []models.BoardMember
	
	if err := config.DB.Order("order_index ASC, created_at ASC").Find(&boardMembers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch board members",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    boardMembers,
	})
}

// CreateBoardMember creates a new board member (protected)
func CreateBoardMember(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Position    string `json:"position" binding:"required"`
		Image       string `json:"image"`
		Description string `json:"description" binding:"required"`
		LinkedIn    string `json:"linkedin"`
		Facebook    string `json:"facebook"`
		Instagram   string `json:"instagram"`
		OrderIndex  int    `json:"order_index"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Get the next order index if not provided
	if input.OrderIndex == 0 {
		var maxOrder int
		config.DB.Model(&models.BoardMember{}).Select("COALESCE(MAX(order_index), 0)").Scan(&maxOrder)
		input.OrderIndex = maxOrder + 1
	}

	boardMember := models.BoardMember{
		Name:        input.Name,
		Position:    input.Position,
		Image:       input.Image,
		Description: input.Description,
		LinkedIn:    input.LinkedIn,
		Facebook:    input.Facebook,
		Instagram:   input.Instagram,
		OrderIndex:  input.OrderIndex,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := config.DB.Create(&boardMember).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create board member",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    boardMember,
		"message": "Board member created successfully",
	})
}

// GetBoardMember returns a single board member by ID (protected)
func GetBoardMember(c *gin.Context) {
	id := c.Param("id")
	
	var boardMember models.BoardMember
	if err := config.DB.First(&boardMember, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Board member not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to fetch board member",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    boardMember,
	})
}

// UpdateBoardMember updates a board member (protected)
func UpdateBoardMember(c *gin.Context) {
	id := c.Param("id")
	
	var boardMember models.BoardMember
	if err := config.DB.First(&boardMember, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Board member not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to fetch board member",
			})
		}
		return
	}

	var input struct {
		Name        string `json:"name"`
		Position    string `json:"position"`
		Image       string `json:"image"`
		Description string `json:"description"`
		LinkedIn    string `json:"linkedin"`
		Facebook    string `json:"facebook"`
		Instagram   string `json:"instagram"`
		OrderIndex  int    `json:"order_index"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Update fields
	updates := map[string]interface{}{
		"updated_at": time.Now(),
	}

	if input.Name != "" {
		updates["name"] = input.Name
	}
	if input.Position != "" {
		updates["position"] = input.Position
	}
	if input.Image != "" {
		updates["image"] = input.Image
	}
	if input.Description != "" {
		updates["description"] = input.Description
	}
	if input.LinkedIn != "" {
		updates["linkedin"] = input.LinkedIn
	}
	if input.Facebook != "" {
		updates["facebook"] = input.Facebook
	}
	if input.Instagram != "" {
		updates["instagram"] = input.Instagram
	}
	if input.OrderIndex > 0 {
		updates["order_index"] = input.OrderIndex
	}

	if err := config.DB.Model(&boardMember).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update board member",
		})
		return
	}

	// Fetch updated record
	config.DB.First(&boardMember, "id = ?", id)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    boardMember,
		"message": "Board member updated successfully",
	})
}

// DeleteBoardMember deletes a board member (protected)
func DeleteBoardMember(c *gin.Context) {
	id := c.Param("id")
	
	var boardMember models.BoardMember
	if err := config.DB.First(&boardMember, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Board member not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to fetch board member",
			})
		}
		return
	}

	if err := config.DB.Delete(&boardMember).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete board member",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Board member deleted successfully",
	})
}

// ReorderBoardMembers updates the order of board members (protected)
func ReorderBoardMembers(c *gin.Context) {
	var input struct {
		Members []struct {
			ID    uint `json:"id"`
			Order int  `json:"order"`
		} `json:"members" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Update order for each member
	for _, member := range input.Members {
		if err := config.DB.Model(&models.BoardMember{}).
			Where("id = ?", member.ID).
			Update("order_index", member.Order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to update member order",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Board member order updated successfully",
	})
}
