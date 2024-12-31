package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID               int
	Name             string
	Balance          float64
	TransactionHistory []string
}

const (
	DepositOption   = 1
	WithdrawOption  = 2
	ViewBalanceOption = 3
	ViewHistoryOption = 4
	ExitOption      = 5
)

func deposit(account *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	account.Balance += amount
	transaction := fmt.Sprintf("Deposited: $%.2f", amount)
	account.TransactionHistory = append(account.TransactionHistory, transaction)
	return nil
}

func withdraw(account *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	if amount > account.Balance {
		return errors.New("insufficient balance")
	}
	account.Balance -= amount
	transaction := fmt.Sprintf("Withdrew: $%.2f", amount)
	account.TransactionHistory = append(account.TransactionHistory, transaction)
	return nil
}

func viewBalance(account Account) {
	fmt.Printf("Account ID: %d\nName: %s\nBalance: $%.2f\n", account.ID, account.Name, account.Balance)
}

func viewTransactionHistory(account Account) {
	fmt.Println("Transaction History:")
	for i, transaction := range account.TransactionHistory {
		fmt.Printf("%d. %s\n", i+1, transaction)
	}
}

func main() {
	accounts := []Account{
		{ID: 1, Name: "John Doe", Balance: 1000.00},
	}

	var choice int
	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Balance")
		fmt.Println("4. View Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == ExitOption {
			fmt.Println("Exiting system. Goodbye!")
			break
		}

		var accountID int
		fmt.Print("Enter Account ID: ")
		fmt.Scan(&accountID)

		var selectedAccount *Account
		for i := range accounts {
			if accounts[i].ID == accountID {
				selectedAccount = &accounts[i]
				break
			}
		}

		if selectedAccount == nil {
			fmt.Println("Account not found.")
			continue
		}

		switch choice {
		case DepositOption:
			var amount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			if err := deposit(selectedAccount, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful.")
			}

		case WithdrawOption:
			var amount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			if err := withdraw(selectedAccount, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful.")
			}

		case ViewBalanceOption:
			viewBalance(*selectedAccount)

		case ViewHistoryOption:
			viewTransactionHistory(*selectedAccount)

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
