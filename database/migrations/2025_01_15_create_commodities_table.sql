-- Create commodities table
CREATE TABLE IF NOT EXISTS commodities (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    specifications TEXT,
    trading_hours VARCHAR(255),
    contract_size VARCHAR(100),
    price_unit VARCHAR(50),
    minimum_price DECIMAL(10,2),
    maximum_price DECIMAL(10,2),
    current_price DECIMAL(10,2),
    price_change DECIMAL(10,2),
    price_change_percent DECIMAL(5,2),
    trading_volume BIGINT DEFAULT 0,
    market_status ENUM('Open', 'Closed', 'Suspended') DEFAULT 'Open',
    image_path VARCHAR(500),
    category VARCHAR(100),
    origin_country VARCHAR(100),
    harvest_season VARCHAR(100),
    storage_requirements TEXT,
    quality_standards TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT INTO commodities (name, code, description, specifications, trading_hours, contract_size, price_unit, minimum_price, maximum_price, current_price, price_change, price_change_percent, trading_volume, market_status, image_path, category, origin_country, harvest_season, storage_requirements, quality_standards) VALUES
('Maize', 'GAPWM2', 'White maize trading information and specifications', 'Moisture content: 12-14%, Foreign matter: 2% max, Broken kernels: 5% max', 'Monday-Friday 9:00-17:00', '50kg bags', 'GHC per bag', 1800.00, 2000.00, 1880.00, 30.00, 1.62, 150000, 'Open', '/commodities/maize.jpg', 'Grains', 'Ghana', 'June-September', 'Cool, dry storage, 12-14% moisture', 'Ghana Standards Authority Grade A'),
('Soya Bean', 'GEJWM2', 'Soya bean market data and trading details', 'Protein content: 35% min, Oil content: 18% min, Moisture: 12% max', 'Monday-Friday 9:00-17:00', '50kg bags', 'GHC per bag', 3800.00, 4200.00, 4030.00, 50.00, 1.26, 75000, 'Open', '/commodities/soya-bean.jpg', 'Oilseeds', 'Ghana', 'October-December', 'Dry storage, 12% moisture max', 'Ghana Standards Authority Grade A'),
('Sorghum', 'GSAWM2', 'Sorghum commodity information and trading specs', 'Moisture: 12% max, Foreign matter: 2% max, Broken grains: 5% max', 'Monday-Friday 9:00-17:00', '50kg bags', 'GHC per bag', 4600.00, 4900.00, 4745.00, 25.00, 0.53, 45000, 'Open', '/commodities/sorghum.jpg', 'Grains', 'Ghana', 'September-December', 'Dry storage, 12% moisture', 'Ghana Standards Authority Grade A'),
('Sesame', 'GKUWM2', 'Sesame trading specifications and market data', 'Oil content: 45% min, Moisture: 8% max, Foreign matter: 1% max', 'Monday-Friday 9:00-17:00', '50kg bags', 'GHC per bag', 4500.00, 4800.00, 4645.00, 45.00, 0.98, 25000, 'Open', '/commodities/sesame.jpg', 'Oilseeds', 'Ghana', 'November-January', 'Cool, dry storage, 8% moisture', 'Ghana Standards Authority Grade A'),
('Rice', 'GRIWM2', 'Rice commodity market data and trading information', 'Moisture: 14% max, Broken grains: 5% max, Foreign matter: 1% max', 'Monday-Friday 9:00-17:00', '50kg bags', 'GHC per bag', 2200.00, 2500.00, 2350.00, -20.00, -0.84, 85000, 'Open', '/commodities/rice.jpg', 'Grains', 'Ghana', 'October-December', 'Dry storage, 14% moisture max', 'Ghana Standards Authority Grade A');
