package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type EventStatus string
type EventType string

const (
	EventStatusUpcoming  EventStatus = "upcoming"
	EventStatusCompleted EventStatus = "completed"
	EventStatusCancelled EventStatus = "cancelled"
)

const (
	EventTypeConference EventType = "Conference"
	EventTypeSummit     EventType = "Summit"
	EventTypeWorkshop   EventType = "Workshop"
	EventTypeForum      EventType = "Forum"
	EventTypeProgram    EventType = "Program"
	EventTypeMeeting    EventType = "Meeting"
)

// Event represents an event or achievement
type Event struct {
	ID                   uint           `json:"id" gorm:"primaryKey"`
	Title                string         `json:"title" gorm:"not null"`
	Slug                 string         `json:"slug" gorm:"uniqueIndex;not null"`
	Date                 time.Time      `json:"date" gorm:"not null"`
	Time                 *string        `json:"time"`
	Location             string         `json:"location" gorm:"not null"`
	Venue                *string        `json:"venue"`
	Address              *string        `json:"address"`
	Type                 EventType      `json:"type" gorm:"not null"`
	Category             string         `json:"category" gorm:"not null"`
	Status               EventStatus    `json:"status" gorm:"default:upcoming"`
	Description          *string        `json:"description"`
	FullDescription      *string        `json:"full_description" gorm:"type:longtext"`
	Attendees            int            `json:"attendees" gorm:"default:0"`
	Image                *string        `json:"image"`
	RegistrationOpen     bool           `json:"registration_open" gorm:"default:false"`
	RegistrationDeadline *time.Time     `json:"registration_deadline"`
	Price                *string        `json:"price"`
	Speakers             datatypes.JSON `json:"speakers" gorm:"type:json"`
	Agenda               datatypes.JSON `json:"agenda" gorm:"type:json"`
	Requirements         datatypes.JSON `json:"requirements" gorm:"type:json"`
	MetaTitle            *string        `json:"meta_title"`
	MetaDescription      *string        `json:"meta_description"`
	SortOrder            int            `json:"sort_order" gorm:"default:0"`
	IsActive             bool           `json:"is_active" gorm:"default:true"`
	IsFeatured           bool           `json:"is_featured" gorm:"default:false"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `json:"-" gorm:"index"`
}

// Speaker represents an event speaker
type Speaker struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Image string `json:"image"`
}

// AgendaItem represents an agenda item
type AgendaItem struct {
	Time    string `json:"time"`
	Session string `json:"session"`
}

// IsUpcoming checks if the event is upcoming
func (e *Event) IsUpcoming() bool {
	return e.Status == EventStatusUpcoming && e.Date.After(time.Now())
}

// IsCompleted checks if the event is completed
func (e *Event) IsCompleted() bool {
	return e.Status == EventStatusCompleted || (e.Date.Before(time.Now()) && e.Status != EventStatusCancelled)
}

// CanRegister checks if registration is open for the event
func (e *Event) CanRegister() bool {
	if !e.RegistrationOpen || !e.IsUpcoming() {
		return false
	}
	if e.RegistrationDeadline != nil {
		return e.RegistrationDeadline.After(time.Now())
	}
	return true
}

// MarkAsCompleted marks the event as completed
func (e *Event) MarkAsCompleted() {
	e.Status = EventStatusCompleted
}

// Cancel cancels the event
func (e *Event) Cancel() {
	e.Status = EventStatusCancelled
	e.RegistrationOpen = false
}

// BeforeCreate hook to ensure defaults are set
func (e *Event) BeforeCreate(tx *gorm.DB) error {
	// Initialize JSON fields if nil
	if e.Speakers == nil {
		e.Speakers = datatypes.JSON([]byte("[]"))
	}
	if e.Agenda == nil {
		e.Agenda = datatypes.JSON([]byte("[]"))
	}
	if e.Requirements == nil {
		e.Requirements = datatypes.JSON([]byte("[]"))
	}
	return nil
}

// TableName returns the table name for Event model
func (Event) TableName() string {
	return "events"
}

// EventRegistration represents an event registration
type EventRegistration struct {
	ID                  uint           `json:"id" gorm:"primaryKey"`
	EventID             uint           `json:"event_id" gorm:"not null"`
	FullName            string         `json:"full_name" gorm:"not null"`
	Email               string         `json:"email" gorm:"not null"`
	Phone               string         `json:"phone" gorm:"not null"`
	Organization        *string        `json:"organization"`
	Position            *string        `json:"position"`
	DietaryRequirements *string        `json:"dietary_requirements"`
	SpecialNeeds        *string        `json:"special_needs"`
	RegistrationStatus  string         `json:"registration_status" gorm:"default:pending"` // pending, confirmed, cancelled
	ConfirmationSent    bool           `json:"confirmation_sent" gorm:"default:false"`
	ConfirmationSentAt  *time.Time     `json:"confirmation_sent_at"`
	CheckedIn           bool           `json:"checked_in" gorm:"default:false"`
	CheckedInAt         *time.Time     `json:"checked_in_at"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Event Event `json:"event" gorm:"foreignKey:EventID"`
}

// TableName returns the table name for EventRegistration model
func (EventRegistration) TableName() string {
	return "event_registrations"
}
