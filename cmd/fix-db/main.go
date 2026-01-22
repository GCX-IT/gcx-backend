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
	case "fix-slugs":
		fixSlugs()
	case "cleanup-duplicates":
		cleanupDuplicates()
	default:
		fmt.Println("Available commands:")
		fmt.Println("  fix-slugs           - Fix page slugs to match frontend expectations")
		fmt.Println("  cleanup-duplicates  - Remove duplicate pages")
	}
}

func fixSlugs() {
	// Fix the about page slug
	var aboutPage models.Page
	if err := config.DB.Where("title = ? AND slug LIKE ?", "About GCX", "about%").First(&aboutPage).Error; err != nil {
		log.Fatal("Failed to find about page:", err)
	}

	// Update the slug to exactly "about"
	aboutPage.Slug = "about"
	if err := config.DB.Save(&aboutPage).Error; err != nil {
		log.Fatal("Failed to update about page slug:", err)
	}

	fmt.Printf("✅ Updated page ID %d slug to 'about'\n", aboutPage.ID)

	// List all pages to confirm
	var pages []models.Page
	if err := config.DB.Find(&pages).Error; err != nil {
		log.Fatal("Failed to fetch pages:", err)
	}

	fmt.Printf("\nCurrent pages:\n")
	for _, page := range pages {
		fmt.Printf("- ID: %d, Title: %s, Slug: %s, Status: %s\n", 
			page.ID, page.Title, page.Slug, page.Status)
	}
}

func cleanupDuplicates() {
	// Keep only one About GCX page and delete the rest
	var aboutPages []models.Page
	if err := config.DB.Where("title = ?", "About GCX").Find(&aboutPages).Error; err != nil {
		log.Fatal("Failed to find about pages:", err)
	}

	if len(aboutPages) <= 1 {
		fmt.Println("No duplicate pages found")
		return
	}

	// Keep the first one, delete the rest
	keepPage := aboutPages[0]
	keepPage.Slug = "about" // Ensure correct slug
	if err := config.DB.Save(&keepPage).Error; err != nil {
		log.Fatal("Failed to update kept page:", err)
	}

	fmt.Printf("✅ Keeping page ID %d with slug 'about'\n", keepPage.ID)

	// Delete duplicates
	for i := 1; i < len(aboutPages); i++ {
		if err := config.DB.Delete(&aboutPages[i]).Error; err != nil {
			log.Printf("Failed to delete page ID %d: %v", aboutPages[i].ID, err)
		} else {
			fmt.Printf("🗑️  Deleted duplicate page ID %d (slug: %s)\n", aboutPages[i].ID, aboutPages[i].Slug)
		}
	}

	fmt.Println("✅ Cleanup completed")
}
