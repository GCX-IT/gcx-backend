package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/database"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

// GetEvents retrieves all events (public endpoint)
func GetEvents(c *gin.Context) {
	db := database.GetDB()

	// Parse query parameters
	status := c.Query("status")     // upcoming, completed, cancelled
	eventType := c.Query("type")    // Conference, Summit, etc.
	category := c.Query("category") // Conference, Summit, etc.
	featured := c.Query("featured") // true/false
	search := c.Query("search")     // search term
	year := c.Query("year")         // filter by year
	limit := c.DefaultQuery("limit", "50")
	offset := c.DefaultQuery("offset", "0")

	var events []models.Event
	query := db.Where("is_active = ?", true)

	// Apply filters
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if eventType != "" {
		query = query.Where("type = ?", eventType)
	}

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if featured == "true" {
		query = query.Where("is_featured = ?", true)
	}

	if search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("title LIKE ? OR description LIKE ? OR location LIKE ?", searchTerm, searchTerm, searchTerm)
	}

	if year != "" {
		query = query.Where("YEAR(date) = ?", year)
	}

	// Order by date (upcoming first, then past events in reverse chronological order)
	query = query.Order("CASE WHEN date >= CURDATE() THEN 0 ELSE 1 END, date DESC")

	// Apply pagination
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)
	query = query.Limit(limitInt).Offset(offsetInt)

	if err := query.Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}

// GetEvent retrieves a single event by ID (public endpoint)
func GetEvent(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var event models.Event
	if err := db.Where("id = ? AND is_active = ?", id, true).First(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": event})
}

// GetEventBySlug retrieves a single event by slug (public endpoint)
func GetEventBySlug(c *gin.Context) {
	db := database.GetDB()
	eventSlug := c.Param("slug")

	var event models.Event
	if err := db.Where("slug = ? AND is_active = ?", eventSlug, true).First(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": event})
}

// GetUpcomingEvents retrieves all upcoming events (public endpoint)
func GetUpcomingEvents(c *gin.Context) {
	db := database.GetDB()
	limit := c.DefaultQuery("limit", "10")

	var events []models.Event
	limitInt, _ := strconv.Atoi(limit)

	if err := db.Where("is_active = ? AND status = ? AND date >= ?", true, models.EventStatusUpcoming, time.Now()).
		Order("date ASC").
		Limit(limitInt).
		Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve upcoming events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}

// GetPastEvents retrieves all past/completed events (public endpoint)
func GetPastEvents(c *gin.Context) {
	db := database.GetDB()
	limit := c.DefaultQuery("limit", "10")

	var events []models.Event
	limitInt, _ := strconv.Atoi(limit)

	if err := db.Where("is_active = ? AND (status = ? OR date < ?)", true, models.EventStatusCompleted, time.Now()).
		Order("date DESC").
		Limit(limitInt).
		Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve past events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}

// GetAllEvents retrieves all events for CMS (protected endpoint)
func GetAllEvents(c *gin.Context) {
	db := database.GetDB()

	var events []models.Event
	if err := db.Order("date DESC").Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}

// CreateEvent creates a new event (protected endpoint)
func CreateEvent(c *gin.Context) {
	db := database.GetDB()

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate slug from title if not provided
	if event.Slug == "" {
		event.Slug = slug.Make(event.Title)
	} else {
		event.Slug = slug.Make(event.Slug)
	}

	// Check if slug already exists
	var existingEvent models.Event
	if err := db.Where("slug = ?", event.Slug).First(&existingEvent).Error; err == nil {
		// Slug exists, append timestamp
		event.Slug = event.Slug + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	}

	// Create event
	if err := db.Create(&event).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": "Event with this slug already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"data":    event,
	})
}

// UpdateEvent updates an existing event (protected endpoint)
func UpdateEvent(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var event models.Event
	if err := db.First(&event, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}

	var updates models.Event
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update slug if title changed
	if updates.Title != "" && updates.Title != event.Title {
		updates.Slug = slug.Make(updates.Title)

		// Check if new slug already exists (excluding current event)
		var existingEvent models.Event
		if err := db.Where("slug = ? AND id != ?", updates.Slug, id).First(&existingEvent).Error; err == nil {
			// Slug exists, append timestamp
			updates.Slug = updates.Slug + "-" + strconv.FormatInt(time.Now().Unix(), 10)
		}
	}

	if err := db.Model(&event).Updates(&updates).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": "Event with this slug already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"data":    event,
	})
}

// DeleteEvent deletes an event (soft delete) (protected endpoint)
func DeleteEvent(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	var event models.Event
	if err := db.First(&event, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}

	if err := db.Delete(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

// GetEventStats retrieves statistics about events (protected endpoint)
func GetEventStats(c *gin.Context) {
	db := database.GetDB()

	var totalEvents int64
	var upcomingEvents int64
	var completedEvents int64
	var totalAttendees int64

	db.Model(&models.Event{}).Where("is_active = ?", true).Count(&totalEvents)
	db.Model(&models.Event{}).Where("is_active = ? AND status = ?", true, models.EventStatusUpcoming).Count(&upcomingEvents)
	db.Model(&models.Event{}).Where("is_active = ? AND status = ?", true, models.EventStatusCompleted).Count(&completedEvents)

	var attendeesSum struct {
		Total int64
	}
	db.Model(&models.Event{}).Where("is_active = ?", true).Select("COALESCE(SUM(attendees), 0) as total").Scan(&attendeesSum)
	totalAttendees = attendeesSum.Total

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"total_events":     totalEvents,
			"upcoming_events":  upcomingEvents,
			"completed_events": completedEvents,
			"total_attendees":  totalAttendees,
		},
	})
}

// RegisterForEvent handles event registration (public endpoint)
func RegisterForEvent(c *gin.Context) {
	db := database.GetDB()
	eventID := c.Param("id")

	// Check if event exists and registration is open
	var event models.Event
	if err := db.First(&event, eventID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}

	if !event.CanRegister() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Registration is not open for this event"})
		return
	}

	var registration models.EventRegistration
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registration.EventID = event.ID
	registration.RegistrationStatus = "pending"

	// Check for duplicate registration
	var existingRegistration models.EventRegistration
	if err := db.Where("event_id = ? AND email = ?", event.ID, registration.Email).First(&existingRegistration).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "You have already registered for this event"})
		return
	}

	if err := db.Create(&registration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register for event"})
		return
	}

	// Send confirmation emails
	go func() {
		// This runs in a goroutine to avoid blocking the response
		// The email service will be called from the frontend
		// TODO: Implement backend email service if needed
	}()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
		"data":    registration,
	})
}

// GetEventRegistrations retrieves all registrations for an event (protected endpoint)
func GetEventRegistrations(c *gin.Context) {
	db := database.GetDB()
	eventID := c.Param("id")

	var registrations []models.EventRegistration
	if err := db.Where("event_id = ?", eventID).Preload("Event").Find(&registrations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve registrations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": registrations})
}
