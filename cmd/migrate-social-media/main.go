package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database connection
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "gcx_cms")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Connected to database successfully!")

	// Read and execute the migration
	migrationSQL := `
		ALTER TABLE team_members 
		ADD COLUMN linkedin_url VARCHAR(500),
		ADD COLUMN twitter_url VARCHAR(500),
		ADD COLUMN facebook_url VARCHAR(500),
		ADD COLUMN instagram_url VARCHAR(500),
		ADD COLUMN email VARCHAR(255),
		ADD COLUMN phone VARCHAR(50),
		ADD COLUMN website_url VARCHAR(500);
	`

	_, err = db.Exec(migrationSQL)
	if err != nil {
		log.Fatalf("Failed to execute migration: %v", err)
	}

	fmt.Println("✅ Social media columns added to team_members table successfully!")
	fmt.Println("Added columns:")
	fmt.Println("  - linkedin_url")
	fmt.Println("  - twitter_url")
	fmt.Println("  - facebook_url")
	fmt.Println("  - instagram_url")
	fmt.Println("  - email")
	fmt.Println("  - phone")
	fmt.Println("  - website_url")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
