-- Migration: Create Video Library Tables
-- Description: Creates tables for video library categories and individual videos
-- Date: 2025-01-20

-- Create video libraries table (categories/folders)
CREATE TABLE IF NOT EXISTS video_libraries (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    category VARCHAR(100) NOT NULL COMMENT 'Tutorial, Analysis, Stories, Events, Education, Technology',
    cover_image VARCHAR(500) COMMENT 'URL to cover/thumbnail image',
    date DATE COMMENT 'Date of video/program',
    location VARCHAR(255),
    tags JSON COMMENT 'Array of tags',
    video_count INT DEFAULT 0,
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

-- Create library videos table (individual videos within categories)
CREATE TABLE IF NOT EXISTS library_videos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    library_id BIGINT UNSIGNED NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    video_url VARCHAR(500) NOT NULL COMMENT 'URL to video file (AWS S3 or external)',
    thumbnail_url VARCHAR(500) COMMENT 'URL to video thumbnail',
    duration VARCHAR(20) COMMENT 'Video duration (e.g., 5:30)',
    file_size BIGINT COMMENT 'Video file size in bytes',
    video_type VARCHAR(50) DEFAULT 'mp4' COMMENT 'mp4, webm, etc.',
    resolution VARCHAR(20) COMMENT '720p, 1080p, 4K, etc.',
    view_count INT DEFAULT 0,
    is_featured BOOLEAN DEFAULT FALSE,
    is_cover BOOLEAN DEFAULT FALSE COMMENT 'Is this the cover video for the library',
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (library_id) REFERENCES video_libraries(id) ON DELETE CASCADE,
    INDEX idx_library_id (library_id),
    INDEX idx_is_cover (is_cover),
    INDEX idx_view_count (view_count),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert sample video libraries (video_count will auto-increment when videos are added)
INSERT INTO video_libraries (title, slug, description, category, cover_image, date, location, tags, video_count, is_featured, is_active) VALUES
('Tutorial Videos', 'tutorial-videos', 'Step-by-step guides and tutorials for using GCX platform and services.', 'Tutorial', '/Picture3.png', '2024-01-15', 'GCX Headquarters', '["Tutorial", "Platform", "Guide"]', 0, TRUE, TRUE),
('Market Analysis', 'market-analysis', 'Comprehensive analysis of commodity market trends and performance.', 'Analysis', '/trading.jpg', '2024-01-10', 'GCX Trading Floor', '["Analysis", "Market", "Trends"]', 0, TRUE, TRUE),
('Success Stories', 'success-stories', 'Inspiring stories from farmers and traders who have benefited from GCX services.', 'Stories', '/farmer.jpg', '2024-01-05', 'Various Locations', '["Stories", "Success", "Farmers"]', 0, TRUE, TRUE),
('Events & Conferences', 'events-conferences', 'Highlights from GCX events, conferences, and important announcements.', 'Events', '/conference.jpg', '2023-12-20', 'Accra, Ghana', '["Events", "Conference", "Highlights"]', 0, TRUE, TRUE),
('Educational Content', 'educational-content', 'Educational videos about commodity trading, agriculture, and market knowledge.', 'Education', '/maize.jpg', '2023-12-15', 'Training Centers', '["Education", "Training", "Knowledge"]', 0, TRUE, TRUE),
('Technology & Innovation', 'technology-innovation', 'Latest technology solutions and innovations in agriculture and trading.', 'Technology', '/trading.jpg', '2023-12-10', 'Tech Hub', '["Technology", "Innovation", "Digital"]', 0, TRUE, TRUE);

-- Insert sample videos for first library
INSERT INTO library_videos (library_id, title, description, video_url, thumbnail_url, duration, file_size, video_type, resolution, view_count, is_featured, is_cover) VALUES
(1, 'GCX Trading Platform Overview', 'Learn about the Ghana Commodity Exchange trading platform and how to get started with commodity trading.', 'https://example.com/videos/gcx-platform-overview.mp4', '/Picture3.png', '5:30', 52428800, 'mp4', '1080p', 1300, TRUE, TRUE),
(1, 'How to Register as a Trader', 'Step-by-step guide on how to register and become a certified trader on the GCX platform.', 'https://example.com/videos/trader-registration.mp4', '/Picture3.png', '3:45', 31457280, 'mp4', '720p', 890, FALSE, FALSE),
(1, 'Understanding Trading Orders', 'Learn about different types of trading orders and how to place them effectively.', 'https://example.com/videos/trading-orders.mp4', '/Picture3.png', '4:20', 41943040, 'mp4', '1080p', 650, FALSE, FALSE);
