-- Fix: Update gallery photo URLs to use actual available images
-- Run this if you've already run the main migration and photos show as black

-- Update the sample photos to use actual image files
UPDATE gallery_photos 
SET 
    image_url = CASE 
        WHEN id = 1 THEN '/trading.jpg'
        WHEN id = 2 THEN '/farmer.jpg'
        WHEN id = 3 THEN '/maize.jpg'
        WHEN id = 4 THEN '/crop.jpg'
        ELSE image_url
    END,
    thumbnail_url = CASE 
        WHEN id = 1 THEN '/trading.jpg'
        WHEN id = 2 THEN '/farmer.jpg'
        WHEN id = 3 THEN '/maize.jpg'
        WHEN id = 4 THEN '/crop.jpg'
        ELSE thumbnail_url
    END
WHERE gallery_id = 1;

-- Show the updated URLs
SELECT id, title, image_url, thumbnail_url 
FROM gallery_photos 
WHERE gallery_id = 1 
ORDER BY sort_order;
