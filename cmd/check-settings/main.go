package main

import (
	"fmt"
	"log"

	"gcx-cms/internal/cms/models"
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

	// List all existing settings
	var settings []models.Setting
	if err := config.DB.Find(&settings).Error; err != nil {
		log.Fatal("Failed to fetch settings:", err)
	}

	fmt.Printf("Found %d settings:\n", len(settings))
	for _, setting := range settings {
		fmt.Printf("- Key: %s, Type: %s, Group: %s, Public: %t\n", 
			setting.Key, setting.Type, setting.Group, setting.IsPublic)
	}

	// Check specifically for about_content
	var aboutSetting models.Setting
	result := config.DB.Where("`key` = ?", "about_content").First(&aboutSetting)
	if result.Error != nil {
		fmt.Println("\n❌ about_content setting does not exist")
		
		// Try to create it
		newSetting := models.Setting{
			Key:         "about_content",
			Value:       `{"hero_title":"What is GCX","hero_subtitle":"A regulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement."}`,
			Type:        "json",
			Group:       "pages",
			Label:       "About GCX Content",
			Description: "Content for the About GCX page",
			IsPublic:    true,
		}

		if err := config.DB.Create(&newSetting).Error; err != nil {
			fmt.Printf("❌ Failed to create about_content setting: %v\n", err)
		} else {
			fmt.Println("✅ Created about_content setting successfully")
		}
	} else {
		fmt.Printf("\n✅ about_content setting exists: ID=%d, Value length=%d\n", 
			aboutSetting.ID, len(aboutSetting.Value))
	}
}
