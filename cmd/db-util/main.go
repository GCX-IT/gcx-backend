package main

import (
	"fmt"
	"log"
	"os"

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

	action := ""
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	switch action {
	case "list-pages":
		listPages()
	case "fix-about-page":
		fixAboutPage()
	case "create-sample-settings":
		createSampleSettings()
	default:
		fmt.Println("Available commands:")
		fmt.Println("  list-pages           - List all pages in database")
		fmt.Println("  fix-about-page       - Fix the about page status")
		fmt.Println("  create-sample-settings - Create sample settings")
	}
}

func listPages() {
	var pages []models.Page
	if err := config.DB.Find(&pages).Error; err != nil {
		log.Fatal("Failed to fetch pages:", err)
	}

	fmt.Printf("Found %d pages:\n", len(pages))
	for _, page := range pages {
		fmt.Printf("- ID: %d, Title: %s, Slug: %s, Status: %s\n", 
			page.ID, page.Title, page.Slug, page.Status)
	}
}

func fixAboutPage() {
	var page models.Page
	if err := config.DB.Where("slug = ?", "about").First(&page).Error; err != nil {
		log.Printf("About page not found, creating new one...")
		
		// Create new about page
		page = models.Page{
			Title:           "About GCX",
			Slug:            "about",
			Content:         `{"hero_title":"What is GCX","hero_subtitle":"A regulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.","hero_image":"/trading dashboard.jpg","about_title":"About GCX","about_description":"The Ghana Commodity Exchange is a private company limited by shares, structured as a Public-Private Partnership, with the Government of Ghana currently the sole shareholder.","ceo_name":"Mr. Robert Dowuona Owoo","ceo_title":"Acting Chief Executive Officer","ceo_image":"/Mr. Robert Dowuona Owoo.jpeg","key_goal_title":"Our Key Goal","key_goal_description":"To link Ghanaian smallholder farmers to agricultural and financial markets in Ghana and across the West Africa Region to ensure Ghana farmers secure competitive prices for their commodities."}`,
			Excerpt:         "Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.",
			Template:        "default",
			Status:          models.PageStatusPublished,
			MetaTitle:       "About GCX - Ghana Commodity Exchange",
			MetaDescription: "Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.",
			MetaKeywords:    "about, gcx, ghana commodity exchange, mission, vision, leadership",
			AuthorID:        1, // Assuming admin user has ID 1
		}
		
		if err := config.DB.Create(&page).Error; err != nil {
			log.Fatal("Failed to create about page:", err)
		}
		
		fmt.Println("✅ Created new about page with published status")
		return
	}

	// Update existing page status
	page.Status = models.PageStatusPublished
	
	if err := config.DB.Save(&page).Error; err != nil {
		log.Fatal("Failed to update about page:", err)
	}

	fmt.Printf("✅ Updated about page (ID: %d) status to published\n", page.ID)
}

func createSampleSettings() {
	settings := []models.Setting{
		{
			Key:         "about_content",
			Value:       `{"hero_title":"What is GCX","hero_subtitle":"A regulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.","hero_image":"/trading dashboard.jpg","about_title":"About GCX","about_description":"The Ghana Commodity Exchange is a private company limited by shares, structured as a Public-Private Partnership, with the Government of Ghana currently the sole shareholder.","ceo_name":"Mr. Robert Dowuona Owoo","ceo_title":"Acting Chief Executive Officer","ceo_image":"/Mr. Robert Dowuona Owoo.jpeg","key_goal_title":"Our Key Goal","key_goal_description":"To link Ghanaian smallholder farmers to agricultural and financial markets in Ghana and across the West Africa Region to ensure Ghana farmers secure competitive prices for their commodities."}`,
			Type:        "json",
			Group:       "pages",
			Label:       "About GCX Content",
			Description: "Content for the About GCX page",
			IsPublic:    true,
		},
		{
			Key:         "site_name",
			Value:       "Ghana Commodity Exchange",
			Type:        "string",
			Group:       "general",
			Label:       "Site Name",
			Description: "The name of the website",
			IsPublic:    true,
		},
		{
			Key:         "site_tagline",
			Value:       "Linking Farmers to Markets",
			Type:        "string",
			Group:       "general",
			Label:       "Site Tagline",
			Description: "The tagline of the website",
			IsPublic:    true,
		},
	}

	for _, setting := range settings {
		// Check if setting already exists
		var existingSetting models.Setting
		if err := config.DB.Where("`key` = ?", setting.Key).First(&existingSetting).Error; err == nil {
			fmt.Printf("⚠️  Setting '%s' already exists, skipping\n", setting.Key)
			continue
		}

		if err := config.DB.Create(&setting).Error; err != nil {
			log.Printf("Failed to create setting '%s': %v", setting.Key, err)
		} else {
			fmt.Printf("✅ Created setting: %s\n", setting.Key)
		}
	}
}
