package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// BlogPost represents a blog post structure
type BlogPost struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Timestamp time.Time `json:"timestamp"`
}

var db *sql.DB

func main() {
	// Initialize SQLite database
	var err error
	db, err = sql.Open("sqlite3", "./blog.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create table if not exists
	createTable()

	// Set up router
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	// Routes
	r.HandleFunc("/blog", createBlog).Methods("POST")
	r.HandleFunc("/blog/{id:[0-9]+}", getBlogByID).Methods("GET")
	r.HandleFunc("/blogs", getAllBlogs).Methods("GET")
	r.HandleFunc("/blog/{id:[0-9]+}", updateBlogByID).Methods("PUT")
	r.HandleFunc("/blog/{id:[0-9]+}", deleteBlogByID).Methods("DELETE")

	// Start server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Create Table
func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

// Middleware: Logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// Handlers

// Create a new blog
func createBlog(w http.ResponseWriter, r *http.Request) {
	var blog BlogPost
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	query := `INSERT INTO blogs (title, content, author) VALUES (?, ?, ?)`
	result, err := db.Exec(query, blog.Title, blog.Content, blog.Author)
	if err != nil {
		http.Error(w, "Failed to create blog post", http.StatusInternalServerError)
		return
	}
	id, _ := result.LastInsertId()
	blog.ID = int(id)
	blog.Timestamp = time.Now()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

// Get a blog by ID
func getBlogByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var blog BlogPost
	query := `SELECT id, title, content, author, timestamp FROM blogs WHERE id = ?`
	row := db.QueryRow(query, id)
	if err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp); err != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

// Get all blogs
func getAllBlogs(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, title, content, author, timestamp FROM blogs ORDER BY timestamp DESC`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch blogs", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var blogs []BlogPost
	for rows.Next() {
		var blog BlogPost
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp); err != nil {
			http.Error(w, "Failed to read blog", http.StatusInternalServerError)
			return
		}
		blogs = append(blogs, blog)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

// Update a blog by ID
func updateBlogByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var blog BlogPost
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	query := `UPDATE blogs SET title = ?, content = ?, author = ? WHERE id = ?`
	_, err := db.Exec(query, blog.Title, blog.Content, blog.Author, id)
	if err != nil {
		http.Error(w, "Failed to update blog post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Blog post updated successfully"))
}

// Delete a blog by ID
func deleteBlogByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	query := `DELETE FROM blogs WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete blog post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Blog post deleted successfully"))
}
