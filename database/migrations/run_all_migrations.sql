-- Complete migration script for file upload support
-- Run this script to update the database for careers and commodities file support

USE gcx_market_data;

-- 1. Add file fields to careers table
ALTER TABLE careers 
ADD COLUMN file_path VARCHAR(500) DEFAULT NULL,
ADD COLUMN file_name VARCHAR(255) DEFAULT NULL;

-- 2. Verify commodities table has image_path field (should already exist)
-- If it doesn't exist, uncomment the line below:
-- ALTER TABLE commodities ADD COLUMN image_path VARCHAR(500) DEFAULT NULL;

-- 3. Show updated table structures
SELECT 'Careers table structure:' as info;
DESCRIBE careers;

SELECT 'Commodities table structure:' as info;
DESCRIBE commodities;

-- 4. Show sample data to verify
SELECT 'Sample careers data:' as info;
SELECT id, title, file_path, file_name FROM careers LIMIT 3;

SELECT 'Sample commodities data:' as info;
SELECT id, name, image_path FROM commodities LIMIT 3;
