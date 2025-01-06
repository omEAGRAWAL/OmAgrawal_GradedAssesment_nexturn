package services

import (
	"bank-transaction-system/models"
	"bank-transaction-system/utils"
	"strconv"
)

var accounts []models.Account

// CreateAccount creates a new account
func CreateAccount(id, name string) error {
	for _, account := range accounts {
		if account.ID == id {
			return utils.NewAppError("Account ID must be unique")
		}
	}

	account := models.Account{
		ID:                 id,
		Name:               name,
		Balance:            0,
		TransactionHistory: []string{},
	}
	accounts = append(accounts, account)
	return nil
}

// Deposit adds money to an account
func Deposit(id, amountStr string) error {
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		return utils.NewAppError("Deposit amount must be greater than zero")
	}

	for i, account := range accounts {
		if account.ID == id {
			accounts[i].Balance += amount
			accounts[i].TransactionHistory = append(accounts[i].TransactionHistory,
				"Deposited: "+amountStr)
			return nil
		}
	}

	return utils.NewAppError("Account not found")
}

// Withdraw subtracts money from an account
func Withdraw(id, amountStr string) error {
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		return utils.NewAppError("Withdrawal amount must be greater than zero")
	}

	for i, account := range accounts {
		if account.ID == id {
			if account.Balance < amount {
				return utils.NewAppError("Insufficient balance")
			}
			accounts[i].Balance -= amount
			accounts[i].TransactionHistory = append(accounts[i].TransactionHistory,
				"Withdrew: "+amountStr)
			return nil
		}
	}

	return utils.NewAppError("Account not found")
}

// GetBalance retrieves the balance of an account
func GetBalance(id string) (float64, error) {
	for _, account := range accounts {
		if account.ID == id {
			return account.Balance, nil
		}
	}

	return 0, utils.NewAppError("Account not found")
}

// GetTransactionHistory retrieves the transaction history of an account
func GetTransactionHistory(id string) ([]string, error) {
	for _, account := range accounts {
		if account.ID == id {
			return account.TransactionHistory, nil
		}
	}

	return nil, utils.NewAppError("Account not found")
}
