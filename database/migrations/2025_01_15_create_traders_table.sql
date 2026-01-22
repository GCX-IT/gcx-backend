-- Create traders table for membership list
CREATE TABLE IF NOT EXISTS traders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    industry VARCHAR(255) DEFAULT '-',
    member_type ENUM('Associates', 'Full Members', 'Brokers', 'Warehouse Operators') NOT NULL DEFAULT 'Associates',
    phone_no VARCHAR(50),
    email VARCHAR(255),
    address TEXT,
    registration_date DATE,
    status ENUM('Active', 'Inactive', 'Suspended') DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT INTO traders (name, industry, member_type, phone_no, email, address, registration_date, status) VALUES
('Eric Kwabena Agyei', '-', 'Associates', '0244 692 089/ 0245 071 061', 'eric.agyei@example.com', 'Accra, Ghana', '2024-01-15', 'Active'),
('Joseph Awudu Malik', '-', 'Associates', '0244 869 387', 'joseph.malik@example.com', 'Kumasi, Ghana', '2024-01-20', 'Active'),
('Kofi Adusei Koduah', '-', 'Associates', '0540 122295', 'kofi.koduah@example.com', 'Tamale, Ghana', '2024-02-01', 'Active'),
('Kwabena Duah Agyemang', '-', 'Associates', '0201 689497', 'kwabena.agyemang@example.com', 'Cape Coast, Ghana', '2024-02-10', 'Active'),
('Maame Adjoa Thompson', '-', 'Associates', '0302 200748/ 0266 802388', 'maame.thompson@example.com', 'Tema, Ghana', '2024-02-15', 'Active'),
('Monica . .', '-', 'Associates', '0269 382146', 'monica@example.com', 'Takoradi, Ghana', '2024-02-20', 'Active'),
('Praise Awisi Bogobley', '-', 'Associates', '0244 590 621', 'praise.bogobley@example.com', 'Ho, Ghana', '2024-03-01', 'Active'),
('Roland Apindem Ajiabadek', '-', 'Associates', '020 535 1223', 'roland.ajiabadek@example.com', 'Koforidua, Ghana', '2024-03-05', 'Active');
