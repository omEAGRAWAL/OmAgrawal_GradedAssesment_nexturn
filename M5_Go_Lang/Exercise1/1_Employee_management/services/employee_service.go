package services

import (
	"employee-management/models"
	"employee-management/utils"
	"strconv"
)

var employees []models.Employee

// AddEmployee adds a new employee
func AddEmployee(id, name, ageStr, department string) error {
	for _, emp := range employees {
		if emp.ID == id {
			return utils.NewAppError("ID must be unique")
		}
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil || age <= 18 {
		return utils.NewAppError("Age must be greater than 18")
	}

	if department != "HR" && department != "IT" {
		return utils.NewAppError("Department must be HR or IT")
	}

	employee := models.Employee{ID: id, Name: name, Age: age, Department: department}
	employees = append(employees, employee)
	return nil
}

// SearchEmployee searches for an employee by ID or Name
func SearchEmployee(query string) (*models.Employee, error) {
	for _, emp := range employees {
		if emp.ID == query || emp.Name == query {
			return &emp, nil
		}
	}
	return nil, utils.NewAppError("Employee not found")
}

// ListEmployeesByDepartment lists employees by department
func ListEmployeesByDepartment(department string) []models.Employee {
	var result []models.Employee
	for _, emp := range employees {
		if emp.Department == department {
			result = append(result, emp)
		}
	}
	return result
}

// CountEmployees counts employees by department
func CountEmployees() map[string]int {
	counts := map[string]int{"HR": 0, "IT": 0}
	for _, emp := range employees {
		counts[emp.Department]++
	}
	return counts
}
