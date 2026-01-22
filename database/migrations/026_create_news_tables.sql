-- Create news_categories table
CREATE TABLE IF NOT EXISTS news_categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    color VARCHAR(7), -- For hex color codes
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_slug (slug),
    INDEX idx_active (is_active)
);

-- Create news_items table
CREATE TABLE IF NOT EXISTS news_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    content TEXT,
    source ENUM('gcx', 'partner', 'external', 'api', 'firebase') NOT NULL DEFAULT 'gcx',
    source_name VARCHAR(255), -- For external sources
    source_url VARCHAR(500), -- Link to original source
    category VARCHAR(100), -- e.g., 'market', 'announcement', 'event'
    priority INT DEFAULT 0, -- Higher numbers = higher priority
    status ENUM('draft', 'published', 'archived') DEFAULT 'draft',
    is_breaking BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    published_at TIMESTAMP NULL,
    expires_at TIMESTAMP NULL, -- Auto-hide after this time
    external_id VARCHAR(255), -- ID from external system
    external_data JSON, -- JSON data from external source
    last_sync_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    INDEX idx_status (status),
    INDEX idx_source (source),
    INDEX idx_category (category),
    INDEX idx_priority (priority),
    INDEX idx_breaking (is_breaking),
    INDEX idx_active (is_active),
    INDEX idx_published (published_at),
    INDEX idx_expires (expires_at),
    INDEX idx_external_id (external_id),
    INDEX idx_deleted_at (deleted_at),
    INDEX idx_status_active_published (status, is_active, published_at),
    INDEX idx_breaking_priority (is_breaking, priority)
);

-- Create news_source_configs table
CREATE TABLE IF NOT EXISTS news_source_configs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    source_name VARCHAR(255) NOT NULL UNIQUE,
    source_type ENUM('gcx', 'partner', 'external', 'api', 'firebase') NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    api_endpoint VARCHAR(500),
    api_key VARCHAR(255),
    refresh_interval INT DEFAULT 300, -- seconds
    last_sync_at TIMESTAMP NULL,
    config JSON, -- Additional config as JSON
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_source_name (source_name),
    INDEX idx_source_type (source_type),
    INDEX idx_active (is_active)
);

-- Insert default news categories
INSERT INTO news_categories (name, slug, description, color) VALUES
('Market Updates', 'market-updates', 'Latest market news and price updates', '#3B82F6'),
('Announcements', 'announcements', 'Important announcements from GCX', '#10B981'),
('Events', 'events', 'Upcoming events and programs', '#F59E0B'),
('Partnerships', 'partnerships', 'Partnership and collaboration news', '#8B5CF6'),
('Regulations', 'regulations', 'Regulatory updates and compliance', '#EF4444'),
('Technology', 'technology', 'Technology and innovation news', '#06B6D4')
ON DUPLICATE KEY UPDATE name = VALUES(name);

-- Insert default news source configs
INSERT INTO news_source_configs (source_name, source_type, is_active, refresh_interval) VALUES
('GCX Internal', 'gcx', TRUE, 0),
('Firebase Public', 'firebase', TRUE, 300),
('External API', 'api', FALSE, 600)
ON DUPLICATE KEY UPDATE source_type = VALUES(source_type);
