-- Sample News Data for GCX News Ticker (Fixed Version)
-- This script handles existing categories and avoids duplicate errors

-- Insert news categories only if they don't exist (ignore duplicates)
INSERT IGNORE INTO news_categories (name, slug, created_at, updated_at) VALUES
('Market Updates', 'market-updates', NOW(), NOW()),
('Announcements', 'announcements', NOW(), NOW()),
('Partnerships', 'partnerships', NOW(), NOW()),
('Events', 'events', NOW(), NOW()),
('Regulations', 'regulations', NOW(), NOW());

-- Get the category IDs (these will work even if categories already existed)
SET @market_category_id = (SELECT id FROM news_categories WHERE slug = 'market-updates' LIMIT 1);
SET @announcement_category_id = (SELECT id FROM news_categories WHERE slug = 'announcements' LIMIT 1);
SET @partnership_category_id = (SELECT id FROM news_categories WHERE slug = 'partnerships' LIMIT 1);
SET @event_category_id = (SELECT id FROM news_categories WHERE slug = 'events' LIMIT 1);

-- Insert sample news items (ignore duplicates based on title)
INSERT IGNORE INTO news_items (
    title, 
    content, 
    source, 
    source_name, 
    source_url, 
    category_id, 
    tags, 
    status, 
    is_breaking, 
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
    @market_category_id,
    'trading,volume,record,quarterly',
    'published',
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
    @announcement_category_id,
    'maize,contracts,specifications,trading',
    'published',
    0,
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
    @partnership_category_id,
    'partnership,adb,financing,credit',
    'published',
    0,
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
    @market_category_id,
    'rice,prices,analysis,market-trends',
    'published',
    0,
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
    @event_category_id,
    'agm,meeting,events,calendar',
    'published',
    0,
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
    @announcement_category_id,
    'soybean,trading-hours,international,market',
    'published',
    0,
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
    @market_category_id,
    'wheat,prices,international,weather',
    'published',
    0,
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
    @announcement_category_id,
    'mobile-app,technology,trading,innovation',
    'published',
    0,
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
    @announcement_category_id,
    'maintenance,closure,system,trading-suspension',
    'published',
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
    @partnership_category_id,
    'farmers,partnership,market-access,pricing',
    'published',
    0,
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
    nc.name as category,
    ni.status,
    ni.is_breaking,
    ni.published_at
FROM news_items ni
LEFT JOIN news_categories nc ON ni.category_id = nc.id
WHERE ni.status = 'published'
ORDER BY ni.published_at DESC;

-- Show count by status
SELECT 
    status,
    COUNT(*) as count
FROM news_items 
GROUP BY status;

-- Show breaking news count
SELECT 
    COUNT(*) as breaking_news_count
FROM news_items 
WHERE is_breaking = 1 AND status = 'published';

-- Show total published news for ticker
SELECT 
    COUNT(*) as total_published_news
FROM news_items 
WHERE status = 'published';
