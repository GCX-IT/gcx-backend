package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"gcx-cms/internal/cms/models"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// FirebaseService handles Firebase operations for news publishing
type FirebaseService struct {
	client *firestore.Client
	ctx    context.Context
}

// FirebaseNewsItem represents the structure of news items in Firebase
type FirebaseNewsItem struct {
	ID          string    `json:"id" firestore:"id"`
	Title       string    `json:"title" firestore:"title"`
	Content     string    `json:"content" firestore:"content"`
	Source      string    `json:"source" firestore:"source"`
	SourceName  string    `json:"source_name" firestore:"sourceName"`
	SourceURL   string    `json:"source_url" firestore:"sourceUrl"`
	Category    string    `json:"category" firestore:"category"`
	Priority    int       `json:"priority" firestore:"priority"`
	IsBreaking  bool      `json:"is_breaking" firestore:"isBreaking"`
	PublishedAt time.Time `json:"published_at" firestore:"publishedAt"`
	ExpiresAt   time.Time `json:"expires_at" firestore:"expiresAt"`
	CreatedAt   time.Time `json:"created_at" firestore:"createdAt"`
	UpdatedAt   time.Time `json:"updated_at" firestore:"updatedAt"`
}

// FirebaseConfig holds Firebase configuration
type FirebaseConfig struct {
	ProjectID   string `json:"project_id"`
	Credentials string `json:"credentials"` // Base64 encoded service account JSON
	Collection  string `json:"collection"`  // Firestore collection name
}

var firebaseService *FirebaseService

// InitializeFirebaseService initializes the Firebase service
func InitializeFirebaseService(config FirebaseConfig) error {
	ctx := context.Background()

	var client *firestore.Client
	var err error

	if config.Credentials != "" {
		// Use service account credentials
		client, err = firestore.NewClient(ctx, config.ProjectID, option.WithCredentialsJSON([]byte(config.Credentials)))
	} else {
		// Use default credentials (for local development or when running on GCP)
		client, err = firestore.NewClient(ctx, config.ProjectID)
	}

	if err != nil {
		return fmt.Errorf("failed to create Firebase client: %v", err)
	}

	firebaseService = &FirebaseService{
		client: client,
		ctx:    ctx,
	}

	return nil
}

// GetFirebaseService returns the singleton Firebase service instance
func GetFirebaseService() *FirebaseService {
	return firebaseService
}

// PublishNewsToFirebase publishes a news item to Firebase for public access
func (fs *FirebaseService) PublishNewsToFirebase(newsItem *models.NewsItem) error {
	if fs.client == nil {
		return fmt.Errorf("Firebase client not initialized")
	}

	// Convert to Firebase format
	firebaseItem := FirebaseNewsItem{
		ID:         fmt.Sprintf("%d", newsItem.ID),
		Title:      newsItem.Title,
		Content:    newsItem.Content,
		Source:     string(newsItem.Source),
		Category:   *newsItem.Category,
		Priority:   newsItem.Priority,
		IsBreaking: newsItem.IsBreaking,
		CreatedAt:  newsItem.CreatedAt,
		UpdatedAt:  newsItem.UpdatedAt,
	}

	if newsItem.SourceName != nil {
		firebaseItem.SourceName = *newsItem.SourceName
	}
	if newsItem.SourceURL != nil {
		firebaseItem.SourceURL = *newsItem.SourceURL
	}
	if newsItem.PublishedAt != nil {
		firebaseItem.PublishedAt = *newsItem.PublishedAt
	}
	if newsItem.ExpiresAt != nil {
		firebaseItem.ExpiresAt = *newsItem.ExpiresAt
	}

	// Add to Firestore
	collection := "news_items"
	if fs.client != nil {
		_, err := fs.client.Collection(collection).Doc(firebaseItem.ID).Set(fs.ctx, firebaseItem)
		if err != nil {
			return fmt.Errorf("failed to publish news to Firebase: %v", err)
		}

		log.Printf("Successfully published news item %s to Firebase", firebaseItem.ID)
	}

	return nil
}

// UpdateNewsInFirebase updates a news item in Firebase
func (fs *FirebaseService) UpdateNewsInFirebase(newsItem *models.NewsItem) error {
	if fs.client == nil {
		return fmt.Errorf("Firebase client not initialized")
	}

	// Convert to Firebase format
	firebaseItem := FirebaseNewsItem{
		ID:         fmt.Sprintf("%d", newsItem.ID),
		Title:      newsItem.Title,
		Content:    newsItem.Content,
		Source:     string(newsItem.Source),
		Category:   *newsItem.Category,
		Priority:   newsItem.Priority,
		IsBreaking: newsItem.IsBreaking,
		CreatedAt:  newsItem.CreatedAt,
		UpdatedAt:  newsItem.UpdatedAt,
	}

	if newsItem.SourceName != nil {
		firebaseItem.SourceName = *newsItem.SourceName
	}
	if newsItem.SourceURL != nil {
		firebaseItem.SourceURL = *newsItem.SourceURL
	}
	if newsItem.PublishedAt != nil {
		firebaseItem.PublishedAt = *newsItem.PublishedAt
	}
	if newsItem.ExpiresAt != nil {
		firebaseItem.ExpiresAt = *newsItem.ExpiresAt
	}

	// Update in Firestore
	collection := "news_items"
	if fs.client != nil {
		_, err := fs.client.Collection(collection).Doc(firebaseItem.ID).Set(fs.ctx, firebaseItem)
		if err != nil {
			return fmt.Errorf("failed to update news in Firebase: %v", err)
		}

		log.Printf("Successfully updated news item %s in Firebase", firebaseItem.ID)
	}

	return nil
}

// DeleteNewsFromFirebase removes a news item from Firebase
func (fs *FirebaseService) DeleteNewsFromFirebase(newsItemID uint) error {
	if fs.client == nil {
		return fmt.Errorf("Firebase client not initialized")
	}

	collection := "news_items"
	docID := fmt.Sprintf("%d", newsItemID)

	if fs.client != nil {
		_, err := fs.client.Collection(collection).Doc(docID).Delete(fs.ctx)
		if err != nil {
			return fmt.Errorf("failed to delete news from Firebase: %v", err)
		}

		log.Printf("Successfully deleted news item %s from Firebase", docID)
	}

	return nil
}

// GetNewsFromFirebase retrieves news items from Firebase (for external consumption)
func (fs *FirebaseService) GetNewsFromFirebase(limit int, breakingOnly bool) ([]FirebaseNewsItem, error) {
	if fs.client == nil {
		return nil, fmt.Errorf("Firebase client not initialized")
	}

	var newsItems []FirebaseNewsItem
	collection := "news_items"

	query := fs.client.Collection(collection).Where("isActive", "==", true)

	if breakingOnly {
		query = query.Where("isBreaking", "==", true)
	}

	// Order by priority and published date
	query = query.OrderBy("priority", firestore.Desc).
		OrderBy("publishedAt", firestore.Desc).
		Limit(limit)

	docs, err := query.Documents(fs.ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get news from Firebase: %v", err)
	}

	for _, doc := range docs {
		var item FirebaseNewsItem
		if err := doc.DataTo(&item); err != nil {
			log.Printf("Error converting document %s: %v", doc.Ref.ID, err)
			continue
		}
		item.ID = doc.Ref.ID
		newsItems = append(newsItems, item)
	}

	return newsItems, nil
}

// PublishNewsBatch publishes multiple news items to Firebase
func (fs *FirebaseService) PublishNewsBatch(newsItems []*models.NewsItem) error {
	if fs.client == nil {
		return fmt.Errorf("Firebase client not initialized")
	}

	collection := "news_items"
	batch := fs.client.Batch()

	for _, newsItem := range newsItems {
		if !newsItem.IsPublished() {
			continue // Skip unpublished items
		}

		firebaseItem := FirebaseNewsItem{
			ID:         fmt.Sprintf("%d", newsItem.ID),
			Title:      newsItem.Title,
			Content:    newsItem.Content,
			Source:     string(newsItem.Source),
			Category:   *newsItem.Category,
			Priority:   newsItem.Priority,
			IsBreaking: newsItem.IsBreaking,
			CreatedAt:  newsItem.CreatedAt,
			UpdatedAt:  newsItem.UpdatedAt,
		}

		if newsItem.SourceName != nil {
			firebaseItem.SourceName = *newsItem.SourceName
		}
		if newsItem.SourceURL != nil {
			firebaseItem.SourceURL = *newsItem.SourceURL
		}
		if newsItem.PublishedAt != nil {
			firebaseItem.PublishedAt = *newsItem.PublishedAt
		}
		if newsItem.ExpiresAt != nil {
			firebaseItem.ExpiresAt = *newsItem.ExpiresAt
		}

		docRef := fs.client.Collection(collection).Doc(firebaseItem.ID)
		batch.Set(docRef, firebaseItem)
	}

	_, err := batch.Commit(fs.ctx)
	if err != nil {
		return fmt.Errorf("failed to publish news batch to Firebase: %v", err)
	}

	log.Printf("Successfully published batch of %d news items to Firebase", len(newsItems))
	return nil
}

// Close closes the Firebase client
func (fs *FirebaseService) Close() error {
	if fs.client != nil {
		return fs.client.Close()
	}
	return nil
}

// GetPublicNewsEndpoint returns the public Firebase endpoint URL for news
func (fs *FirebaseService) GetPublicNewsEndpoint() string {
	// This would be the public Firebase endpoint that external systems can use
	// In a real implementation, you might use Firebase Functions or a public API
	return "https://your-project-id-default-rtdb.firebaseio.com/news.json"
}

// CreatePublicAPIKey generates or retrieves a public API key for external access
func (fs *FirebaseService) CreatePublicAPIKey() (string, error) {
	// This is a placeholder - in a real implementation, you would:
	// 1. Generate a secure API key
	// 2. Store it in Firebase with appropriate permissions
	// 3. Return the key for external systems to use

	apiKey := "gcx_news_" + fmt.Sprintf("%d", time.Now().Unix())
	log.Printf("Generated public API key: %s", apiKey)

	return apiKey, nil
}
