-- Add image_path field to publications table for publication covers/thumbnails
ALTER TABLE publications 
ADD COLUMN image_path VARCHAR(500) AFTER category;

-- Update existing publications with placeholder image paths (optional)
-- UPDATE publications SET image_path = '/publications/covers/default-cover.jpg' WHERE image_path = '' OR image_path IS NULL;
