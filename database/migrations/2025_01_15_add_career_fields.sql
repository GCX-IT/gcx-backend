-- Add image_path and contact_email fields to careers table
ALTER TABLE careers 
ADD COLUMN image_path VARCHAR(500) AFTER application_count,
ADD COLUMN contact_email VARCHAR(255) AFTER image_path;
