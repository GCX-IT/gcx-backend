# Events & Achievements CMS Setup Guide

This document describes the Events and Achievements management system that has been implemented for the GCX platform.

## Overview

The Events & Achievements system allows administrators to manage events (conferences, summits, workshops, forums, programs, meetings) through a CMS interface, with public-facing pages that display event information and allow user registration.

## What Was Implemented

### Backend (Go)

#### 1. Database Migration
- **File**: `database/migrations/2025_01_20_create_events_table.sql`
- Creates the `events` table with all necessary fields
- Includes sample data for 6 events (3 upcoming, 3 past)
- Creates indexes for optimal query performance

#### 2. Event Model
- **File**: `internal/cms/models/event.go`
- Defines the Event struct with all fields
- Includes EventRegistration model for handling registrations
- Helper methods: `IsUpcoming()`, `IsCompleted()`, `CanRegister()`, etc.

#### 3. Event Handlers
- **File**: `internal/cms/handlers/event.go`
- **Public endpoints**:
  - `GET /api/events` - List all active events with filtering
  - `GET /api/events/upcoming` - Get upcoming events
  - `GET /api/events/past` - Get past/completed events
  - `GET /api/events/:id` - Get single event by ID
  - `GET /api/events/slug/:slug` - Get event by slug
  - `POST /api/events/:id/register` - Register for an event

- **Protected endpoints (CMS)**:
  - `GET /api/cms/events` - List all events for CMS
  - `POST /api/cms/events` - Create new event
  - `PUT /api/cms/events/:id` - Update event
  - `DELETE /api/cms/events/:id` - Delete event (soft delete)
  - `GET /api/cms/events/stats` - Get event statistics
  - `GET /api/cms/events/:id/registrations` - View registrations

#### 4. Routes
- **File**: `routes/cms.go`
- All event routes registered with proper middleware
- Public routes accessible without authentication
- Protected routes require authentication

### Frontend (Vue.js)

#### 1. Event Service
- **File**: `src/services/eventService.ts`
- TypeScript service for all event-related API calls
- Handles both public and authenticated requests
- Type definitions for Event, Speaker, AgendaItem, etc.

#### 2. Public Views

##### ArchivesView.vue
- **Route**: `/media/archives` (this is the achievements page)
- Displays past events with filtering
- Search, category, and year filters
- Fetches data from API
- Responsive grid layout

##### EventArchivesView.vue
- **Route**: `/media/archives`
- Shows both upcoming and past events
- Advanced filtering (status, type, category, year)
- Fetches data from API
- Click to view event details

##### EventDetailView.vue
- **Route**: `/media/archives/:slug`
- Full event details page
- Event registration form
- Speaker information
- Event agenda
- Requirements checklist
- Fetches event by slug from API

#### 3. CMS Components

##### EventManager.vue
- **Component**: `src/components/CMS/EventManager.vue`
- Full CRUD interface for managing events
- Filtering by status, type, and search
- Add/Edit modal with all event fields
- View registrations functionality
- Color-coded status and type badges

#### 4. Router & Navigation
- **Files Updated**:
  - `src/router/index.ts` - Added `/cms/events` route
  - `src/views/CMSView.vue` - Added EventManager component
  - `src/components/CMS/CMSSidebar.vue` - Added Events menu item

## Event Data Structure

```typescript
{
  id: number
  title: string
  slug: string
  date: string (ISO date)
  time?: string
  location: string
  venue?: string
  address?: string
  type: 'Conference' | 'Summit' | 'Workshop' | 'Forum' | 'Program' | 'Meeting'
  category: string
  status: 'upcoming' | 'completed' | 'cancelled'
  description?: string
  full_description?: string
  attendees: number
  image?: string
  registration_open: boolean
  registration_deadline?: string
  price?: string
  speakers?: Speaker[]
  agenda?: AgendaItem[]
  requirements?: string[]
  is_active: boolean
  is_featured: boolean
  sort_order: number
}
```

## Installation & Setup

### 1. Run Database Migration

**Option A: Using the batch script**
```bash
cd gcx-go-backend
migrate-events.bat
```

**Option B: Manual MySQL command**
```bash
mysql -u root -p gcx_cms < database/migrations/2025_01_20_create_events_table.sql
```

### 2. Start Backend Server

```bash
cd gcx-go-backend
go run main.go
```

The server will start on `http://localhost:8080`

### 3. Start Frontend Server

```bash
cd gcx-frontend1
npm install  # if not already done
npm run dev
```

The frontend will start on `http://localhost:5173`

## Usage

### Public Access

1. **View All Events**: Navigate to `/media/archives`
2. **Filter Events**: Use the search and filter options
3. **View Event Details**: Click on any event card
4. **Register for Event**: Fill out the registration form on event detail page

### CMS Access

1. **Login**: Go to `/login` and authenticate
2. **Access CMS**: Navigate to `/cms/events`
3. **Add Event**: Click "Add Event" button
4. **Edit Event**: Click pencil icon on any event
5. **Delete Event**: Click trash icon (soft delete)
6. **View Registrations**: Click users icon to see registrations

## Features

### Event Management
- ✅ Create, Read, Update, Delete events
- ✅ Slug generation from title
- ✅ Status management (upcoming, completed, cancelled)
- ✅ Type categorization
- ✅ Featured events
- ✅ Registration management
- ✅ Soft delete with recovery option

### Filtering & Search
- ✅ Search by title, location, description
- ✅ Filter by status
- ✅ Filter by type/category
- ✅ Filter by year
- ✅ Combined filters

### Event Registration
- ✅ Public registration form
- ✅ Registration deadline validation
- ✅ Duplicate registration prevention
- ✅ Registration tracking
- ✅ Email confirmation (TODO: email service)

### UI/UX
- ✅ Responsive design
- ✅ Dark mode support
- ✅ Loading states
- ✅ Error handling
- ✅ Color-coded badges
- ✅ Image support
- ✅ Modern card layout

## API Endpoints Reference

### Public Endpoints

```
GET    /api/events                     - List events (with filters)
GET    /api/events/upcoming            - Upcoming events
GET    /api/events/past                - Past events
GET    /api/events/:id                 - Get event by ID
GET    /api/events/slug/:slug          - Get event by slug
POST   /api/events/:id/register        - Register for event
```

### CMS Endpoints (Authenticated)

```
GET    /api/cms/events                 - List all events
POST   /api/cms/events                 - Create event
PUT    /api/cms/events/:id             - Update event
DELETE /api/cms/events/:id             - Delete event
GET    /api/cms/events/stats           - Event statistics
GET    /api/cms/events/:id/registrations - View registrations
```

## Filter Parameters

The `/api/events` endpoint supports the following query parameters:

- `status` - Filter by status (upcoming, completed, cancelled)
- `type` - Filter by type (Conference, Summit, Workshop, etc.)
- `category` - Filter by category
- `featured` - Filter featured events (true/false)
- `search` - Search in title, location, description
- `year` - Filter by year
- `limit` - Number of results (default: 50)
- `offset` - Pagination offset (default: 0)

Example: `/api/events?status=upcoming&type=Conference&year=2025&limit=10`

## Sample Data

The migration includes 6 sample events:

1. **GCX Annual Conference 2025** (Upcoming, Featured)
2. **Agricultural Technology Summit 2025** (Upcoming, Featured)
3. **Youth in Agriculture Program 2025** (Upcoming)
4. **GCX Annual Conference 2023** (Completed)
5. **Agricultural Innovation Summit 2023** (Completed)
6. **Commodity Market Analysis Workshop 2023** (Completed)

## Future Enhancements

### Potential additions:
- [ ] Email notifications for registrations
- [ ] Calendar export (iCal format)
- [ ] Event image upload functionality
- [ ] Advanced speaker management
- [ ] Attendee check-in system
- [ ] Event feedback/ratings
- [ ] Multi-language support for events
- [ ] Event templates
- [ ] Recurring events
- [ ] Event capacity management

## Security Notes

- All CMS endpoints require authentication
- Soft delete preserves data integrity
- Input validation on all forms
- SQL injection prevention via parameterized queries
- XSS prevention via proper escaping
- CSRF protection (ensure tokens are implemented)

## Troubleshooting

### Events not showing up
1. Check if events are marked as `is_active = true`
2. Verify date formats are correct
3. Check browser console for API errors

### Can't register for event
1. Verify `registration_open = true`
2. Check registration deadline hasn't passed
3. Ensure event status is 'upcoming'
4. Check for duplicate email registration

### CMS not accessible
1. Verify user is authenticated
2. Check user role has proper permissions
3. Ensure backend server is running
4. Check CORS settings if frontend/backend on different ports

## Support

For issues or questions, please contact the development team or create an issue in the project repository.

---

**Version**: 1.0.0  
**Last Updated**: January 20, 2025  
**Author**: GCX Development Team
