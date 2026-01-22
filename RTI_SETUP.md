# RTI (Right to Information) System Setup Guide

This document describes the RTI Request Management system that has been implemented for the GCX platform.

## Overview

The RTI system allows citizens to submit Right to Information requests through the public website, and administrators can manage, respond to, and track these requests through the CMS.

## What Was Implemented

### Backend (Go)

#### 1. Database Migration
- **File**: `database/migrations/2025_01_20_create_rti_requests_table.sql`
- Creates the `rti_requests` table with all necessary fields
- Includes 3 sample RTI requests (pending, under review, completed)
- Auto-generates unique Request IDs (RTI-2025-001, RTI-2025-002, etc.)

#### 2. RTI Request Model
- **File**: `internal/cms/models/rti.go`
- Complete data structure for RTI requests
- Status management: pending → under_review → approved → completed/rejected
- Priority levels: low, normal, high, urgent
- Helper methods for status transitions

#### 3. RTI Handlers
- **File**: `internal/cms/handlers/rti.go`
- **Public endpoint**:
  - `POST /api/rti/request` - Submit RTI request

- **Protected endpoints (CMS)**:
  - `GET /api/cms/rti/requests` - List all requests (with filters)
  - `GET /api/cms/rti/requests/:id` - Get single request
  - `PUT /api/cms/rti/requests/:id` - Update request
  - `POST /api/cms/rti/requests/:id/respond` - Respond to request
  - `PUT /api/cms/rti/requests/:id/status` - Update status
  - `DELETE /api/cms/rti/requests/:id` - Delete request
  - `GET /api/cms/rti/stats` - Get statistics

#### 4. Routes
- **File**: `routes/cms.go`
- All RTI routes registered with proper middleware

### Frontend (Vue.js)

#### 1. RTI Service
- **File**: `src/services/rtiService.ts`
- TypeScript service for all RTI-related API calls
- Type definitions for RTIRequest

#### 2. Public RTI Page
- **File**: `src/views/RTIView.vue`
- RTI application form
- Submits to backend API
- Shows generated Request ID
- Email confirmation (TODO: implement email service)

#### 3. CMS RTI Manager
- **File**: `src/components/CMS/RTIManager.vue`
- Complete request management interface
- Features:
  - Statistics dashboard
  - Advanced filtering (status, priority, search)
  - Pagination
  - View full request details
  - Respond with text and file upload
  - Status management
  - Delete requests
  - Export functionality

#### 4. Router & Navigation
- **Files Updated**:
  - `src/router/index.ts` - Added `/cms/rti` route
  - `src/views/CMSView.vue` - Added RTIManager component
  - `src/components/CMS/CMSSidebar.vue` - Added "RTI Requests" menu item

## RTI Request Data Structure

```typescript
{
  id: number
  request_id: string          // Auto-generated: RTI-2025-001
  full_name: string
  email: string
  phone: string
  address?: string
  organization?: string
  request_type: string        // Type of information requested
  subject: string
  description: string
  preferred_format: string    // Electronic, Hard Copy, Both
  status: 'pending' | 'under_review' | 'approved' | 'rejected' | 'completed'
  priority: 'low' | 'normal' | 'high' | 'urgent'
  assigned_to?: string
  response_text?: string
  response_file?: string
  response_date?: string
  responded_by?: string
  notes?: string
  rejection_reason?: string
  created_at: string
  updated_at: string
}
```

## Installation & Setup

### 1. Run Database Migration

```bash
cd gcx-go-backend
migrate-rti.bat
```

Or manually:
```bash
mysql -u root -p gcx_cms < database/migrations/2025_01_20_create_rti_requests_table.sql
```

### 2. Start Backend Server

```bash
cd gcx-go-backend
go run main.go
```

### 3. Start Frontend Server

```bash
cd gcx-frontend1
npm run dev
```

## Usage

### Public Access (Citizens)

1. **Submit RTI Request**: Navigate to `/rti`
2. **Fill out the form**:
   - Full Name *
   - Email *
   - Phone *
   - Address
   - Organization
   - Type of Information *
   - Details of Information Required *
   - Purpose of Request *
   - Preferred Format
3. **Submit** → Receive auto-generated Request ID
4. **Keep Request ID** for future reference

### CMS Access (Administrators)

1. **Login**: Go to `/login`
2. **Access RTI Manager**: `/cms/rti` or click "RTI Requests" in sidebar
3. **View Dashboard**:
   - Total Requests
   - Pending
   - Under Review
   - Completed

4. **Manage Requests**:
   - 👁️ View - See full request details
   - 💬 Respond - Add response with optional file
   - 🗑️ Delete - Remove request

5. **Change Status**:
   - Pending → Under Review
   - Under Review → Approved
   - Any → Rejected (with reason)
   - Approved → Completed (after response)

6. **Filter & Search**:
   - Search by name, email, Request ID
   - Filter by status
   - Filter by priority
   - Pagination for large datasets

## Features

### Request Management
- ✅ Auto-generated unique Request IDs
- ✅ Full CRUD operations
- ✅ Status workflow management
- ✅ Priority assignment
- ✅ Response system (text + file)
- ✅ Search and filtering
- ✅ Pagination
- ✅ Statistics dashboard

### Status Workflow
```
New Request → Pending
     ↓
Under Review (assigned to officer)
     ↓
Approved → Response Added → Completed
     ↓
  OR Rejected (with reason)
```

### Response System
- ✅ Text response
- ✅ File upload (documents, PDFs)
- ✅ Response tracking (who responded, when)
- ✅ Download response files

### Tracking
- ✅ Request creation date
- ✅ Review date and reviewer
- ✅ Response date and responder
- ✅ Completion date
- ✅ Internal notes

## API Endpoints

### Public Endpoints

```
POST /api/rti/request - Submit RTI request
```

### CMS Endpoints (Authenticated)

```
GET    /api/cms/rti/requests              - List all requests
GET    /api/cms/rti/requests/:id          - Get single request
PUT    /api/cms/rti/requests/:id          - Update request
POST   /api/cms/rti/requests/:id/respond  - Add response
PUT    /api/cms/rti/requests/:id/status   - Update status
DELETE /api/cms/rti/requests/:id          - Delete request
GET    /api/cms/rti/stats                 - Get statistics
```

## Sample Workflow

### Scenario: Citizen Requests Market Data

1. **Citizen** visits `/rti` and submits request:
   - Name: "John Mensah"
   - Email: "john@example.com"
   - Type: "Market Data"
   - Subject: "Annual Trading Volume Data"

2. **System** generates Request ID: `RTI-2025-004`

3. **Admin** sees request in CMS:
   - Status: 🟡 Pending
   - Priority: Normal

4. **Admin** clicks "Mark as Under Review":
   - Status: 🔵 Under Review

5. **Admin** clicks "Approve":
   - Status: 🟢 Approved

6. **Admin** clicks "Respond":
   - Enters response text
   - Uploads supporting PDF document
   - Submits

7. **Status** changes to: 🟣 Completed
8. **Citizen** receives email with response (TODO: email integration)

## Statistics Dashboard

The RTI Manager shows real-time statistics:
- **Total Requests** - All requests submitted
- **Pending** - Awaiting review (🟡)
- **Under Review** - Being processed (🔵)
- **Completed** - Responded (🟣)

## Security

- ✅ Public can only submit requests
- ✅ All management endpoints require authentication
- ✅ Soft delete preserves data
- ✅ Input validation on all forms
- ✅ SQL injection prevention

## Future Enhancements

- [ ] Email notifications (confirmation, status updates, response)
- [ ] SMS notifications
- [ ] Request tracking by Request ID (public page)
- [ ] Automated response templates
- [ ] SLA tracking (response time monitoring)
- [ ] Export to Excel/PDF
- [ ] Advanced reporting and analytics
- [ ] Bulk operations
- [ ] Email integration for responses

## Troubleshooting

### Requests not showing
1. Check if `rti_requests` table exists
2. Verify backend is running
3. Check authentication token
4. Check browser console for errors

### Can't submit request
1. Verify all required fields are filled
2. Check backend logs for errors
3. Verify API endpoint is accessible

### Can't respond to request
1. Ensure request status allows responses
2. Check file upload if document is attached
3. Verify "responded_by" field is filled

---

**Version**: 1.0.0  
**Last Updated**: January 20, 2025  
**Author**: GCX Development Team
