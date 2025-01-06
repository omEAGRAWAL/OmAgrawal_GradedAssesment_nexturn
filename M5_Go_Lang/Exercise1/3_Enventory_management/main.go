package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"inventory-management/services"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nInventory Management System")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			addProduct(scanner)
		case "2":
			updateStock(scanner)
		case "3":
			searchProduct(scanner)
		case "4":
			displayInventory()
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addProduct(scanner *bufio.Scanner) {
	fmt.Print("Enter Product ID: ")
	scanner.Scan()
	id := scanner.Text()

	fmt.Print("Enter Product Name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter Product Price: ")
	scanner.Scan()
	price := scanner.Text()

	fmt.Print("Enter Product Stock: ")
	scanner.Scan()
	stock := scanner.Text()

	err := services.AddProduct(id, name, price, stock)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Product added successfully!")
	}
}

func updateStock(scanner *bufio.Scanner) {
	fmt.Print("Enter Product ID: ")
	scanner.Scan()
	id := scanner.Text()

	fmt.Print("Enter Stock Quantity to Update: ")
	scanner.Scan()
	stock := scanner.Text()

	err := services.UpdateStock(id, stock)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Stock updated successfully!")
	}
}

func searchProduct(scanner *bufio.Scanner) {
	fmt.Print("Search by ID or Name: ")
	scanner.Scan()
	query := scanner.Text()

	product, err := services.SearchProduct(query)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Product Found: %+v\n", product)
	}
}

func displayInventory() {
	fmt.Println("Available Products:")
	products := services.DisplayInventory()
	if len(products) == 0 {
		fmt.Println("No products in inventory.")
	} else {
		fmt.Printf("%-10s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
		fmt.Println(strings.Repeat("-", 50))
		for _, product := range products {
			fmt.Printf("%-10s %-20s %-10.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
		}
	}
}
