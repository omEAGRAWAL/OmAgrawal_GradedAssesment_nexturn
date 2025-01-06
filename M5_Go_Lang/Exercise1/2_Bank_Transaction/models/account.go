package models

// Account struct represents a bank account
type Account struct {
	ID             string
	Name           string
	Balance        float64
	TransactionHistory []string
}
