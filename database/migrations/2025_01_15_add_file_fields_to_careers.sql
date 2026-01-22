-- Add file fields to careers table
-- Run this migration to add file upload support to careers

USE gcx_market_data;

-- Add file fields to careers table
ALTER TABLE careers 
ADD COLUMN file_path VARCHAR(500) DEFAULT NULL,
ADD COLUMN file_name VARCHAR(255) DEFAULT NULL;

-- Verify the changes
DESCRIBE careers;