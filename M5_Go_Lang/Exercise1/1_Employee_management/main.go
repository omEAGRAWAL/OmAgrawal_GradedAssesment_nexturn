package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"employee-management/services"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nEmployee Management System")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. List Employees by Department")
		fmt.Println("4. Count Employees by Department")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			addEmployee(scanner)
		case "2":
			searchEmployee(scanner)
		case "3":
			listEmployeesByDepartment(scanner)
		case "4":
			countEmployeesByDepartment()
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addEmployee(scanner *bufio.Scanner) {
	fmt.Print("Enter ID: ")
	scanner.Scan()
	id := scanner.Text()

	fmt.Print("Enter Name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter Age: ")
	scanner.Scan()
	age := scanner.Text()

	fmt.Print("Enter Department (HR/IT): ")
	scanner.Scan()
	department := scanner.Text()

	err := services.AddEmployee(id, name, age, department)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee added successfully!")
	}
}

func searchEmployee(scanner *bufio.Scanner) {
	fmt.Print("Search by ID or Name: ")
	scanner.Scan()
	query := scanner.Text()

	employee, err := services.SearchEmployee(query)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Employee Found: %+v\n", employee)
	}
}

func listEmployeesByDepartment(scanner *bufio.Scanner) {
	fmt.Print("Enter Department (HR/IT): ")
	scanner.Scan()
	department := scanner.Text()

	employees := services.ListEmployeesByDepartment(department)
	if len(employees) == 0 {
		fmt.Println("No employees found in this department.")
	} else {
		fmt.Println("Employees:")
		for _, emp := range employees {
			fmt.Printf("- %+v\n", emp)
		}
	}
}

func countEmployeesByDepartment() {
	counts := services.CountEmployees()
	fmt.Println("Employee Counts by Department:")
	for department, count := range counts {
		fmt.Printf("%s: %d\n", department, count)
	}
}
