package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Product struct
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Category    string
}

// Initialize the database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "products.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create the products table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		price REAL,
		category TEXT
	);`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	log.Println("Database initialized successfully")
}
