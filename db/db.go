package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading Env File %v", err)
	}
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	fmt.Println("✅ Connected to database successfully!")
	return db
}