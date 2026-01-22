package models

import (
	"time"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin   UserRole = "admin"
	RoleBlogger UserRole = "blogger"
	RoleUser    UserRole = "user"
	RoleTrader  UserRole = "trader"    // New role for market data access
	RolePremium UserRole = "premium"   // New role for premium market data
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"` // Hidden from JSON
	Role      UserRole  `json:"role" gorm:"default:user"`
	Avatar    *string   `json:"avatar"`
	Bio       *string   `json:"bio"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	LastLogin *time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Market Data specific fields
	Company     *string `json:"company"`      // Company name for business users
	Phone       *string `json:"phone"`        // Phone number
	Country     *string `json:"country"`      // Country for regional access
	TimeZone    *string `json:"time_zone"`    // User's timezone
	Preferences string  `json:"preferences" gorm:"type:text"` // JSON string for user preferences
}

// HashPassword hashes the user's password
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword verifies the password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// CanManageBlog checks if user can manage blog posts
func (u *User) CanManageBlog() bool {
	return u.Role == RoleAdmin || u.Role == RoleBlogger
}

// CanManageUsers checks if user can manage other users
func (u *User) CanManageUsers() bool {
	return u.Role == RoleAdmin
}

// CanManageSettings checks if user can manage system settings
func (u *User) CanManageSettings() bool {
	return u.Role == RoleAdmin
}

// CanAccessMarketData checks if user can access market data
func (u *User) CanAccessMarketData() bool {
	return u.Role == RoleAdmin || u.Role == RoleTrader || u.Role == RolePremium || u.Role == RoleUser
}

// CanAccessRealTimeData checks if user can access real-time market data
func (u *User) CanAccessRealTimeData() bool {
	return u.Role == RoleAdmin || u.Role == RoleTrader || u.Role == RolePremium
}

// CanAccessHistoricalData checks if user can access historical market data
func (u *User) CanAccessHistoricalData() bool {
	return u.Role == RoleAdmin || u.Role == RoleTrader || u.Role == RolePremium
}

// CanAccessAnalytics checks if user can access market analytics
func (u *User) CanAccessAnalytics() bool {
	return u.Role == RoleAdmin || u.Role == RoleTrader || u.Role == RolePremium
}

// BeforeCreate hook to hash password before creating user
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Password != "" {
		return u.HashPassword()
	}
	return nil
}

// TableName returns the table name for User model
func (User) TableName() string {
	return "users"
}
