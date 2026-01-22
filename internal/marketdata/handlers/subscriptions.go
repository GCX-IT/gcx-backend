package handlers

import (
	"net/http"
	"strconv"
	"time"

	"gcx-cms/internal/marketdata/models"
	"gcx-cms/internal/shared/config"

	"github.com/gin-gonic/gin"
)

// GetSubscriptionPlans returns all available subscription plans
func GetSubscriptionPlans(c *gin.Context) {
	var plans []models.SubscriptionPlan

	if err := config.DB.Where("is_active = ?", true).
		Order("sort_order ASC, price ASC").
		Find(&plans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch subscription plans",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    plans,
		"count":   len(plans),
	})
}

// GetUserSubscription returns the current user's subscription
func GetUserSubscription(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	var subscription models.UserSubscription

	if err := config.DB.Preload("Plan").
		Where("user_id = ? AND status = ?", userID, "active").
		First(&subscription).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No active subscription found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    subscription,
	})
}

// CreateSubscription creates a new subscription for the user
func CreateSubscription(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	var req struct {
		PlanID           uint   `json:"plan_id" binding:"required"`
		PaymentMethod    string `json:"payment_method" binding:"required"`
		PaymentReference string `json:"payment_reference"`
		AutoRenew        bool   `json:"auto_renew"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	// Check if plan exists and is active
	var plan models.SubscriptionPlan
	if err := config.DB.Where("id = ? AND is_active = ?", req.PlanID, true).
		First(&plan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid subscription plan",
		})
		return
	}

	// Check if user already has an active subscription
	var existingSubscription models.UserSubscription
	if err := config.DB.Where("user_id = ? AND status = ?", userID, "active").
		First(&existingSubscription).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already has an active subscription",
		})
		return
	}

	// Create new subscription
	subscription := models.UserSubscription{
		UserID:           userID.(uint),
		PlanID:           req.PlanID,
		Status:           "active",
		StartDate:        time.Now(),
		EndDate:          time.Now().AddDate(0, 0, plan.Duration),
		AutoRenew:        req.AutoRenew,
		PaymentMethod:    req.PaymentMethod,
		PaymentReference: req.PaymentReference,
		AmountPaid:       plan.Price,
		Currency:         plan.Currency,
		LastBillingDate:  &time.Time{},
		NextBillingDate:  &time.Time{},
	}

	if err := config.DB.Create(&subscription).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create subscription",
			"details": err.Error(),
		})
		return
	}

	// Load the plan details for response
	if err := config.DB.Preload("Plan").First(&subscription, subscription.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to load subscription details",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Subscription created successfully",
		"data":    subscription,
	})
}

// UpdateSubscription updates an existing subscription
func UpdateSubscription(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	subscriptionID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid subscription ID",
		})
		return
	}

	var req struct {
		AutoRenew     *bool   `json:"auto_renew"`
		PaymentMethod *string `json:"payment_method"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	// Get user's subscription
	var subscription models.UserSubscription
	if err := config.DB.Where("id = ? AND user_id = ?", subscriptionID, userID).
		First(&subscription).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Subscription not found",
		})
		return
	}

	// Update fields if provided
	updates := make(map[string]interface{})
	if req.AutoRenew != nil {
		updates["auto_renew"] = *req.AutoRenew
	}
	if req.PaymentMethod != nil {
		updates["payment_method"] = *req.PaymentMethod
	}

	if len(updates) > 0 {
		if err := config.DB.Model(&subscription).Updates(updates).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to update subscription",
				"details": err.Error(),
			})
			return
		}
	}

	// Load updated subscription
	if err := config.DB.Preload("Plan").First(&subscription, subscription.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to load updated subscription",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Subscription updated successfully",
		"data":    subscription,
	})
}

// CancelSubscription cancels a user's subscription
func CancelSubscription(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	subscriptionID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid subscription ID",
		})
		return
	}

	// Get user's subscription
	var subscription models.UserSubscription
	if err := config.DB.Where("id = ? AND user_id = ?", subscriptionID, userID).
		First(&subscription).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Subscription not found",
		})
		return
	}

	// Cancel subscription
	if err := config.DB.Model(&subscription).Updates(map[string]interface{}{
		"status":     "cancelled",
		"auto_renew": false,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to cancel subscription",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Subscription cancelled successfully",
	})
}
