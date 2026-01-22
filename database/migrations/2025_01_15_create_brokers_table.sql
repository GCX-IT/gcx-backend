-- Create brokers table
CREATE TABLE IF NOT EXISTS brokers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    company VARCHAR(255),
    license_number VARCHAR(100),
    phone_no VARCHAR(50),
    email VARCHAR(255),
    address TEXT,
    specialization VARCHAR(255),
    experience_years INT DEFAULT 0,
    status ENUM('Active', 'Inactive', 'Suspended') DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT INTO brokers (name, company, license_number, phone_no, email, address, specialization, experience_years, status) VALUES
('John Mensah', 'AgriTrade Ltd', 'BRK001', '0244 123 456', 'john.mensah@agritrade.com', 'Accra, Ghana', 'Agricultural Commodities', 5, 'Active'),
('Sarah Osei', 'Commodity Brokers Ghana', 'BRK002', '0244 234 567', 'sarah.osei@cbg.com', 'Kumasi, Ghana', 'Grain Trading', 8, 'Active'),
('Michael Asante', 'West Africa Trading Co', 'BRK003', '0244 345 678', 'michael.asante@watc.com', 'Tema, Ghana', 'Export/Import', 12, 'Active'),
('Grace Adjei', 'FarmLink Brokers', 'BRK004', '0244 456 789', 'grace.adjei@farmlink.com', 'Tamale, Ghana', 'Livestock Trading', 6, 'Active'),
('David Nkrumah', 'Ghana Commodity Brokers', 'BRK005', '0244 567 890', 'david.nkrumah@gcb.com', 'Cape Coast, Ghana', 'Oil Seeds', 10, 'Active');
