package main

import (
	"fmt"
	"log"
	"time"

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

	// Restore the soft-deleted page with slug "about"
	var page models.Page
	result := config.DB.Unscoped().Where("slug = ?", "about").First(&page)
	
	if result.Error != nil {
		log.Fatal("No page found with slug 'about':", result.Error)
	}

	fmt.Printf("Found page: ID=%d, Title=%s, Status=%s, DeletedAt=%v\n", 
		page.ID, page.Title, page.Status, page.DeletedAt)

	if page.DeletedAt.Valid {
		// Restore the page by setting DeletedAt to nil
		if err := config.DB.Unscoped().Model(&page).Update("deleted_at", nil).Error; err != nil {
			log.Fatal("Failed to restore page:", err)
		}
		fmt.Println("✅ Page restored from soft delete")
	}

	// Make sure it's published
	now := time.Now()
	updates := map[string]interface{}{
		"status": models.PageStatusPublished,
		"published_at": &now,
	}
	
	if err := config.DB.Unscoped().Model(&page).Where("id = ?", page.ID).Updates(updates).Error; err != nil {
		log.Fatal("Failed to update page:", err)
	}
	
	fmt.Println("✅ Page status updated to published")

	// Verify the page is now accessible
	var verifyPage models.Page
	if err := config.DB.Where("slug = ? AND status = ?", "about", models.PageStatusPublished).First(&verifyPage).Error; err != nil {
		log.Fatal("Failed to verify page restoration:", err)
	}

	fmt.Printf("✅ Verification successful: Page ID %d with slug 'about' is now published\n", verifyPage.ID)
}
