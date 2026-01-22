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

	// Check for any page with slug "about" including soft deleted
	var page models.Page
	result := config.DB.Unscoped().Where("slug = ?", "about").First(&page)
	
	if result.Error != nil {
		fmt.Println("No page found with slug 'about'")
		fmt.Println("Creating new page with slug 'about'...")
		
		// Delete all duplicate about pages first
		if err := config.DB.Where("title = ?", "About GCX").Delete(&models.Page{}).Error; err != nil {
			log.Printf("Warning: Failed to delete existing pages: %v", err)
		}
		
		// Create new about page
		newPage := models.Page{
			Title:           "About GCX",
			Slug:            "about",
			Content:         `{"hero_title":"What is GCX","hero_subtitle":"A regulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.","hero_image":"/trading dashboard.jpg","about_title":"About GCX","about_description":"The Ghana Commodity Exchange is a private company limited by shares, structured as a Public-Private Partnership, with the Government of Ghana currently the sole shareholder.","ceo_name":"Mr. Robert Dowuona Owoo","ceo_title":"Acting Chief Executive Officer","ceo_image":"/Mr. Robert Dowuona Owoo.jpeg","key_goal_title":"Our Key Goal","key_goal_description":"To link Ghanaian smallholder farmers to agricultural and financial markets in Ghana and across the West Africa Region to ensure Ghana farmers secure competitive prices for their commodities."}`,
			Excerpt:         "Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.",
			Template:        "default",
			Status:          models.PageStatusPublished,
			MetaTitle:       "About GCX - Ghana Commodity Exchange",
			MetaDescription: "Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.",
			MetaKeywords:    "about, gcx, ghana commodity exchange, mission, vision, leadership",
			AuthorID:        1,
		}
		
		if err := config.DB.Create(&newPage).Error; err != nil {
			log.Fatal("Failed to create about page:", err)
		}
		
		fmt.Printf("✅ Created new about page with ID: %d\n", newPage.ID)
	} else {
		fmt.Printf("Found page with slug 'about': ID=%d, Title=%s, Status=%s, DeletedAt=%v\n", 
			page.ID, page.Title, page.Status, page.DeletedAt)
		
		if page.DeletedAt.Valid {
			fmt.Println("Page is soft deleted, restoring it...")
			page.DeletedAt = page.DeletedAt
			if err := config.DB.Unscoped().Save(&page).Error; err != nil {
				log.Fatal("Failed to restore page:", err)
			}
			fmt.Println("✅ Page restored")
		}
		
		// Make sure it's published
		if page.Status != models.PageStatusPublished {
			page.Status = models.PageStatusPublished
			if err := config.DB.Save(&page).Error; err != nil {
				log.Fatal("Failed to publish page:", err)
			}
			fmt.Println("✅ Page status updated to published")
		}
	}
	
	// List final state
	var pages []models.Page
	if err := config.DB.Find(&pages).Error; err != nil {
		log.Fatal("Failed to fetch pages:", err)
	}

	fmt.Printf("\nFinal pages list:\n")
	for _, p := range pages {
		fmt.Printf("- ID: %d, Title: %s, Slug: %s, Status: %s\n", 
			p.ID, p.Title, p.Slug, p.Status)
	}
}
