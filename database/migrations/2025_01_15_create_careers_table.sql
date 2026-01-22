-- Create careers table
CREATE TABLE IF NOT EXISTS careers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    category ENUM('Job Openings', 'Internship', 'Job Functional Areas') NOT NULL,
    department VARCHAR(255),
    location VARCHAR(255),
    employment_type ENUM('Full-time', 'Part-time', 'Contract', 'Internship') NOT NULL,
    experience_level ENUM('Entry Level', 'Mid Level', 'Senior Level', 'Executive') NOT NULL,
    requirements TEXT,
    responsibilities TEXT,
    benefits TEXT,
    salary_range VARCHAR(100),
    application_deadline DATE,
    start_date DATE,
    status ENUM('Open', 'Closed', 'On Hold') DEFAULT 'Open',
    application_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT INTO careers (title, description, category, department, location, employment_type, experience_level, requirements, responsibilities, benefits, salary_range, application_deadline, start_date, status) VALUES
('Senior Market Analyst', 'Lead market analysis and research for commodity trading', 'Job Openings', 'Market Research', 'Accra, Ghana', 'Full-time', 'Senior Level', 'Masters in Economics, 5+ years experience, Strong analytical skills', 'Analyze market trends, Prepare reports, Lead research projects', 'Health insurance, Pension, Professional development', 'GHC 8,000 - GHC 12,000', '2025-02-15', '2025-03-01', 'Open'),
('IT Support Intern', 'Support IT operations and system maintenance', 'Internship', 'Information Technology', 'Accra, Ghana', 'Internship', 'Entry Level', 'Computer Science degree, Basic programming knowledge', 'Support users, Maintain systems, Learn new technologies', 'Stipend, Mentorship, Certificate', 'GHC 1,500 - GHC 2,000', '2025-01-31', '2025-02-15', 'Open'),
('Trading Operations', 'Manage day-to-day trading operations and member services', 'Job Functional Areas', 'Trading Operations', 'Accra, Ghana', 'Full-time', 'Mid Level', 'Finance background, Trading experience preferred', 'Execute trades, Manage member accounts, Monitor markets', 'Competitive salary, Performance bonus, Career growth', 'GHC 5,000 - GHC 8,000', '2025-02-28', '2025-03-15', 'Open'),
('Marketing Coordinator', 'Coordinate marketing activities and member engagement', 'Job Openings', 'Marketing', 'Accra, Ghana', 'Full-time', 'Mid Level', 'Marketing degree, 3+ years experience, Creative thinking', 'Develop campaigns, Manage social media, Organize events', 'Health insurance, Flexible hours, Team building', 'GHC 4,000 - GHC 6,000', '2025-02-20', '2025-03-01', 'Open'),
('Finance Intern', 'Assist with financial reporting and analysis', 'Internship', 'Finance', 'Accra, Ghana', 'Internship', 'Entry Level', 'Accounting/Finance degree, Excel proficiency', 'Prepare reports, Assist with audits, Data analysis', 'Stipend, Learning opportunities, Reference letter', 'GHC 1,200 - GHC 1,800', '2025-01-25', '2025-02-01', 'Open');
