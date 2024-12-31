package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"project_root/handlers"
	"project_root/middleware"
	"project_root/config"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize the router
	router := mux.NewRouter()

	// Middleware
	router.Use(middleware.Logging)
	router.Use(middleware.Validation)
	router.Use(middleware.Authentication)

	// Blog Routes
	router.HandleFunc("/blog", handlers.CreateBlog).Methods("POST")
	router.HandleFunc("/blog/{id}", handlers.GetBlog).Methods("GET")
	router.HandleFunc("/blogs", handlers.GetBlogs).Methods("GET")
	router.HandleFunc("/blog/{id}", handlers.UpdateBlog).Methods("PUT")
	router.HandleFunc("/blog/{id}", handlers.DeleteBlog).Methods("DELETE")

	// Start the HTTP server
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}
