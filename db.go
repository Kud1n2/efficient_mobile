package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// initDB godoc
// @Summary Initialize database connection
// @Description Connects to PostgreSQL database using environment variables
// @Tags Database
// @Success 200 "Database connection established"
// @Failure 500 {object} ErrorResponse "Database connection failed"
func initDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("DB connect error")
		return nil, err
	}
	if err := DB.Ping(); err != nil {
		return nil, err
	}
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS subscriptions (
		id SERIAL PRIMARY KEY,
		service_name TEXT NOT NULL,
		price INTEGER NOT NULL,
		user_id TEXT NOT NULL,
		start_date TEXT NOT NULL,
		finish_date TEXT
	);`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	log.Println("Database connection established and table verified/created")
	return DB, nil
}
