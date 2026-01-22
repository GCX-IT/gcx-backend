-- Create commodity_contract_types table
-- Note: Price data (current_price, price_change, trading_volume) comes from Firebase API
-- CMS only handles static content (descriptions, images, specifications)
CREATE TABLE IF NOT EXISTS commodity_contract_types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    commodity_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) NOT NULL,
    description TEXT,
    full_description TEXT,
    specifications TEXT,
    trading_hours VARCHAR(255),
    contract_size VARCHAR(100),
    price_unit VARCHAR(50),
    image_path VARCHAR(500),
    contract_file VARCHAR(500),
    delivery_months VARCHAR(255),
    storage_requirements TEXT,
    quality_standards TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (commodity_id) REFERENCES commodities(id) ON DELETE CASCADE,
    INDEX idx_commodity_id (commodity_id),
    INDEX idx_code (code),
    INDEX idx_is_active (is_active)
);

-- Note: Contract types should be added through the CMS interface
-- This migration only creates the table structure
-- Content managers will add contract types through the admin panel
