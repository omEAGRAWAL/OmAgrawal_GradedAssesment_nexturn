package services

import (
	"inventory-management/models"
	"inventory-management/utils"
	"strconv"
)

var inventory []models.Product

// AddProduct adds a new product to the inventory
func AddProduct(id, name, priceStr, stockStr string) error {
	// Convert price to float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil || price <= 0 {
		return utils.NewAppError("Price must be a positive number")
	}

	// Convert stock to int
	stock, err := strconv.Atoi(stockStr)
	if err != nil || stock < 0 {
		return utils.NewAppError("Stock must be a non-negative integer")
	}

	// Check if product ID is unique
	for _, product := range inventory {
		if product.ID == id {
			return utils.NewAppError("Product ID must be unique")
		}
	}

	// Add product to inventory
	product := models.Product{ID: id, Name: name, Price: price, Stock: stock}
	inventory = append(inventory, product)
	return nil
}

// UpdateStock updates the stock of a product
func UpdateStock(id, stockStr string) error {
	stock, err := strconv.Atoi(stockStr)
	if err != nil || stock < 0 {
		return utils.NewAppError("Stock must be a non-negative integer")
	}

	// Find product by ID and update stock
	for i, product := range inventory {
		if product.ID == id {
			inventory[i].Stock = stock
			return nil
		}
	}

	return utils.NewAppError("Product not found")
}

// SearchProduct searches for a product by ID or Name
func SearchProduct(query string) (*models.Product, error) {
	for _, product := range inventory {
		if product.ID == query || product.Name == query {
			return &product, nil
		}
	}

	return nil, utils.NewAppError("Product not found")
}

// DisplayInventory returns the list of all products in the inventory
func DisplayInventory() []models.Product {
	return inventory
}
