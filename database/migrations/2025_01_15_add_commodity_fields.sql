-- Add missing fields to commodities table for frontend compatibility
ALTER TABLE commodities 
ADD COLUMN full_description TEXT AFTER description,
ADD COLUMN contract_file VARCHAR(500) AFTER image_path,
ADD COLUMN delivery_months VARCHAR(255) AFTER harvest_season;

-- Update existing records with full descriptions using ID
UPDATE commodities SET full_description = CONCAT(description, ' The commodity exchange provides a platform for farmers to access better prices and reduce post-harvest losses through proper storage and marketing. This commodity is essential for food security and economic development in Ghana.') WHERE id = 1;

UPDATE commodities SET full_description = CONCAT(description, ' The commodity exchange provides a platform for soybean farmers to access better markets and pricing, supporting the development of the legume sector and contributing to food security and nutrition.') WHERE id = 2;

UPDATE commodities SET full_description = CONCAT(description, ' The commodity exchange facilitates better market access for sorghum farmers, helping them get fair prices for their produce and contributing to food security in the region.') WHERE id = 3;

UPDATE commodities SET full_description = CONCAT(description, ' The commodity exchange provides a platform for sesame farmers to access better markets and pricing, contributing to rural development and poverty reduction.') WHERE id = 4;

UPDATE commodities SET full_description = CONCAT(description, ' Rice trading on the commodity exchange provides price discovery and market access for both producers and consumers, ensuring food security and fair pricing.') WHERE id = 5;

-- Add contract files using ID
UPDATE commodities SET contract_file = '/gcx-online-trading-member.pdf' WHERE id IN (1, 2, 3, 4, 5);

-- Add delivery months using ID
UPDATE commodities SET delivery_months = 'March, May, July, September' WHERE id = 1;
UPDATE commodities SET delivery_months = 'April, June, August, October' WHERE id = 2;
UPDATE commodities SET delivery_months = 'March, May, July, September' WHERE id = 3;
UPDATE commodities SET delivery_months = 'February, April, June, August' WHERE id = 4;
UPDATE commodities SET delivery_months = 'April, June, August, October' WHERE id = 5;
