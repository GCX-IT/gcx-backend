package handlers

import (
	"net/http"
	"strconv"

	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateRTIRequest creates a new RTI request (public endpoint)
func CreateRTIRequest(c *gin.Context) {
	db := database.GetDB()

	var request models.RTIRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create request (request_id will be auto-generated)
	if err := db.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create RTI request"})
		return
	}

	// TODO: Send confirmation email to requester

	c.JSON(http.StatusCreated, gin.H{
		"message": "RTI request submitted successfully",
		"data":    request,
	})
}

// GetAllRTIRequests retrieves all RTI requests for CMS (protected endpoint)
func GetAllRTIRequests(c *gin.Context) {
	db := database.GetDB()

	// Parse query parameters
	status := c.Query("status")
	priority := c.Query("priority")
	search := c.Query("search")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	var requests []models.RTIRequest
	query := db.Model(&models.RTIRequest{})

	// Apply filters
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if priority != "" {
		query = query.Where("priority = ?", priority)
	}

	if search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("full_name LIKE ? OR email LIKE ? OR request_id LIKE ? OR subject LIKE ?",
			searchTerm, searchTerm, searchTerm, searchTerm)
	}

	// Count total
	var total int64
	query.Count(&total)

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&requests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve RTI requests"})
		return
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	c.JSON(http.StatusOK, gin.H{
		"data": requests,
		"pagination": gin.H{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}

// GetRTIRequest retrieves a single RTI request (protected endpoint)
func GetRTIRequest(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var request models.RTIRequest
	if err := db.First(&request, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "RTI request not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve RTI request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": request})
}

// UpdateRTIRequest updates an RTI request (protected endpoint)
func UpdateRTIRequest(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var request models.RTIRequest
	if err := db.First(&request, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "RTI request not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve RTI request"})
		return
	}

	var updates models.RTIRequest
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&request).Updates(&updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update RTI request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "RTI request updated successfully",
		"data":    request,
	})
}

// RespondToRTIRequest adds a response to an RTI request (protected endpoint)
func RespondToRTIRequest(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var request models.RTIRequest
	if err := db.First(&request, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "RTI request not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve RTI request"})
		return
	}

	var response struct {
		ResponseText string  `json:"response_text" binding:"required"`
		ResponseFile *string `json:"response_file"`
		RespondedBy  string  `json:"responded_by" binding:"required"`
	}

	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.Complete(response.ResponseText, response.ResponseFile, response.RespondedBy)

	if err := db.Save(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to respond to RTI request"})
		return
	}

	// TODO: Send email notification to requester

	c.JSON(http.StatusOK, gin.H{
		"message": "Response submitted successfully",
		"data":    request,
	})
}

// UpdateRTIStatus updates the status of an RTI request (protected endpoint)
func UpdateRTIStatus(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var request models.RTIRequest
	if err := db.First(&request, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "RTI request not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve RTI request"})
		return
	}

	var statusUpdate struct {
		Status          string  `json:"status" binding:"required"`
		ReviewerName    *string `json:"reviewer_name"`
		RejectionReason *string `json:"rejection_reason"`
	}

	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.Status = models.RTIStatus(statusUpdate.Status)

	if statusUpdate.Status == string(models.RTIStatusUnderReview) && statusUpdate.ReviewerName != nil {
		request.MarkAsReviewed(*statusUpdate.ReviewerName)
	}

	if statusUpdate.Status == string(models.RTIStatusRejected) && statusUpdate.RejectionReason != nil {
		request.Reject(*statusUpdate.RejectionReason)
	}

	if err := db.Save(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update RTI request status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Status updated successfully",
		"data":    request,
	})
}

// DeleteRTIRequest deletes an RTI request (soft delete) (protected endpoint)
func DeleteRTIRequest(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var request models.RTIRequest
	if err := db.First(&request, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "RTI request not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve RTI request"})
		return
	}

	if err := db.Delete(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete RTI request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "RTI request deleted successfully"})
}

// GetRTIStats retrieves statistics about RTI requests (protected endpoint)
func GetRTIStats(c *gin.Context) {
	db := database.GetDB()

	var totalRequests int64
	var pendingRequests int64
	var underReviewRequests int64
	var completedRequests int64

	db.Model(&models.RTIRequest{}).Count(&totalRequests)
	db.Model(&models.RTIRequest{}).Where("status = ?", models.RTIStatusPending).Count(&pendingRequests)
	db.Model(&models.RTIRequest{}).Where("status = ?", models.RTIStatusUnderReview).Count(&underReviewRequests)
	db.Model(&models.RTIRequest{}).Where("status = ?", models.RTIStatusCompleted).Count(&completedRequests)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"total_requests":        totalRequests,
			"pending_requests":      pendingRequests,
			"under_review_requests": underReviewRequests,
			"completed_requests":    completedRequests,
		},
	})
}
