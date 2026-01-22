package main

import (
	"log"

	cms_models "gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/config"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	config.InitDB()

	log.Println("🔄 Running database migration...")

	// Force create the settings table
	if err := config.DB.AutoMigrate(&cms_models.Setting{}); err != nil {
		log.Fatal("Failed to migrate settings table:", err)
	}

	log.Println("✅ Settings table created successfully")

	// Create some initial settings
	initialSettings := []cms_models.Setting{
		{
			Key:         "site_name",
			Value:       "Ghana Commodity Exchange",
			Type:        cms_models.SettingTypeText,
			Group:       "general",
			Label:       "Site Name",
			Description: "The name of your website",
			IsPublic:    true,
			SortOrder:   1,
		},
		{
			Key:         "site_tagline",
			Value:       "Connecting Markets, Connecting People, Providing Opportunities",
			Type:        cms_models.SettingTypeText,
			Group:       "general",
			Label:       "Site Tagline",
			Description: "Your site tagline or motto",
			IsPublic:    true,
			SortOrder:   2,
		},
		{
			Key:         "site_logo",
			Value:       "/logo_black.png",
			Type:        cms_models.SettingTypeImage,
			Group:       "general",
			Label:       "Site Logo",
			Description: "Your site logo",
			IsPublic:    true,
			SortOrder:   3,
		},
		{
			Key:         "contact_email",
			Value:       "info@gcx.com.gh",
			Type:        cms_models.SettingTypeEmail,
			Group:       "contact",
			Label:       "Contact Email",
			Description: "Main contact email address",
			IsPublic:    true,
			SortOrder:   1,
		},
		{
			Key:         "contact_phone",
			Value:       "+233 302 123 456",
			Type:        cms_models.SettingTypeText,
			Group:       "contact",
			Label:       "Contact Phone",
			Description: "Main contact phone number",
			IsPublic:    true,
			SortOrder:   2,
		},
	}

	// Insert initial settings
	for _, setting := range initialSettings {
		var existingSetting cms_models.Setting
		if err := config.DB.Where("`key` = ?", setting.Key).First(&existingSetting).Error; err != nil {
			// Setting doesn't exist, create it
			if err := config.DB.Create(&setting).Error; err != nil {
				log.Printf("Failed to create setting %s: %v", setting.Key, err)
			} else {
				log.Printf("✅ Created setting: %s", setting.Key)
			}
		} else {
			log.Printf("⚠️  Setting %s already exists, skipping", setting.Key)
		}
	}

	log.Println("🎉 Database migration completed successfully!")
	log.Println("📊 You can now use the CMS settings functionality")
}
