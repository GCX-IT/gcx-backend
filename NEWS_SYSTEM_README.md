# GCX News Ticker System

This document describes the GCX News Ticker System that provides real-time news updates for the website and enables external channels to post news through various APIs.

## 🎯 Overview

The news ticker system consists of:
- **Go Backend**: Manages news items, categories, and API endpoints
- **Firebase Integration**: Provides public access to news data
- **Frontend Component**: Displays news in a TV-style ticker format
- **Multi-channel Support**: Allows posting from different sources

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   External      │    │   GCX Backend   │    │   Firebase      │
│   Channels      │───▶│   (Go API)      │───▶│   (Public)      │
│                 │    │                 │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         │                       ▼                       │
         │              ┌─────────────────┐              │
         │              │   Database      │              │
         │              │   (MySQL)       │              │
         │              └─────────────────┘              │
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │   CMS Panel     │    │   External      │
│   News Ticker   │    │   (Admin)       │    │   Consumers     │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 📊 Database Schema

### news_items Table
```sql
CREATE TABLE news_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    content TEXT,
    source ENUM('gcx', 'partner', 'external', 'api', 'firebase') NOT NULL,
    source_name VARCHAR(255),
    source_url VARCHAR(500),
    category VARCHAR(100),
    priority INT DEFAULT 0,
    status ENUM('draft', 'published', 'archived') DEFAULT 'draft',
    is_breaking BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    published_at TIMESTAMP NULL,
    expires_at TIMESTAMP NULL,
    external_id VARCHAR(255),
    external_data JSON,
    last_sync_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
```

### news_categories Table
```sql
CREATE TABLE news_categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    color VARCHAR(7),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🚀 API Endpoints

### Public Endpoints (No Authentication Required)

#### Get Active News Items
```http
GET /api/news?limit=20&source=gcx&category=market&breaking=false
```

#### Get Breaking News Only
```http
GET /api/news/breaking?limit=10
```

#### Get Single News Item
```http
GET /api/news/{id}
```

#### Get News Categories
```http
GET /api/news-categories
```

### Protected Endpoints (Authentication Required)

#### Create News Item
```http
POST /api/cms/news
Content-Type: application/json

{
    "title": "News Title",
    "content": "News content...",
    "source": "gcx",
    "source_name": "GCX Official",
    "category": "announcement",
    "priority": 5,
    "is_breaking": false,
    "status": "published"
}
```

#### Update News Item
```http
PUT /api/cms/news/{id}
Content-Type: application/json

{
    "title": "Updated Title",
    "is_breaking": true
}
```

#### Publish News Item
```http
POST /api/cms/news/{id}/publish
```

#### Archive News Item
```http
POST /api/cms/news/{id}/archive
```

#### Set Breaking News
```http
PUT /api/cms/news/{id}/breaking
Content-Type: application/json

{
    "is_breaking": true
}
```

## 🔥 Firebase Integration

### Public Firebase Endpoint
```
https://your-project-id-default-rtdb.firebaseio.com/news.json
```

### Firebase Data Structure
```json
{
  "news_items": {
    "1": {
      "id": "1",
      "title": "News Title",
      "content": "News content...",
      "source": "gcx",
      "sourceName": "GCX Official",
      "sourceUrl": "https://example.com",
      "category": "announcement",
      "priority": 5,
      "isBreaking": false,
      "publishedAt": "2024-01-01T12:00:00Z",
      "expiresAt": "2024-01-02T12:00:00Z",
      "createdAt": "2024-01-01T12:00:00Z",
      "updatedAt": "2024-01-01T12:00:00Z"
    }
  }
}
```

## 📱 Frontend Integration

### News Ticker Component
The news ticker is automatically integrated into the hero section and displays:
- Breaking news with special highlighting
- Category badges with color coding
- Source information and timestamps
- Automatic rotation with pause on hover
- Responsive design for all screen sizes

### Usage in Vue Components
```vue
<template>
  <div>
    <NewsTicker />
  </div>
</template>

<script setup>
import NewsTicker from '@/components/Landing/NewsTicker.vue'
</script>
```

### Using the News Composable
```vue
<script setup>
import { useNews, useNewsTicker } from '@/composables/useNews'

const { newsItems, isLoading, error } = useNews()
const { currentNewsItem, nextNewsItem, pauseTicker } = useNewsTicker()
</script>
```

## 🔌 External Channel Integration

### 1. Direct API Integration

#### Post News via REST API
```javascript
const postNews = async (newsItem) => {
  const response = await fetch('https://gcx-api.com/api/cms/news', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer YOUR_API_TOKEN'
    },
    body: JSON.stringify({
      title: newsItem.title,
      content: newsItem.content,
      source: 'partner',
      source_name: 'Your Organization',
      category: 'market',
      priority: 5,
      is_breaking: false,
      status: 'published'
    })
  });
  
  return response.json();
};
```

### 2. Firebase Integration

#### Read News from Firebase
```javascript
import { initializeApp } from 'firebase/app';
import { getDatabase, ref, onValue } from 'firebase/database';

const firebaseConfig = {
  // Your Firebase config
};

const app = initializeApp(firebaseConfig);
const database = getDatabase(app);

// Listen for news updates
const newsRef = ref(database, 'news_items');
onValue(newsRef, (snapshot) => {
  const newsData = snapshot.val();
  console.log('News updated:', newsData);
});
```

### 3. Webhook Integration

#### Set up webhook endpoint
```go
func webhookHandler(c *gin.Context) {
    var webhookData map[string]interface{}
    if err := c.ShouldBindJSON(&webhookData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // Process webhook data and create news item
    newsItem := &models.NewsItem{
        Title: webhookData["title"].(string),
        Content: webhookData["content"].(string),
        Source: models.NewsSourceExternal,
        SourceName: webhookData["source_name"].(string),
        IsBreaking: webhookData["is_breaking"].(bool),
        Status: models.NewsStatusPublished,
    }
    
    // Save to database and sync to Firebase
    newsService.CreateNewsItem(newsItem, true)
}
```

## 🎨 News Sources

### Supported Sources
- **gcx**: Internal GCX news
- **partner**: Partner organization news
- **external**: External news sources
- **api**: API-generated news
- **firebase**: Firebase-synced news

### News Categories
- **market**: Market updates and price information
- **announcement**: Official announcements
- **event**: Events and programs
- **partnership**: Partnership news
- **regulation**: Regulatory updates
- **technology**: Technology and innovation

## 🔧 Configuration

### Environment Variables
```bash
# Firebase Configuration
FIREBASE_PROJECT_ID=your-project-id
FIREBASE_CREDENTIALS=base64-encoded-service-account-json
FIREBASE_COLLECTION=news_items

# News System Settings
NEWS_AUTO_SYNC_INTERVAL=300  # seconds
NEWS_BREAKING_PRIORITY=10
NEWS_DEFAULT_EXPIRY=86400    # seconds (24 hours)
```

### Database Migration
```bash
# Run the news tables migration
mysql -u username -p database_name < database/migrations/026_create_news_tables.sql
```

## 🧪 Testing

### Test API Integration
```bash
# Run the test script
cd gcx-go-backend
go run test_news_api.go
```

### Test Firebase Integration
```javascript
// Test Firebase connection
import { getFirebaseService } from './services/firebaseService.js';

const firebase = getFirebaseService();
const newsItems = await firebase.GetNewsFromFirebase(10, false);
console.log('Firebase news items:', newsItems);
```

## 📈 Performance Considerations

### Optimization Features
- **Caching**: News items are cached for 5 minutes
- **Pagination**: Large news lists are paginated
- **Lazy Loading**: News ticker loads items on demand
- **Auto-refresh**: Background updates every 5 minutes
- **Breaking News**: Immediate updates for breaking news

### Monitoring
- **API Response Times**: Monitor endpoint performance
- **Firebase Sync**: Track sync success/failure rates
- **News Engagement**: Monitor click-through rates
- **Error Logging**: Comprehensive error tracking

## 🔒 Security

### Authentication
- **API Keys**: Required for protected endpoints
- **Rate Limiting**: Prevents API abuse
- **Input Validation**: Sanitizes all input data
- **CORS**: Configured for specific domains

### Data Privacy
- **Source Attribution**: All news items include source information
- **Expiry Dates**: Automatic cleanup of expired news
- **Audit Trail**: Complete history of news item changes

## 🚀 Deployment

### Production Setup
1. **Database**: Run migrations on production database
2. **Firebase**: Configure production Firebase project
3. **Environment**: Set production environment variables
4. **API Keys**: Configure API authentication
5. **Monitoring**: Set up logging and monitoring

### Scaling Considerations
- **Database Indexing**: Optimize queries with proper indexes
- **Firebase Limits**: Monitor Firebase usage and limits
- **CDN**: Use CDN for static news assets
- **Load Balancing**: Distribute API load across servers

## 📞 Support

For technical support or questions about the news system:
- **Documentation**: Check this README and API docs
- **Issues**: Report bugs via GitHub issues
- **Contact**: Reach out to the development team

---

## 🎉 Conclusion

The GCX News Ticker System provides a comprehensive solution for managing and displaying news across multiple channels. With its flexible API, Firebase integration, and modern frontend components, it enables seamless news distribution and consumption.

Key benefits:
- ✅ Multi-channel news posting
- ✅ Real-time updates via Firebase
- ✅ Professional TV-style ticker display
- ✅ Comprehensive API for external integration
- ✅ Automatic news management and expiry
- ✅ Breaking news system
- ✅ Responsive design for all devices
