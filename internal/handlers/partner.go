package handlers

import (
	"net/http"
	"strconv"

	"gcx-cms/internal/models"
	"gcx-cms/internal/services"

	"github.com/gin-gonic/gin"
)

type PartnerHandler struct {
	partnerService *services.PartnerService
}

func NewPartnerHandler(partnerService *services.PartnerService) *PartnerHandler {
	return &PartnerHandler{partnerService: partnerService}
}

// CreatePartner creates a new partner
func (h *PartnerHandler) CreatePartner(c *gin.Context) {
	var partner models.Partner
	if err := c.ShouldBindJSON(&partner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	if err := h.partnerService.CreatePartner(&partner); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create partner"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": partner})
}

// GetPartner retrieves a partner by ID
func (h *PartnerHandler) GetPartner(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid partner ID"})
		return
	}

	partner, err := h.partnerService.GetPartnerByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "Partner not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": partner})
}

// GetAllPartners retrieves all partners with pagination and filtering
func (h *PartnerHandler) GetAllPartners(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")
	category := c.Query("category")
	status := c.Query("status")

	partners, total, err := h.partnerService.GetAllPartners(page, limit, search, category, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to retrieve partners"})
		return
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    partners,
		"pagination": gin.H{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}

// UpdatePartner updates a partner
func (h *PartnerHandler) UpdatePartner(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid partner ID"})
		return
	}

	var partner models.Partner
	if err := c.ShouldBindJSON(&partner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	if err := h.partnerService.UpdatePartner(uint(id), &partner); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to update partner"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": partner})
}

// DeletePartner deletes a partner
func (h *PartnerHandler) DeletePartner(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid partner ID"})
		return
	}

	if err := h.partnerService.DeletePartner(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to delete partner"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Partner deleted successfully"})
}

// GetPartnersByCategory retrieves partners by category
func (h *PartnerHandler) GetPartnersByCategory(c *gin.Context) {
	category := c.Param("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Category is required"})
		return
	}

	partners, err := h.partnerService.GetPartnersByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to retrieve partners"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": partners})
}

// GetActivePartners retrieves all active partners
func (h *PartnerHandler) GetActivePartners(c *gin.Context) {
	partners, err := h.partnerService.GetActivePartners()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to retrieve partners"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": partners})
}
