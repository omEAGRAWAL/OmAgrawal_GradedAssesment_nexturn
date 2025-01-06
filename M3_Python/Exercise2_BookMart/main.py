import book_management
import customer_management
import sales_management

def main_menu():
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("Enter your choice: ")

        if choice == '1':
            book_management_menu()
        elif choice == '2':
            customer_management_menu()
        elif choice == '3':
            sales_management_menu()
        elif choice == '4':
            print("Exiting the application. Goodbye!")
            break
        else:
            print("Invalid choice! Please try again.")

def book_management_menu():
    while True:
        print("\nBook Management:")
        print("1. Add Book")
        print("2. View All Books")
        print("3. Search Book")
        print("4. Back to Main Menu")
        choice = input("Enter your choice: ")

        if choice == '1':
            book_management.add_book()
        elif choice == '2':
            book_management.view_books()
        elif choice == '3':
            book_management.search_book()
        elif choice == '4':
            break
        else:
            print("Invalid choice! Please try again.")

def customer_management_menu():
    while True:
        print("\nCustomer Management:")
        print("1. Add Customer")
        print("2. View All Customers")
        print("3. Back to Main Menu")
        choice = input("Enter your choice: ")

        if choice == '1':
            customer_management.add_customer()
        elif choice == '2':
            customer_management.view_customers()
        elif choice == '3':
            break
        else:
            print("Invalid choice! Please try again.")

def sales_management_menu():
    while True:
        print("\nSales Management:")
        print("1. Sell Book")
        print("2. View All Sales")
        print("3. Back to Main Menu")
        choice = input("Enter your choice: ")

        if choice == '1':
            sales_management.sell_book()
        elif choice == '2':
            sales_management.view_sales()
        elif choice == '3':
            break
        else:
            print("Invalid choice! Please try again.")

if __name__ == "__main__":
    main_menu()
