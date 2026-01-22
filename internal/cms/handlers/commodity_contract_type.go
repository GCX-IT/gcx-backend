package handlers

import (
	"net/http"

	"gcx-cms/internal/cms/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllCommodityContractTypes retrieves all contract types for a commodity
func GetAllCommodityContractTypes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	commodityID := c.Param("commodityId")

	var contractTypes []models.CommodityContractType
	if err := db.Where("commodity_id = ? AND is_active = ?", commodityID, true).
		Order("sort_order ASC").
		Find(&contractTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch contract types",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    contractTypes,
	})
}

// GetCommodityContractType retrieves a single contract type by ID
func GetCommodityContractType(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var contractType models.CommodityContractType
	if err := db.Preload("Commodity").First(&contractType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Contract type not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    contractType,
	})
}

// CreateCommodityContractType creates a new contract type
func CreateCommodityContractType(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var contractType models.CommodityContractType
	if err := c.ShouldBindJSON(&contractType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid contract type data",
		})
		return
	}

	if err := db.Create(&contractType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create contract type",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    contractType,
	})
}

// UpdateCommodityContractType updates an existing contract type
func UpdateCommodityContractType(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var contractType models.CommodityContractType
	if err := db.First(&contractType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Contract type not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&contractType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid contract type data",
		})
		return
	}

	if err := db.Save(&contractType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update contract type",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    contractType,
	})
}

// DeleteCommodityContractType deletes a contract type
func DeleteCommodityContractType(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var contractType models.CommodityContractType
	if err := db.First(&contractType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Contract type not found",
		})
		return
	}

	if err := db.Delete(&contractType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete contract type",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Contract type deleted successfully",
	})
}

// GetCommoditiesWithContractTypes retrieves all commodities with their contract types
func GetCommoditiesWithContractTypes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var commodities []models.Commodity
	if err := db.Preload("ContractTypes", "is_active = ?", true).
		Order("name ASC").
		Find(&commodities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch commodities with contract types",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    commodities,
	})
}

// UpdateContractTypeSortOrder updates the sort order of contract types
func UpdateContractTypeSortOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type SortOrderUpdate struct {
		ID        uint `json:"id" binding:"required"`
		SortOrder int  `json:"sort_order" binding:"required"`
	}

	var updates []SortOrderUpdate
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid sort order data",
		})
		return
	}

	// Update each contract type's sort order
	for _, update := range updates {
		if err := db.Model(&models.CommodityContractType{}).
			Where("id = ?", update.ID).
			Update("sort_order", update.SortOrder).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to update sort order",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Sort order updated successfully",
	})
}
