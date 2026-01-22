package handlers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetTeamMembers returns all team members (protected)
func GetTeamMembers(c *gin.Context) {
	var teamMembers []models.TeamMember

	// Get type filter from query parameter
	teamType := c.Query("type")

	// Let custom sorting handle all ordering
	query := config.DB

	if teamType != "" {
		query = query.Where("type = ?", teamType)
	}

	if err := query.Find(&teamMembers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch team members",
		})
		return
	}

	// Debug: Log before sorting
	fmt.Printf("🔍 Before sorting - Found %d team members of type '%s'\n", len(teamMembers), teamType)
	for i, member := range teamMembers {
		fmt.Printf("  %d: %s (%s) - order_index: %d\n", i+1, member.Name, member.Title, member.OrderIndex)
	}

	// Sort team members by intelligent hierarchy
	sort.Slice(teamMembers, func(i, j int) bool {
		a, b := teamMembers[i], teamMembers[j]

		// Check if either is CEO - CEO always comes first regardless of order_index
		// More specific detection to avoid matching "Deputy Chief Executive Officer"
		aIsCEO := (strings.Contains(strings.ToLower(a.Title), "chief executive officer") && !strings.Contains(strings.ToLower(a.Title), "deputy")) || strings.Contains(strings.ToLower(a.Title), "ceo")
		bIsCEO := (strings.Contains(strings.ToLower(b.Title), "chief executive officer") && !strings.Contains(strings.ToLower(b.Title), "deputy")) || strings.Contains(strings.ToLower(b.Title), "ceo")

		// Debug logging
		fmt.Printf("🔍 Sorting: %s (%s) vs %s (%s) - aIsCEO: %v, bIsCEO: %v\n",
			a.Name, a.Title, b.Name, b.Title, aIsCEO, bIsCEO)

		if aIsCEO && !bIsCEO {
			fmt.Printf("✅ CEO %s comes first\n", a.Name)
			return true // CEO comes first
		}
		if !aIsCEO && bIsCEO {
			fmt.Printf("✅ CEO %s comes first\n", b.Name)
			return false // CEO comes first
		}

		// Also check for Chairman in board members
		if a.Type == "board" {
			aIsChairman := strings.Contains(strings.ToLower(a.Title), "chairman") || strings.Contains(strings.ToLower(a.Title), "board chairman")
			bIsChairman := strings.Contains(strings.ToLower(b.Title), "chairman") || strings.Contains(strings.ToLower(b.Title), "board chairman")

			if aIsChairman && !bIsChairman && !bIsCEO {
				return true // Chairman comes first (after CEO)
			}
			if !aIsChairman && bIsChairman && !aIsCEO {
				return false // Chairman comes first (after CEO)
			}
		}

		// If both or neither are CEO, then sort by order_index
		if a.OrderIndex != 0 && b.OrderIndex != 0 {
			return a.OrderIndex < b.OrderIndex
		}
		// If order_index is missing for one, prioritize the one with order_index
		if a.OrderIndex != 0 {
			return true
		}
		if b.OrderIndex != 0 {
			return false
		}

		// If both missing order_index, sort by title hierarchy
		aTitle := strings.ToLower(a.Title)
		bTitle := strings.ToLower(b.Title)

		// Define hierarchy based on team type
		var hierarchy []string
		switch a.Type {
		case "board":
			hierarchy = []string{"chief executive officer", "ceo", "chairman", "board chairman", "vice chairman", "board secretary", "treasurer", "director"}
		case "executive":
			hierarchy = []string{"chief executive officer", "ceo", "deputy chief executive officer", "deputy ceo", "chief operating officer", "coo", "chief financial officer", "cfo", "president", "managing director", "executive director"}
		case "functional":
			hierarchy = []string{"head of operations", "head of", "director of", "manager of", "senior manager", "manager", "coordinator", "specialist", "officer"}
		default:
			hierarchy = []string{}
		}

		aIndex := -1
		bIndex := -1
		for idx, title := range hierarchy {
			if strings.Contains(aTitle, title) {
				aIndex = idx
				break
			}
		}
		for idx, title := range hierarchy {
			if strings.Contains(bTitle, title) {
				bIndex = idx
				break
			}
		}

		// If both have hierarchy titles, sort by hierarchy
		if aIndex != -1 && bIndex != -1 {
			return aIndex < bIndex
		}
		// If only one has hierarchy title, prioritize it
		if aIndex != -1 {
			return true
		}
		if bIndex != -1 {
			return false
		}

		// If neither has hierarchy title, sort by name
		return a.Name < b.Name
	})

	// Debug logging for team member ordering
	if len(teamMembers) > 0 {
		teamType := teamMembers[0].Type
		for i, member := range teamMembers {
			fmt.Printf("✅ %s #%d: %s (%s) - order_index: %d\n",
				strings.Title(teamType), i+1, member.Name, member.Title, member.OrderIndex)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    teamMembers,
	})
}

// GetTeamMember returns a single team member by ID (protected)
func GetTeamMember(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid team member ID",
		})
		return
	}

	var teamMember models.TeamMember
	if err := config.DB.First(&teamMember, uint(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Team member not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to fetch team member",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    teamMember,
	})
}

// CreateTeamMember creates a new team member (protected)
func CreateTeamMember(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Image       string `json:"image"`
		Type        string `json:"type" binding:"required,oneof=board executive functional"`
		OrderIndex  int    `json:"order_index"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// If no order index provided, set it to the next available index for this type
	if input.OrderIndex == 0 {
		var maxOrder int
		config.DB.Model(&models.TeamMember{}).Where("type = ?", input.Type).Select("COALESCE(MAX(order_index), 0)").Scan(&maxOrder)
		input.OrderIndex = maxOrder + 1
	}

	teamMember := models.TeamMember{
		Name:        input.Name,
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		Type:        input.Type,
		OrderIndex:  input.OrderIndex,
	}

	if err := config.DB.Create(&teamMember).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create team member",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    teamMember,
		"message": "Team member created successfully",
	})
}

// UpdateTeamMember updates an existing team member (protected)
func UpdateTeamMember(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid team member ID",
		})
		return
	}

	var input struct {
		Name        string `json:"name"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Image       string `json:"image"`
		Type        string `json:"type" binding:"omitempty,oneof=board executive functional"`
		OrderIndex  int    `json:"order_index"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	var teamMember models.TeamMember
	if err := config.DB.First(&teamMember, uint(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Team member not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to fetch team member",
			})
		}
		return
	}

	// Update fields if provided
	updates := make(map[string]interface{})
	if input.Name != "" {
		updates["name"] = input.Name
	}
	if input.Title != "" {
		updates["title"] = input.Title
	}
	if input.Description != "" {
		updates["description"] = input.Description
	}
	if input.Image != "" {
		updates["image"] = input.Image
	}
	if input.Type != "" {
		updates["type"] = input.Type
	}
	if input.OrderIndex != 0 {
		updates["order_index"] = input.OrderIndex
	}

	if err := config.DB.Model(&teamMember).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update team member",
		})
		return
	}

	// Fetch updated team member
	config.DB.First(&teamMember, uint(id))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    teamMember,
		"message": "Team member updated successfully",
	})
}

// DeleteTeamMember deletes a team member (protected)
func DeleteTeamMember(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid team member ID",
		})
		return
	}

	var teamMember models.TeamMember
	if err := config.DB.First(&teamMember, uint(id)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Team member not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to fetch team member",
			})
		}
		return
	}

	if err := config.DB.Delete(&teamMember).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete team member",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Team member deleted successfully",
	})
}

// ReorderTeamMembers reorders team members (protected)
func ReorderTeamMembers(c *gin.Context) {
	var input struct {
		TeamType string `json:"team_type" binding:"required,oneof=board executive functional"`
		Members  []struct {
			ID         uint `json:"id" binding:"required"`
			OrderIndex int  `json:"order_index" binding:"required"`
		} `json:"members" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Update order indices in a transaction
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, member := range input.Members {
		if err := tx.Model(&models.TeamMember{}).Where("id = ? AND type = ?", member.ID, input.TeamType).Update("order_index", member.OrderIndex).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to reorder team members",
			})
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to commit reorder changes",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Team members reordered successfully",
	})
}
