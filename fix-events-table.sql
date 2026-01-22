-- Fix: Add deleted_at column to events table
-- This fixes the "Unknown column 'events.deleted_at'" error

ALTER TABLE events ADD COLUMN deleted_at TIMESTAMP NULL DEFAULT NULL AFTER updated_at;
ALTER TABLE events ADD INDEX idx_deleted_at (deleted_at);

-- Verify the column was added
DESCRIBE events;
