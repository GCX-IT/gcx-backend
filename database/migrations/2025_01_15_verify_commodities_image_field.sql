-- Verify commodities table has image_path field
-- This migration ensures the image_path field exists for commodity images

USE gcx_market_data;

-- Check if image_path field exists, if not add it
SET @sql = (SELECT IF(
    (SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS 
     WHERE TABLE_SCHEMA = 'gcx_market_data' 
     AND TABLE_NAME = 'commodities' 
     AND COLUMN_NAME = 'image_path') > 0,
    'SELECT "image_path field already exists" as status;',
    'ALTER TABLE commodities ADD COLUMN image_path VARCHAR(500) DEFAULT NULL;'
));
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- Verify the commodities table structure
DESCRIBE commodities;
