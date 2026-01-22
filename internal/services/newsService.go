package services

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/database"

	"gorm.io/gorm"
)

// NewsService handles news operations and Firebase integration
type NewsService struct {
	db       *gorm.DB
	firebase *FirebaseService
}

// NewNewsService creates a new news service instance
func NewNewsService() *NewsService {
	return &NewsService{
		db:       database.GetDB(),
		firebase: GetFirebaseService(),
	}
}

// CreateNewsItem creates a new news item and optionally publishes to Firebase
func (ns *NewsService) CreateNewsItem(newsItem *models.NewsItem, publishToFirebase bool) error {
	// Create in database
	if err := ns.db.Create(newsItem).Error; err != nil {
		return fmt.Errorf("failed to create news item: %v", err)
	}

	// If published and Firebase is enabled, publish to Firebase
	if publishToFirebase && newsItem.IsPublished() && ns.firebase != nil {
		if err := ns.firebase.PublishNewsToFirebase(newsItem); err != nil {
			log.Printf("Warning: Failed to publish news to Firebase: %v", err)
			// Don't fail the entire operation if Firebase fails
		}
	}

	log.Printf("Created news item: %s (ID: %d)", newsItem.Title, newsItem.ID)
	return nil
}

// UpdateNewsItem updates a news item and syncs with Firebase
func (ns *NewsService) UpdateNewsItem(newsItem *models.NewsItem) error {
	// Update in database
	if err := ns.db.Save(newsItem).Error; err != nil {
		return fmt.Errorf("failed to update news item: %v", err)
	}

	// Sync with Firebase if it's published
	if newsItem.IsPublished() && ns.firebase != nil {
		if err := ns.firebase.UpdateNewsInFirebase(newsItem); err != nil {
			log.Printf("Warning: Failed to update news in Firebase: %v", err)
		}
	}

	log.Printf("Updated news item: %s (ID: %d)", newsItem.Title, newsItem.ID)
	return nil
}

// DeleteNewsItem deletes a news item and removes from Firebase
func (ns *NewsService) DeleteNewsItem(id uint) error {
	var newsItem models.NewsItem
	if err := ns.db.First(&newsItem, id).Error; err != nil {
		return fmt.Errorf("failed to find news item: %v", err)
	}

	// Delete from database
	if err := ns.db.Delete(&newsItem).Error; err != nil {
		return fmt.Errorf("failed to delete news item: %v", err)
	}

	// Remove from Firebase
	if ns.firebase != nil {
		if err := ns.firebase.DeleteNewsFromFirebase(id); err != nil {
			log.Printf("Warning: Failed to delete news from Firebase: %v", err)
		}
	}

	log.Printf("Deleted news item: %s (ID: %d)", newsItem.Title, newsItem.ID)
	return nil
}

// PublishNewsItem publishes a news item and syncs with Firebase
func (ns *NewsService) PublishNewsItem(id uint) error {
	var newsItem models.NewsItem
	if err := ns.db.First(&newsItem, id).Error; err != nil {
		return fmt.Errorf("failed to find news item: %v", err)
	}

	newsItem.Publish()

	// Update in database
	if err := ns.db.Save(&newsItem).Error; err != nil {
		return fmt.Errorf("failed to publish news item: %v", err)
	}

	// Publish to Firebase
	if ns.firebase != nil {
		if err := ns.firebase.PublishNewsToFirebase(&newsItem); err != nil {
			log.Printf("Warning: Failed to publish news to Firebase: %v", err)
		}
	}

	log.Printf("Published news item: %s (ID: %d)", newsItem.Title, newsItem.ID)
	return nil
}

// ArchiveNewsItem archives a news item and updates Firebase
func (ns *NewsService) ArchiveNewsItem(id uint) error {
	var newsItem models.NewsItem
	if err := ns.db.First(&newsItem, id).Error; err != nil {
		return fmt.Errorf("failed to find news item: %v", err)
	}

	newsItem.Archive()

	// Update in database
	if err := ns.db.Save(&newsItem).Error; err != nil {
		return fmt.Errorf("failed to archive news item: %v", err)
	}

	// Update in Firebase
	if ns.firebase != nil {
		if err := ns.firebase.UpdateNewsInFirebase(&newsItem); err != nil {
			log.Printf("Warning: Failed to update news in Firebase: %v", err)
		}
	}

	log.Printf("Archived news item: %s (ID: %d)", newsItem.Title, newsItem.ID)
	return nil
}

// SetBreakingNews sets a news item as breaking news
func (ns *NewsService) SetBreakingNews(id uint, isBreaking bool) error {
	var newsItem models.NewsItem
	if err := ns.db.First(&newsItem, id).Error; err != nil {
		return fmt.Errorf("failed to find news item: %v", err)
	}

	newsItem.SetBreaking(isBreaking)

	// Update in database
	if err := ns.db.Save(&newsItem).Error; err != nil {
		return fmt.Errorf("failed to update breaking news status: %v", err)
	}

	// Update in Firebase if published
	if newsItem.IsPublished() && ns.firebase != nil {
		if err := ns.firebase.UpdateNewsInFirebase(&newsItem); err != nil {
			log.Printf("Warning: Failed to update news in Firebase: %v", err)
		}
	}

	log.Printf("Set breaking news status for item: %s (ID: %d, Breaking: %v)", newsItem.Title, newsItem.ID, isBreaking)
	return nil
}

// SyncAllPublishedNews syncs all published news items to Firebase
func (ns *NewsService) SyncAllPublishedNews() error {
	if ns.firebase == nil {
		return fmt.Errorf("Firebase service not initialized")
	}

	var newsItems []models.NewsItem
	if err := ns.db.Where("status = ? AND is_active = ?", models.NewsStatusPublished, true).
		Find(&newsItems).Error; err != nil {
		return fmt.Errorf("failed to fetch published news items: %v", err)
	}

	// Convert []models.NewsItem to []*models.NewsItem
	var newsItemPtrs []*models.NewsItem
	for i := range newsItems {
		newsItemPtrs = append(newsItemPtrs, &newsItems[i])
	}

	if err := ns.firebase.PublishNewsBatch(newsItemPtrs); err != nil {
		return fmt.Errorf("failed to sync news to Firebase: %v", err)
	}

	log.Printf("Synced %d published news items to Firebase", len(newsItems))
	return nil
}

// GetActiveNewsItems returns active news items for the ticker
func (ns *NewsService) GetActiveNewsItems(limit int, source string, category string, breakingOnly bool) ([]models.NewsItem, error) {
	var newsItems []models.NewsItem

	query := ns.db.Where("status = ? AND is_active = ?", models.NewsStatusPublished, true).
		Where("(published_at IS NULL OR published_at <= ?)", time.Now()).
		Where("(expires_at IS NULL OR expires_at > ?)", time.Now())

	if source != "" {
		query = query.Where("source = ?", source)
	}

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if breakingOnly {
		query = query.Where("is_breaking = ?", true)
	}

	err := query.Order("is_breaking DESC, priority DESC, published_at DESC").
		Limit(limit).
		Find(&newsItems).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch active news items: %v", err)
	}

	return newsItems, nil
}

// GetBreakingNews returns only breaking news items
func (ns *NewsService) GetBreakingNews(limit int) ([]models.NewsItem, error) {
	return ns.GetActiveNewsItems(limit, "", "", true)
}

// CreateNewsFromExternalSource creates a news item from external source data
func (ns *NewsService) CreateNewsFromExternalSource(sourceData map[string]interface{}, source models.NewsSource) (*models.NewsItem, error) {
	newsItem := &models.NewsItem{
		Source: source,
		Status: models.NewsStatusDraft,
	}

	// Extract common fields
	if title, ok := sourceData["title"].(string); ok {
		newsItem.Title = title
	}

	if content, ok := sourceData["content"].(string); ok {
		newsItem.Content = content
	}

	if sourceName, ok := sourceData["source_name"].(string); ok {
		newsItem.SourceName = &sourceName
	}

	if sourceURL, ok := sourceData["source_url"].(string); ok {
		newsItem.SourceURL = &sourceURL
	}

	if category, ok := sourceData["category"].(string); ok {
		newsItem.Category = &category
	}

	if priority, ok := sourceData["priority"].(float64); ok {
		newsItem.Priority = int(priority)
	}

	if isBreaking, ok := sourceData["is_breaking"].(bool); ok {
		newsItem.IsBreaking = isBreaking
	}

	// Store external data as JSON
	if externalDataBytes, err := json.Marshal(sourceData); err == nil {
		externalData := string(externalDataBytes)
		newsItem.ExternalData = &externalData
	}

	// Set external ID if provided
	if externalID, ok := sourceData["external_id"].(string); ok {
		newsItem.ExternalID = &externalID
	}

	now := time.Now()
	newsItem.LastSyncAt = &now

	return newsItem, nil
}

// SyncFromExternalAPI syncs news from an external API
func (ns *NewsService) SyncFromExternalAPI(apiEndpoint string, source models.NewsSource) error {
	// This is a placeholder for external API integration
	// In a real implementation, you would:
	// 1. Make HTTP request to the API endpoint
	// 2. Parse the response
	// 3. Create news items from the data
	// 4. Handle duplicates and updates

	log.Printf("Syncing news from external API: %s", apiEndpoint)

	// Example implementation:
	// resp, err := http.Get(apiEndpoint)
	// if err != nil {
	//     return fmt.Errorf("failed to fetch from API: %v", err)
	// }
	// defer resp.Body.Close()
	//
	// var apiResponse map[string]interface{}
	// if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
	//     return fmt.Errorf("failed to decode API response: %v", err)
	// }
	//
	// // Process each news item from the API
	// for _, itemData := range apiResponse["items"].([]interface{}) {
	//     newsItem, err := ns.CreateNewsFromExternalSource(itemData.(map[string]interface{}), source)
	//     if err != nil {
	//         log.Printf("Error creating news item: %v", err)
	//         continue
	//     }
	//
	//     // Check if item already exists (by external_id)
	//     var existing models.NewsItem
	//     if err := ns.db.Where("external_id = ? AND source = ?", newsItem.ExternalID, source).First(&existing).Error; err == nil {
	//         // Update existing item
	//         existing.Title = newsItem.Title
	//         existing.Content = newsItem.Content
	//         existing.ExternalData = newsItem.ExternalData
	//         existing.LastSyncAt = newsItem.LastSyncAt
	//         ns.db.Save(&existing)
	//     } else {
	//         // Create new item
	//         ns.db.Create(newsItem)
	//     }
	// }

	return nil
}
