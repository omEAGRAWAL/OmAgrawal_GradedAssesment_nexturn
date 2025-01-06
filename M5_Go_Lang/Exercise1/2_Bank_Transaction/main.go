package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"bank-transaction-system/services"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. View Balance")
		fmt.Println("5. View Transaction History")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			createAccount(scanner)
		case "2":
			performDeposit(scanner)
		case "3":
			performWithdraw(scanner)
		case "4":
			viewBalance(scanner)
		case "5":
			viewTransactionHistory(scanner)
		case "6":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func createAccount(scanner *bufio.Scanner) {
	fmt.Print("Enter Account ID: ")
	scanner.Scan()
	id := scanner.Text()

	fmt.Print("Enter Account Holder Name: ")
	scanner.Scan()
	name := scanner.Text()

	err := services.CreateAccount(id, name)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Account created successfully!")
	}
}

func performDeposit(scanner *bufio.Scanner) {
	fmt.Print("Enter Account ID: ")
	scanner.Scan()
	id := scanner.Text()

	fmt.Print("Enter Deposit Amount: ")
	scanner.Scan()
	amount := scanner.Text()

	err := services.Deposit(id, amount)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Deposit successful!")
	}
}

func performWithdraw(scanner *bufio.Scanner) {
	fmt.Print("Enter Account ID: ")
	scanner.Scan()
	id := scanner.Text()

	fmt.Print("Enter Withdrawal Amount: ")
	scanner.Scan()
	amount := scanner.Text()

	err := services.Withdraw(id, amount)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Withdrawal successful!")
	}
}

func viewBalance(scanner *bufio.Scanner) {
	fmt.Print("Enter Account ID: ")
	scanner.Scan()
	id := scanner.Text()

	balance, err := services.GetBalance(id)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Current Balance: %.2f\n", balance)
	}
}

func viewTransactionHistory(scanner *bufio.Scanner) {
	fmt.Print("Enter Account ID: ")
	scanner.Scan()
	id := scanner.Text()

	history, err := services.GetTransactionHistory(id)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Transaction History:")
		for _, entry := range history {
			fmt.Println("- " + entry)
		}
	}
}
