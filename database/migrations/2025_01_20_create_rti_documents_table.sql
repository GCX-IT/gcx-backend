-- Migration: Create RTI Documents Table
-- Description: Creates table for RTI downloadable resources (manuals, forms, guides)
-- Date: 2025-01-20

CREATE TABLE IF NOT EXISTS rti_documents (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    category VARCHAR(100) NOT NULL COMMENT 'manual, form, guide, policy, other',
    file_path VARCHAR(500) NOT NULL COMMENT 'URL to the PDF/document file',
    file_name VARCHAR(255),
    file_size BIGINT COMMENT 'File size in bytes',
    icon VARCHAR(50) DEFAULT 'pi-file-pdf' COMMENT 'PrimeIcons class name',
    download_count INT DEFAULT 0,
    is_featured BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_category (category),
    INDEX idx_is_active (is_active),
    INDEX idx_is_featured (is_featured),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert sample RTI documents
INSERT INTO rti_documents (
    title, description, category, file_path, file_name, icon, is_featured, is_active, sort_order
) VALUES
(
    'GCX RTI Manual',
    'Comprehensive guide to understanding the Right to Information process at GCX',
    'manual',
    '/documents/gcx-rti-manual.pdf',
    'GCX-RTI-Manual.pdf',
    'pi-book',
    TRUE,
    TRUE,
    1
),
(
    'RTI Application Form',
    'Download the official application form to submit your information request',
    'form',
    '/documents/rti-application-form.pdf',
    'RTI-Application-Form.pdf',
    'pi-file-edit',
    TRUE,
    TRUE,
    2
),
(
    'RTI Guidelines',
    'Guidelines for submitting and processing RTI requests',
    'guide',
    '/documents/rti-guidelines.pdf',
    'RTI-Guidelines.pdf',
    'pi-info-circle',
    TRUE,
    TRUE,
    3
),
(
    'RTI Act 2019',
    'Full text of the Right to Information Act, 2019',
    'policy',
    '/documents/rti-act-2019.pdf',
    'RTI-Act-2019.pdf',
    'pi-file',
    FALSE,
    TRUE,
    4
),
(
    'Exemptions Guide',
    'Information about exemptions under the RTI Act',
    'guide',
    '/documents/rti-exemptions.pdf',
    'RTI-Exemptions-Guide.pdf',
    'pi-shield',
    FALSE,
    TRUE,
    5
);
