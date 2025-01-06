package main

import (
	"log"
	"net/http"
	"product-management-system/controllers"
	"product-management-system/models"
)

func main() {
	// Initialize the database
	models.InitDB()

	// Handle routes
	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/products/add", controllers.AddProduct)
	http.HandleFunc("/products/edit", controllers.EditProduct)
	http.HandleFunc("/products/delete", controllers.DeleteProduct)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Start the server
	log.Println("Server is running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
