package main

import (
	"errors"
	"fmt"
	"strings"
)

// Employee struct to hold employee details
type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

// Departments as constants
const (
	HR = "HR"
	IT = "IT"
	Finance = "Finance"
)

var employees []Employee // Slice to store employees

// AddEmployee adds a new employee to the list
func AddEmployee(id int, name string, age int, department string) error {
	// Validate ID uniqueness
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}

	// Validate age
	if age <= 18 {
		return errors.New("employee age must be greater than 18")
	}

	// Add employee to the list
	employees = append(employees, Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	})
	return nil
}

// SearchEmployee searches for an employee by ID or name
func SearchEmployee(query string) (Employee, error) {
	for _, emp := range employees {
		if fmt.Sprint(emp.ID) == query || strings.EqualFold(emp.Name, query) {
			return emp, nil
		}
	}
	return Employee{}, errors.New("employee not found")
}

// ListEmployeesByDepartment lists all employees in a given department
func ListEmployeesByDepartment(department string) []Employee {
	var deptEmployees []Employee
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			deptEmployees = append(deptEmployees, emp)
		}
	}
	return deptEmployees
}

// CountEmployeesByDepartment counts the number of employees in a given department
func CountEmployeesByDepartment(department string) int {
	count := 0
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			count++
		}
	}
	return count
}

func main() {
	// Adding employees
	if err := AddEmployee(1, "Alice", 25, HR); err != nil {
		fmt.Println("Error adding employee:", err)
	}
	if err := AddEmployee(2, "Bob", 30, IT); err != nil {
		fmt.Println("Error adding employee:", err)
	}
	if err := AddEmployee(3, "Charlie", 22, Finance); err != nil {
		fmt.Println("Error adding employee:", err)
	}

	// Attempting to add an invalid employee
	if err := AddEmployee(2, "Duplicate", 20, HR); err != nil {
		fmt.Println("Error adding employee:", err)
	}

	// Searching for employees
	emp, err := SearchEmployee("Alice")
	if err != nil {
		fmt.Println("Error searching employee:", err)
	} else {
		fmt.Println("Found employee:", emp)
	}

	emp, err = SearchEmployee("10")
	if err != nil {
		fmt.Println("Error searching employee:", err)
	} else {
		fmt.Println("Found employee:", emp)
	}

	// Listing employees by department
	fmt.Println("Employees in HR:", ListEmployeesByDepartment(HR))
	fmt.Println("Employees in IT:", ListEmployeesByDepartment(IT))

	// Counting employees by department
	fmt.Println("Number of employees in Finance:", CountEmployeesByDepartment(Finance))
	fmt.Println("Number of employees in HR:", CountEmployeesByDepartment(HR))
}
