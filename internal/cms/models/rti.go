package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type RTIStatus string
type RTIPriority string

const (
	RTIStatusPending     RTIStatus = "pending"
	RTIStatusUnderReview RTIStatus = "under_review"
	RTIStatusApproved    RTIStatus = "approved"
	RTIStatusRejected    RTIStatus = "rejected"
	RTIStatusCompleted   RTIStatus = "completed"
)

const (
	RTIPriorityLow    RTIPriority = "low"
	RTIPriorityNormal RTIPriority = "normal"
	RTIPriorityHigh   RTIPriority = "high"
	RTIPriorityUrgent RTIPriority = "urgent"
)

// RTIRequest represents a Right to Information request
type RTIRequest struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	RequestID       string         `json:"request_id" gorm:"type:varchar(191);uniqueIndex;not null"`
	FullName        string         `json:"full_name" gorm:"not null"`
	Email           string         `json:"email" gorm:"not null"`
	Phone           string         `json:"phone" gorm:"not null"`
	Address         *string        `json:"address"`
	Organization    *string        `json:"organization"`
	RequestType     string         `json:"request_type" gorm:"not null"`
	Subject         string         `json:"subject" gorm:"not null"`
	Description     string         `json:"description" gorm:"type:longtext;not null"`
	PreferredFormat string         `json:"preferred_format" gorm:"default:'Electronic'"`
	Status          RTIStatus      `json:"status" gorm:"default:'pending'"`
	Priority        RTIPriority    `json:"priority" gorm:"default:'normal'"`
	AssignedTo      *string        `json:"assigned_to"`
	ResponseText    *string        `json:"response_text" gorm:"type:longtext"`
	ResponseFile    *string        `json:"response_file"`
	ResponseDate    *time.Time     `json:"response_date"`
	RespondedBy     *string        `json:"responded_by"`
	ReviewedAt      *time.Time     `json:"reviewed_at"`
	ReviewedBy      *string        `json:"reviewed_by"`
	CompletedAt     *time.Time     `json:"completed_at"`
	Notes           *string        `json:"notes" gorm:"type:longtext"`
	RejectionReason *string        `json:"rejection_reason"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}

// IsPending checks if request is pending
func (r *RTIRequest) IsPending() bool {
	return r.Status == RTIStatusPending
}

// IsCompleted checks if request is completed
func (r *RTIRequest) IsCompleted() bool {
	return r.Status == RTIStatusCompleted
}

// CanRespond checks if request can be responded to
func (r *RTIRequest) CanRespond() bool {
	return r.Status == RTIStatusPending || r.Status == RTIStatusUnderReview || r.Status == RTIStatusApproved
}

// MarkAsReviewed marks the request as reviewed
func (r *RTIRequest) MarkAsReviewed(reviewerName string) {
	r.Status = RTIStatusUnderReview
	now := time.Now()
	r.ReviewedAt = &now
	r.ReviewedBy = &reviewerName
}

// Approve approves the request
func (r *RTIRequest) Approve() {
	r.Status = RTIStatusApproved
}

// Reject rejects the request
func (r *RTIRequest) Reject(reason string) {
	r.Status = RTIStatusRejected
	r.RejectionReason = &reason
}

// Complete marks the request as completed
func (r *RTIRequest) Complete(responseText string, responseFile *string, respondedBy string) {
	r.Status = RTIStatusCompleted
	r.ResponseText = &responseText
	r.ResponseFile = responseFile
	now := time.Now()
	r.ResponseDate = &now
	r.CompletedAt = &now
	r.RespondedBy = &respondedBy
}

// GenerateRequestID generates a unique request ID
func GenerateRequestID(db *gorm.DB) (string, error) {
	year := time.Now().Year()

	// Count requests for this year
	var count int64
	db.Model(&RTIRequest{}).
		Where("request_id LIKE ?", fmt.Sprintf("RTI-%d-%%", year)).
		Count(&count)

	// Generate ID like RTI-2025-001, RTI-2025-002, etc.
	requestID := fmt.Sprintf("RTI-%d-%03d", year, count+1)

	// Check if it exists (edge case)
	var existing RTIRequest
	if err := db.Where("request_id = ?", requestID).First(&existing).Error; err == nil {
		// ID exists, try next number
		requestID = fmt.Sprintf("RTI-%d-%03d", year, count+2)
	}

	return requestID, nil
}

// BeforeCreate hook to generate request ID
func (r *RTIRequest) BeforeCreate(tx *gorm.DB) error {
	if r.RequestID == "" {
		requestID, err := GenerateRequestID(tx)
		if err != nil {
			return err
		}
		r.RequestID = requestID
	}
	return nil
}

// TableName returns the table name for RTIRequest model
func (RTIRequest) TableName() string {
	return "rti_requests"
}
