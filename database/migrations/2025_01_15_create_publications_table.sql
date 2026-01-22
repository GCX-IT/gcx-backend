-- Create publications table
CREATE TABLE IF NOT EXISTS publications (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    category ENUM('Research Papers', 'Annual Reports', 'Policy Documents') NOT NULL,
    file_path VARCHAR(500),
    file_name VARCHAR(255),
    file_size INT,
    file_type VARCHAR(100),
    publication_date DATE,
    author VARCHAR(255),
    tags TEXT,
    status ENUM('Published', 'Draft', 'Archived') DEFAULT 'Published',
    download_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT INTO publications (title, description, category, file_path, file_name, file_size, file_type, publication_date, author, tags, status) VALUES
('GCX Market Analysis 2024', 'Comprehensive analysis of commodity markets and trading patterns for 2024', 'Research Papers', '/publications/research/gcx-market-analysis-2024.pdf', 'gcx-market-analysis-2024.pdf', 2048576, 'application/pdf', '2024-12-01', 'GCX Research Team', 'market analysis, commodities, trading', 'Published'),
('Annual Report 2023', 'GCX Annual Performance Report for the year 2023', 'Annual Reports', '/publications/annual/gcx-annual-report-2023.pdf', 'gcx-annual-report-2023.pdf', 5120000, 'application/pdf', '2024-03-15', 'GCX Management', 'annual report, performance, financial', 'Published'),
('Trading Rules and Regulations', 'Updated trading rules and regulations for GCX members', 'Policy Documents', '/publications/policy/trading-rules-2024.pdf', 'trading-rules-2024.pdf', 1536000, 'application/pdf', '2024-01-01', 'GCX Legal Team', 'trading rules, regulations, policy', 'Published'),
('Maize Market Outlook 2024', 'Research paper on maize market trends and future outlook', 'Research Papers', '/publications/research/maize-outlook-2024.pdf', 'maize-outlook-2024.pdf', 1024000, 'application/pdf', '2024-06-15', 'Dr. Sarah Mensah', 'maize, market outlook, research', 'Published'),
('GCX Compliance Guidelines', 'Guidelines for member compliance and regulatory requirements', 'Policy Documents', '/publications/policy/compliance-guidelines.pdf', 'compliance-guidelines.pdf', 768000, 'application/pdf', '2024-02-01', 'GCX Compliance Team', 'compliance, guidelines, regulatory', 'Published');
