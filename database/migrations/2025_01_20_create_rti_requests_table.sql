-- Migration: Create RTI Requests Table
-- Description: Creates table for Right to Information requests
-- Date: 2025-01-20

CREATE TABLE IF NOT EXISTS rti_requests (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    request_id VARCHAR(50) NOT NULL UNIQUE COMMENT 'Auto-generated request ID like RTI-2025-001',
    
    -- Requester Information
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    address TEXT,
    organization VARCHAR(255),
    
    -- Request Details
    request_type VARCHAR(100) NOT NULL COMMENT 'Type of information requested',
    subject VARCHAR(500) NOT NULL,
    description LONGTEXT NOT NULL,
    preferred_format VARCHAR(50) DEFAULT 'Electronic' COMMENT 'Electronic, Hard Copy, Both',
    
    -- Status and Processing
    status VARCHAR(50) DEFAULT 'pending' COMMENT 'pending, under_review, approved, rejected, completed',
    priority VARCHAR(50) DEFAULT 'normal' COMMENT 'low, normal, high, urgent',
    assigned_to VARCHAR(255) COMMENT 'RTI officer assigned',
    
    -- Response Information
    response_text LONGTEXT COMMENT 'Response from GCX',
    response_file VARCHAR(500) COMMENT 'URL to response document',
    response_date TIMESTAMP NULL,
    responded_by VARCHAR(255) COMMENT 'Officer who responded',
    
    -- Tracking
    reviewed_at TIMESTAMP NULL,
    reviewed_by VARCHAR(255),
    completed_at TIMESTAMP NULL,
    
    -- Metadata
    notes LONGTEXT COMMENT 'Internal notes',
    rejection_reason TEXT COMMENT 'Reason if rejected',
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    INDEX idx_request_id (request_id),
    INDEX idx_email (email),
    INDEX idx_status (status),
    INDEX idx_priority (priority),
    INDEX idx_created_at (created_at),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Sample data
INSERT INTO rti_requests (
    request_id, full_name, email, phone, address, organization,
    request_type, subject, description, preferred_format,
    status, priority, created_at
) VALUES
(
    'RTI-2025-001',
    'John Mensah',
    'john.mensah@example.com',
    '+233 24 123 4567',
    'Accra, Ghana',
    'Agricultural Research Institute',
    'Market Data',
    'Request for Annual Trading Volume Data',
    'I am requesting access to the annual trading volume data for all commodities traded on GCX for the years 2020-2024. This information is needed for academic research purposes.',
    'Electronic',
    'under_review',
    'normal',
    DATE_SUB(NOW(), INTERVAL 5 DAY)
),
(
    'RTI-2025-002',
    'Akosua Boateng',
    'akosua.boateng@example.com',
    '+233 54 987 6543',
    'Kumasi, Ghana',
    'Ghana Farmers Association',
    'Policy Documents',
    'Request for Trading Rules and Regulations',
    'I would like to obtain copies of the current trading rules and regulations, including any amendments made in the last two years.',
    'Both',
    'completed',
    'normal',
    DATE_SUB(NOW(), INTERVAL 15 DAY)
),
(
    'RTI-2025-003',
    'Kwame Asare',
    'kwame.asare@example.com',
    '+233 20 555 7890',
    'Takoradi, Ghana',
    NULL,
    'Financial Information',
    'Request for Membership Fee Structure',
    'I am interested in becoming a member and would like information about the current membership fee structure and associated costs.',
    'Electronic',
    'pending',
    'high',
    DATE_SUB(NOW(), INTERVAL 2 DAY)
);
