-- Fix: Update photo_count to match actual number of photos in each gallery
-- Run this if you've already run the main migration

-- Update photo counts based on actual photos
UPDATE photo_galleries pg
SET photo_count = (
    SELECT COUNT(*)
    FROM gallery_photos gp
    WHERE gp.gallery_id = pg.id
    AND gp.deleted_at IS NULL
)
WHERE pg.deleted_at IS NULL;

-- Show the updated counts
SELECT 
    id,
    title,
    photo_count,
    (SELECT COUNT(*) FROM gallery_photos WHERE gallery_id = photo_galleries.id AND deleted_at IS NULL) as actual_count
FROM photo_galleries
WHERE deleted_at IS NULL
ORDER BY id;
