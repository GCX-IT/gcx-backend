-- Create team_members table
CREATE TABLE IF NOT EXISTS team_members (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    image VARCHAR(500),
    type ENUM('board', 'executive', 'functional') NOT NULL,
    order_index INT DEFAULT 0,
    -- Social Media Handles
    linkedin_url VARCHAR(500),
    twitter_url VARCHAR(500),
    facebook_url VARCHAR(500),
    instagram_url VARCHAR(500),
    email VARCHAR(255),
    phone VARCHAR(50),
    website_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_type (type),
    INDEX idx_order (type, order_index)
);

-- Insert sample data
INSERT INTO team_members (name, title, description, image, type, order_index) VALUES
-- Board Members
('Mr. Kofi S. Yamoah', 'Board Chairman', 'Experienced leader providing strategic direction and governance oversight.', '/Board of directors/Mr. Kofi S. Yamoah \'Board Chairman\'.png', 'board', 1),
('Mr. Kwame Daaku', 'Non-Executive Director', 'Independent director bringing valuable industry expertise and governance experience.', '/Board of directors/Mr. Kwame Daaku \'Non-excutive director\'.jpg', 'board', 2),
('Mr. Stephen Antwi-Asimeng', 'Non-Executive Director', 'Independent director with extensive experience in financial markets and corporate governance.', '/Board of directors/Mr. Stephen Antwi-Asimeng \' non-excutive director\'.png', 'board', 3),
('Mrs. Wendy Malm', 'Board Secretary', 'Experienced professional ensuring effective board governance and compliance.', '/Board of directors/Mrs. Wendy Malm Board Secretary.png', 'board', 4),

-- Executive Heads
('Mr. Robert Dowuona Owoo', 'Acting Chief Executive Officer', 'Experienced leader with extensive background in commodity trading and financial markets.', '/Mr. Robert Dowuona Owoo.jpeg', 'executive', 1),
('Ms. Ophelia Martekuo Atoklo', 'Acting Deputy Chief Executive Officer', 'Strategic leader with deep expertise in operational excellence and stakeholder management.', '/Ms. Ophelia Martekuo Atoklo \'Deputy Chief Executive Officer\'.jpg', 'executive', 2),

-- Functional Heads
('Mr. Opoku Debrah', 'Internal Auditor', 'Ensuring compliance and risk management across all organizational processes.', '/Functional Heads/Mr. Opoku Debrah (Internal Auditor).jpg', 'functional', 1),
('Mr. Richard Ankrah', 'Special Project', 'Leading strategic initiatives and special projects to drive organizational growth.', '/Functional Heads/Mr. Richard Ankrah ( Special Project).jpg', 'functional', 2),
('Mr. Vitus Ninfaakang', 'Risk, Membership and Partnership', 'Managing risk assessment and building strategic partnerships with key stakeholders.', '/Functional Heads/Mr. Vitus Ninfaakang (Risk, Membership and Partnership).jpg', 'functional', 3),
('Mrs. Jemimah Naa Adjeley Oppong-Gyamfi', 'Corporate Services', 'Overseeing corporate governance and administrative excellence.', '/Functional Heads/Mrs. Jemimah Naa Adjeley Oppong-Gyamfi ( Corporate Services).jpg', 'functional', 4),
('Mrs. Wendy Malm', 'Operations', 'Ensuring smooth operational processes and service delivery excellence.', '/Functional Heads/Mrs. Wendy Malm (Operations).jpg', 'functional', 5),
('Mr. Godfred Kofi Nyamekye', 'Value Chain and Product Development', 'Driving innovation in product development and value chain optimization.', '/Functional Heads/Mr. Godfred Kofi Nyamekye (Value Chain and Product Development).jpg', 'functional', 6),
('Mr. Gabriel Aryeetey', 'Warehouse & Quality', 'Maintaining quality standards and efficient warehouse operations.', '/Functional Heads/Mr. Gabriel Aryeetey (Warehouse & Quality).jpg', 'functional', 7),
('Dr. Harold Okai-Tettey', 'Information Technology and Information System', 'Leading digital transformation and technology infrastructure development.', '/Functional Heads/Dr. Harold Okai-Tettey ( Information Technology and Information System).jpg', 'functional', 8),
('Mr. Albert Nii Ayi Tagoe', 'Finance and Investments', 'Managing financial strategy and investment opportunities for sustainable growth.', '/Functional Heads/Mr. Albert Nii Ayi Tagoe (Finance and Investments).jpg', 'functional', 9);
