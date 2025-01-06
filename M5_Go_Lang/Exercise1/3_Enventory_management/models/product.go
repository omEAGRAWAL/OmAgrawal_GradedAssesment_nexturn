package models

// Product represents an item in the inventory
type Product struct {
	ID    string
	Name  string
	Price float64
	Stock int
}
