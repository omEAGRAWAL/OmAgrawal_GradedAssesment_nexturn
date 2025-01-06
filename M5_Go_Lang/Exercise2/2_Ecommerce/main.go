package main

import (
	"ecommerce-inventory/config"
	"ecommerce-inventory/handlers"
	"ecommerce-inventory/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	db := config.InitDB()
	defer db.Close()

	// Router setup
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware, middleware.RateLimitMiddleware)

	// Product routes
	r.HandleFunc("/product", handlers.AddProduct(db)).Methods("POST")
	r.HandleFunc("/product/{id:[0-9]+}", handlers.GetProduct(db)).Methods("GET")
	r.HandleFunc("/product/{id:[0-9]+}", handlers.UpdateProduct(db)).Methods("PUT")
	r.HandleFunc("/product/{id:[0-9]+}", handlers.DeleteProduct(db)).Methods("DELETE")
	r.HandleFunc("/products", handlers.GetProducts(db)).Methods("GET")

	// Start server
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
