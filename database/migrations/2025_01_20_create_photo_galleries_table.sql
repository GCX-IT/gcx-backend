-- Migration: Create Photo Galleries Tables
-- Description: Creates tables for photo gallery albums and individual photos
-- Date: 2025-01-20

-- Create photo galleries table (albums/folders)
CREATE TABLE IF NOT EXISTS photo_galleries (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    category VARCHAR(100) NOT NULL COMMENT 'Events, Training, Operations, Quality, Programs, Analysis, Partnerships, Forums, Technology',
    cover_image VARCHAR(500) COMMENT 'URL to cover/thumbnail image',
    date DATE COMMENT 'Date of event/program',
    location VARCHAR(255),
    tags JSON COMMENT 'Array of tags',
    photo_count INT DEFAULT 0,
    is_featured BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_slug (slug),
    INDEX idx_category (category),
    INDEX idx_is_active (is_active),
    INDEX idx_is_featured (is_featured),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create gallery photos table (individual photos within albums)
CREATE TABLE IF NOT EXISTS gallery_photos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    gallery_id BIGINT UNSIGNED NOT NULL,
    title VARCHAR(255),
    description TEXT,
    image_url VARCHAR(500) NOT NULL,
    thumbnail_url VARCHAR(500),
    photographer VARCHAR(255),
    caption TEXT,
    tags JSON,
    sort_order INT DEFAULT 0,
    is_cover BOOLEAN DEFAULT FALSE COMMENT 'Is this the cover photo for the gallery',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (gallery_id) REFERENCES photo_galleries(id) ON DELETE CASCADE,
    INDEX idx_gallery_id (gallery_id),
    INDEX idx_is_cover (is_cover),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert sample galleries (photo_count will auto-increment when photos are added)
INSERT INTO photo_galleries (title, slug, description, category, cover_image, date, location, tags, photo_count, is_featured, is_active) VALUES
('Events & Conferences', 'events-conferences', 'Key moments from our annual conferences, events, and stakeholder meetings.', 'Events', '/Picture3.png', '2023-12-15', 'Accra, Ghana', '["Conference", "Leadership"]', 0, TRUE, TRUE),
('Training & Workshops', 'training-workshops', 'Training sessions and workshops empowering farmers and traders.', 'Training', '/Picture3.png', '2023-11-20', 'Kumasi, Ghana', '["Training", "Farmers"]', 0, TRUE, TRUE),
('Trading Operations', 'trading-operations', 'Behind-the-scenes look at our trading operations and market activities.', 'Operations', '/trading.jpg', '2023-10-15', 'GCX Trading Floor', '["Trading", "Operations"]', 0, TRUE, TRUE),
('Quality Assurance', 'quality-assurance', 'Our quality assurance team ensuring the highest standards for commodity trading.', 'Quality', '/Picture3.png', '2023-09-10', 'Warehouse Facilities', '["Quality", "Inspection"]', 0, TRUE, TRUE),
('Youth Programs', 'youth-programs', 'Engaging young people in agriculture and commodity trading through education and mentorship.', 'Programs', '/Picture3.png', '2023-08-25', 'Various Locations', '["Youth", "Programs"]', 0, TRUE, TRUE),
('Market Analysis', 'market-analysis', 'Expert analysis sessions on commodity market trends and trading strategies.', 'Analysis', '/trading dashboard.jpg', '2023-07-30', 'GCX Headquarters', '["Analysis", "Market"]', 0, TRUE, TRUE),
('International Partnerships', 'international-partnerships', 'Building strategic partnerships with international commodity exchanges and organizations.', 'Partnerships', '/Picture3.png', '2023-06-15', 'International', '["Partnerships", "International"]', 0, TRUE, TRUE),
('Women in Agriculture', 'women-agriculture', 'Empowering women farmers and traders in the commodity exchange ecosystem.', 'Forums', '/Picture3.png', '2023-05-20', 'Cape Coast, Ghana', '["Women", "Forum"]', 0, TRUE, TRUE),
('Technology & Innovation', 'technology-innovation', 'Showcasing cutting-edge technology solutions for modern agriculture.', 'Technology', '/trading.jpg', '2023-04-10', 'Tech Hub, Accra', '["Technology", "Innovation"]', 0, TRUE, TRUE);

-- Insert sample photos for first gallery (using actual available images)
INSERT INTO gallery_photos (gallery_id, title, description, image_url, thumbnail_url, photographer, sort_order, is_cover) VALUES
(1, 'Opening Ceremony', 'Annual conference opening ceremony', '/trading.jpg', '/trading.jpg', 'GCX Media Team', 1, TRUE),
(1, 'Keynote Address', 'CEO delivering keynote address', '/farmer.jpg', '/farmer.jpg', 'GCX Media Team', 2, FALSE),
(1, 'Panel Discussion', 'Expert panel discussing market trends', '/maize.jpg', '/maize.jpg', 'GCX Media Team', 3, FALSE),
(1, 'Networking Session', 'Attendees networking during break', '/crop.jpg', '/crop.jpg', 'GCX Media Team', 4, FALSE);
