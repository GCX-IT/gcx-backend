-- Add essential social media columns to existing team_members table
ALTER TABLE team_members 
ADD COLUMN linkedin_url VARCHAR(500),
ADD COLUMN twitter_url VARCHAR(500),
ADD COLUMN facebook_url VARCHAR(500),
ADD COLUMN instagram_url VARCHAR(500);
