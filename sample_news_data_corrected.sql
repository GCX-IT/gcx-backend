-- Sample News Data for GCX News Ticker (Corrected Version)
-- This script uses the correct table structure from the migration

-- Insert news categories only if they don't exist (ignore duplicates)
INSERT IGNORE INTO news_categories (name, slug, description, color, created_at, updated_at) VALUES
('Market Updates', 'market-updates', 'Latest market news and price updates', '#3B82F6', NOW(), NOW()),
('Announcements', 'announcements', 'Important announcements from GCX', '#10B981', NOW(), NOW()),
('Partnerships', 'partnerships', 'Partnership and collaboration news', '#8B5CF6', NOW(), NOW()),
('Events', 'events', 'Upcoming events and programs', '#F59E0B', NOW(), NOW()),
('Regulations', 'regulations', 'Regulatory updates and compliance', '#EF4444', NOW(), NOW());

-- Insert sample news items (using correct column names)
INSERT IGNORE INTO news_items (
    title, 
    content, 
    source, 
    source_name, 
    source_url, 
    category, 
    priority, 
    status, 
    is_breaking, 
    is_active,
    published_at, 
    created_at, 
    updated_at
) VALUES
(
    'GCX Records Highest Trading Volume in Q1 2025',
    'The Ghana Commodity Exchange achieved a record-breaking trading volume of $2.5 million in the first quarter of 2025, representing a 35% increase from the previous quarter. This milestone demonstrates the growing confidence of traders and farmers in the exchange platform.',
    'gcx',
    'GCX Official',
    'https://gcx.com/news/q1-2025-results',
    'market',
    9,
    'published',
    1,
    1,
    NOW(),
    NOW(),
    NOW()
),
(
    'New Maize Contract Specifications Released',
    'GCX has released updated contract specifications for maize trading, including new quality standards and delivery requirements effective March 1, 2025. These improvements aim to enhance market efficiency and ensure better price discovery.',
    'gcx',
    'GCX Trading',
    'https://gcx.com/contracts/maize-specs-2025',
    'announcement',
    7,
    'published',
    0,
    1,
    NOW(),
    NOW(),
    NOW()
),
(
    'Strategic Partnership with Agricultural Development Bank',
    'GCX announces strategic partnership with ADB to provide enhanced financing solutions for commodity traders and farmers across Ghana. This collaboration will improve access to credit and reduce transaction costs.',
    'gcx',
    'GCX Partnerships',
    'https://gcx.com/partnerships/adb-collaboration',
    'partnership',
    8,
    'published',
    0,
    1,
    NOW(),
    NOW(),
    NOW()
),
(
    'Market Analysis: Rice Prices Show Steady Growth',
    'Recent market analysis indicates a 12% increase in rice prices over the past month, driven by increased demand and seasonal factors. Experts predict continued growth in the coming weeks.',
    'partner',
    'Market Research Institute',
    'https://example.com/market-analysis-rice',
    'market',
    6,
    'published',
    0,
    1,
    NOW(),
    NOW(),
    NOW()
),
(
    'GCX Annual General Meeting 2025 Scheduled',
    'Mark your calendars for the GCX Annual General Meeting scheduled for March 15, 2025, at the Accra International Conference Centre. Registration opens February 1st.',
    'gcx',
    'GCX Events',
    'https://gcx.com/events/agm-2025',
    'event',
    5,
    'published',
    0,
    1,
    NOW(),
    NOW(),
    NOW()
),
(
    'New Soybean Trading Hours Effective February 1st',
    'GCX announces extended trading hours for soybean contracts, now available from 8:00 AM to 6:00 PM GMT. This change aims to accommodate international market participants.',
    'gcx',
    'GCX Trading',
    'https://gcx.com/trading/hours-update',
    'announcement',
    6,
    'published',
    0,
    1,
    NOW(),
    NOW(),
    NOW()
),
(
    'External Market Alert: Global Wheat Prices Rising',
    'International wheat markets show significant price increases due to weather concerns in major producing regions. This may impact local wheat import costs.',
    'external',
    'International Commodity News',
    'https://example.com/wheat-price-alert',
    'market',
    5,
    'published',
    0,
    1,
    NOW(),
    NOW(),
    NOW()
),
(
    'Technology Upgrade: New Mobile Trading App Launch',
    'GCX launches its new mobile trading application with enhanced features including real-time price alerts, portfolio tracking, and secure payment integration.',
    'gcx',
    'GCX Technology',
    'https://gcx.com/mobile-app-launch',
    'announcement',
    7,
    'published',
    0,
    1,
    NOW(),
    NOW(),
    NOW()
),
(
    'Breaking: Emergency Market Closure Due to System Maintenance',
    'GCX will temporarily suspend trading operations from 2:00 AM to 4:00 AM GMT on February 5th for critical system maintenance. All pending orders will be preserved.',
    'gcx',
    'GCX Operations',
    'https://gcx.com/maintenance-notice',
    'announcement',
    10,
    'published',
    1,
    1,
    NOW(),
    NOW(),
    NOW()
),
(
    'Partnership with Local Farmers Association',
    'GCX strengthens ties with the Ghana Farmers Association to provide better market access and fair pricing for small-scale farmers across the country.',
    'gcx',
    'GCX Partnerships',
    'https://gcx.com/partnerships/farmers-association',
    'partnership',
    6,
    'published',
    0,
    1,
    NOW(),
    NOW(),
    NOW()
),
(
    'Regulatory Update: New Trading Guidelines Published',
    'The Securities and Exchange Commission has published new guidelines for commodity trading. All market participants are required to comply with these regulations by March 1st.',
    'gcx',
    'GCX Compliance',
    'https://gcx.com/regulations/trading-guidelines-2025',
    'regulation',
    8,
    'published',
    0,
    1,
    NOW(),
    NOW(),
    NOW()
),
(
    'External: International Corn Prices Reach New High',
    'Global corn prices have reached their highest level in three years due to supply chain disruptions and increased demand from ethanol production.',
    'external',
    'Global Commodity News',
    'https://example.com/corn-prices-2025',
    'market',
    4,
    'published',
    0,
    1,
    NOW(),
    NOW(),
    NOW()
);

-- Verify the data was inserted correctly
SELECT 
    ni.id,
    ni.title,
    ni.source,
    ni.source_name,
    ni.category,
    ni.status,
    ni.is_breaking,
    ni.priority,
    ni.published_at
FROM news_items ni
WHERE ni.status = 'published' AND ni.is_active = 1
ORDER BY ni.priority DESC, ni.published_at DESC;

-- Show count by status
SELECT 
    status,
    COUNT(*) as count
FROM news_items 
WHERE is_active = 1
GROUP BY status;

-- Show breaking news count
SELECT 
    COUNT(*) as breaking_news_count
FROM news_items 
WHERE is_breaking = 1 AND status = 'published' AND is_active = 1;

-- Show total published news for ticker
SELECT 
    COUNT(*) as total_published_news
FROM news_items 
WHERE status = 'published' AND is_active = 1;

-- Show news by category
SELECT 
    category,
    COUNT(*) as count
FROM news_items 
WHERE status = 'published' AND is_active = 1
GROUP BY category
ORDER BY count DESC;
