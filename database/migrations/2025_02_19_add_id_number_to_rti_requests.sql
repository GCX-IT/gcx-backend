-- Migration: Add ID Number (Ghana Card) field to RTI Requests Table
-- Date: 2025-02-19
-- Description: Adds id_number field to store Ghana Card ID numbers for RTI applicants

-- Add id_number column if it doesn't exist (required)
-- Check if column exists before adding
SET @dbname = DATABASE();
SET @tablename = 'rti_requests';
SET @columnname = 'id_number';
SET @preparedStatement = (SELECT IF(
  (
    SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
    WHERE
      (TABLE_SCHEMA = @dbname)
      AND (TABLE_NAME = @tablename)
      AND (COLUMN_NAME = @columnname)
  ) > 0,
  'SELECT 1', -- Column exists, do nothing
  CONCAT('ALTER TABLE `', @tablename, '` ADD COLUMN `', @columnname, '` VARCHAR(50) NOT NULL DEFAULT '''' AFTER `phone`')
));
PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;

-- Update existing records to have empty string for id_number if NULL
-- Using id (primary key) in WHERE clause to satisfy safe update mode
UPDATE `rti_requests` SET `id_number` = '' WHERE `id` > 0 AND `id_number` IS NULL;

-- Make address required (modify existing column, don't add if already exists)
-- Since address column already exists, we just modify it to be NOT NULL
-- Using id (primary key) in WHERE clause to satisfy safe update mode
UPDATE `rti_requests` SET `address` = '' WHERE `id` > 0 AND `address` IS NULL;
ALTER TABLE `rti_requests` 
MODIFY COLUMN `address` VARCHAR(255) NOT NULL DEFAULT '';

-- Make preferred_format required
ALTER TABLE `rti_requests` 
MODIFY COLUMN `preferred_format` VARCHAR(50) NOT NULL DEFAULT 'electronic';
