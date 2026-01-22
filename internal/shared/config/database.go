package config

import (
	"fmt"
	"log"
	"os"
	"time"

	cms_models "gcx-cms/internal/cms/models"
	marketdata_models "gcx-cms/internal/marketdata/models"
	"gcx-cms/internal/models"
	shared_models "gcx-cms/internal/shared/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	var err error
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "sqlite"
	}

	var dsn string
	var dialector gorm.Dialector

	switch dbType {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"))
		dialector = mysql.Open(dsn)
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"))
		dialector = postgres.Open(dsn)
	default: // sqlite
		dbName := os.Getenv("DB_NAME")
		if dbName == "" {
			dbName = "gcx_cms.db"
		}
		dsn = dbName
		dialector = sqlite.Open(dsn)
	}

	// Configure GORM logger
	gormLogger := logger.Default
	if os.Getenv("GIN_MODE") == "release" {
		gormLogger = logger.Default.LogMode(logger.Silent)
	}

	DB, err = gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get underlying sql.DB:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Printf("✅ Connected to %s database", dbType)

	// Check if database is already migrated (faster than AutoMigrate)
	if !isDatabaseMigrated() {
		log.Println("Database not migrated, running AutoMigrate...")
		if err := AutoMigrate(); err != nil {
			log.Fatal("Failed to migrate database:", err)
		}
	} else {
		log.Println("Database already migrated, skipping AutoMigrate")
	}

	// Create default admin user
	CreateDefaultAdmin()
}

// isDatabaseMigrated checks if key tables exist (faster than AutoMigrate)
func isDatabaseMigrated() bool {
	// Check if a few key tables exist
	var count int64
	DB.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name IN ('users', 'settings', 'pages')").Scan(&count)
	return count >= 3
}

// AutoMigrate runs database migrations
func AutoMigrate() error {
	return DB.AutoMigrate(
		// Shared models
		&shared_models.User{},

		// CMS models
		&cms_models.BlogPost{},
		&cms_models.BlogCategory{},
		&cms_models.MediaFile{},
		&cms_models.Page{},
		&cms_models.Setting{},
		&cms_models.Menu{},
		&cms_models.MenuItem{},
		&cms_models.Translation{},
		&cms_models.Language{},
		&cms_models.BoardMember{},
		&cms_models.TeamMember{},
		&cms_models.Trader{},
		&cms_models.Broker{},
		&models.Partner{},
		&cms_models.Publication{},
		&cms_models.Career{},
		&cms_models.Commodity{},

		// Market Data models
		&marketdata_models.MarketData{},
		&marketdata_models.CommodityInfo{},
		&marketdata_models.TradingSession{},
		&marketdata_models.PriceAlert{},
		&marketdata_models.MarketAnalytics{},

		// Subscription models
		&marketdata_models.SubscriptionPlan{},
		&marketdata_models.UserSubscription{},
		&marketdata_models.SubscriptionFeature{},
		&marketdata_models.UserDataAccess{},
	)
}

// CreateDefaultAdmin creates a default admin user if none exists
func CreateDefaultAdmin() {
	var count int64
	DB.Model(&shared_models.User{}).Where("role = ?", shared_models.RoleAdmin).Count(&count)

	if count == 0 {
		adminEmail := os.Getenv("ADMIN_EMAIL")
		adminPassword := os.Getenv("ADMIN_PASSWORD")
		adminName := os.Getenv("ADMIN_NAME")

		if adminEmail == "" {
			adminEmail = "admin@gcx.com"
		}
		if adminPassword == "" {
			adminPassword = "admin123"
		}
		if adminName == "" {
			adminName = "GCX Admin"
		}

		admin := shared_models.User{
			Name:     adminName,
			Email:    adminEmail,
			Password: adminPassword,
			Role:     shared_models.RoleAdmin,
			IsActive: true,
		}

		if err := DB.Create(&admin).Error; err != nil {
			log.Printf("Failed to create default admin user: %v", err)
		} else {
			log.Printf("✅ Created default admin user: %s", adminEmail)
		}
	}
}
