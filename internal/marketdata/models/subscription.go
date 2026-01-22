package models

import (
	shared_models "gcx-cms/internal/shared/models"
	"time"
)

// SubscriptionPlan represents available subscription plans
type SubscriptionPlan struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"` // Basic, Premium, Enterprise
	Description string    `json:"description" gorm:"type:text"`
	Price       float64   `json:"price" gorm:"not null"` // Monthly price in GHS
	Currency    string    `json:"currency" gorm:"default:GHS"`
	Duration    int       `json:"duration" gorm:"not null"`   // Duration in days
	Features    string    `json:"features" gorm:"type:text"`  // Comma-separated features
	MaxUsers    int       `json:"max_users" gorm:"default:1"` // Maximum users per subscription
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	SortOrder   int       `json:"sort_order" gorm:"default:0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserSubscription represents user's active subscription
type UserSubscription struct {
	ID               uint       `json:"id" gorm:"primaryKey"`
	UserID           uint       `json:"user_id" gorm:"not null"`
	PlanID           uint       `json:"plan_id" gorm:"not null"`
	Status           string     `json:"status" gorm:"default:active"` // active, expired, cancelled, suspended
	StartDate        time.Time  `json:"start_date" gorm:"not null"`
	EndDate          time.Time  `json:"end_date" gorm:"not null"`
	AutoRenew        bool       `json:"auto_renew" gorm:"default:true"`
	PaymentMethod    string     `json:"payment_method"` // stripe, mobile_money, bank_transfer
	PaymentReference string     `json:"payment_reference"`
	AmountPaid       float64    `json:"amount_paid"`
	Currency         string     `json:"currency" gorm:"default:GHS"`
	LastBillingDate  *time.Time `json:"last_billing_date"`
	NextBillingDate  *time.Time `json:"next_billing_date"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`

	// Relationships
	User shared_models.User `json:"user" gorm:"foreignKey:UserID"`
	Plan SubscriptionPlan   `json:"plan" gorm:"foreignKey:PlanID"`
}

// SubscriptionFeature represents features available in each plan
type SubscriptionFeature struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	PlanID      uint      `json:"plan_id" gorm:"not null"`
	Name        string    `json:"name" gorm:"not null"` // real_time_data, historical_data, alerts, etc.
	Description string    `json:"description" gorm:"type:text"`
	IsEnabled   bool      `json:"is_enabled" gorm:"default:true"`
	Limit       *int      `json:"limit"` // Usage limit if applicable
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships
	Plan SubscriptionPlan `json:"plan" gorm:"foreignKey:PlanID"`
}

// UserDataAccess represents user's access to different data types
type UserDataAccess struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	DataType     string    `json:"data_type" gorm:"not null"`      // real_time, historical, analytics, alerts
	AccessLevel  string    `json:"access_level" gorm:"not null"`   // full, limited, none
	Commodities  string    `json:"commodities" gorm:"type:text"`   // Comma-separated commodity codes
	MaxRequests  int       `json:"max_requests"`                   // Daily request limit
	RequestCount int       `json:"request_count" gorm:"default:0"` // Current day's request count
	LastReset    time.Time `json:"last_reset"`                     // Last time request count was reset
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Relationships
	User shared_models.User `json:"user" gorm:"foreignKey:UserID"`
}

// TableName returns the table name for SubscriptionPlan model
func (SubscriptionPlan) TableName() string {
	return "subscription_plans"
}

// TableName returns the table name for UserSubscription model
func (UserSubscription) TableName() string {
	return "user_subscriptions"
}

// TableName returns the table name for SubscriptionFeature model
func (SubscriptionFeature) TableName() string {
	return "subscription_features"
}

// TableName returns the table name for UserDataAccess model
func (UserDataAccess) TableName() string {
	return "user_data_access"
}

// IsActive checks if subscription is currently active
func (us *UserSubscription) IsActive() bool {
	now := time.Now()
	return us.Status == "active" && now.After(us.StartDate) && now.Before(us.EndDate)
}

// DaysUntilExpiry returns days until subscription expires
func (us *UserSubscription) DaysUntilExpiry() int {
	now := time.Now()
	if now.After(us.EndDate) {
		return 0
	}
	duration := us.EndDate.Sub(now)
	return int(duration.Hours() / 24)
}
