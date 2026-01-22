package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database connection
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "root:password@tcp(localhost:3306)/gcx_cms?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Connected to database successfully")

	// Read and execute migration file
	migrationFile := "database/migrations/2025_01_15_create_careers_table.sql"
	sqlBytes, err := ioutil.ReadFile(migrationFile)
	if err != nil {
		log.Fatalf("Failed to read migration file: %v", err)
	}

	sqlContent := string(sqlBytes)

	// Execute the migration
	if _, err := db.Exec(sqlContent); err != nil {
		log.Fatalf("Failed to execute migration: %v", err)
	}

	fmt.Println("Careers table migration completed successfully!")
	fmt.Println("Sample careers data has been inserted.")
}
