package handlers

import (
	"net/http"
	"strconv"

	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetMenus returns all menus
func GetMenus(c *gin.Context) {
	var menus []models.Menu
	query := config.DB.Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Order("sort_order ASC")
	}).Preload("Items.Children", func(db *gorm.DB) *gorm.DB {
		return db.Order("sort_order ASC")
	})

	// Filter by location
	if location := c.Query("location"); location != "" {
		query = query.Where("location = ?", location)
	}

	// Filter by active status
	if isActive := c.Query("is_active"); isActive != "" {
		if isActive == "true" {
			query = query.Where("is_active = ?", true)
		} else if isActive == "false" {
			query = query.Where("is_active = ?", false)
		}
	}

	query = query.Order("name ASC")

	if err := query.Find(&menus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"menus": menus})
}

// GetMenu returns a single menu by ID
func GetMenu(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}

	var menu models.Menu
	if err := config.DB.Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Order("sort_order ASC")
	}).Preload("Items.Children", func(db *gorm.DB) *gorm.DB {
		return db.Order("sort_order ASC")
	}).First(&menu, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

// GetMenuByLocation returns a menu by location (for frontend use)
func GetMenuByLocation(c *gin.Context) {
	location := c.Param("location")

	var menu models.Menu
	if err := config.DB.Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Where("is_active = ?", true).Order("sort_order ASC")
	}).Preload("Items.Children", func(db *gorm.DB) *gorm.DB {
		return db.Where("is_active = ?", true).Order("sort_order ASC")
	}).Where("location = ? AND is_active = ?", location, true).First(&menu).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

// CreateMenu creates a new menu
func CreateMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create menu"})
		return
	}

	// Reload with relationships
	config.DB.Preload("Items").First(&menu, menu.ID)

	c.JSON(http.StatusCreated, gin.H{"menu": menu})
}

// UpdateMenu updates an existing menu
func UpdateMenu(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}

	var menu models.Menu
	if err := config.DB.First(&menu, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu"})
		}
		return
	}

	var updateData models.Menu
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	menu.Name = updateData.Name
	menu.Location = updateData.Location
	menu.IsActive = updateData.IsActive

	if err := config.DB.Save(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update menu"})
		return
	}

	// Reload with relationships
	config.DB.Preload("Items").First(&menu, menu.ID)

	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

// DeleteMenu deletes a menu
func DeleteMenu(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}

	var menu models.Menu
	if err := config.DB.First(&menu, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu"})
		}
		return
	}

	if err := config.DB.Delete(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete menu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}

// Menu Item handlers

// GetMenuItems returns all menu items for a specific menu
func GetMenuItems(c *gin.Context) {
	menuID, err := strconv.ParseUint(c.Param("menu_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}

	var items []models.MenuItem
	if err := config.DB.Where("menu_id = ?", menuID).
		Preload("Children", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC")
		}).
		Order("sort_order ASC").
		Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

// CreateMenuItem creates a new menu item
func CreateMenuItem(c *gin.Context) {
	menuID, err := strconv.ParseUint(c.Param("menu_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}

	var item models.MenuItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.MenuID = uint(menuID)

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create menu item"})
		return
	}

	// Reload with relationships
	config.DB.Preload("Menu").Preload("Parent").Preload("Children").First(&item, item.ID)

	c.JSON(http.StatusCreated, gin.H{"item": item})
}

// UpdateMenuItem updates an existing menu item
func UpdateMenuItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu item ID"})
		return
	}

	var item models.MenuItem
	if err := config.DB.First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Menu item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu item"})
		}
		return
	}

	var updateData models.MenuItem
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	item.Label = updateData.Label
	item.URL = updateData.URL
	item.Target = updateData.Target
	item.IconClass = updateData.IconClass
	item.ParentID = updateData.ParentID
	item.SortOrder = updateData.SortOrder
	item.IsActive = updateData.IsActive

	if err := config.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update menu item"})
		return
	}

	// Reload with relationships
	config.DB.Preload("Menu").Preload("Parent").Preload("Children").First(&item, item.ID)

	c.JSON(http.StatusOK, gin.H{"item": item})
}

// DeleteMenuItem deletes a menu item
func DeleteMenuItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu item ID"})
		return
	}

	var item models.MenuItem
	if err := config.DB.First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Menu item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu item"})
		}
		return
	}

	if err := config.DB.Delete(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete menu item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu item deleted successfully"})
}
