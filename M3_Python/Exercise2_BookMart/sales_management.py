from models import Transaction
from book_management import books

sales = []  # List to store transaction records

def sell_book():
    customer_name = input("Customer Name: ")
    customer_email = input("Customer Email: ")
    customer_phone = input("Customer Phone: ")
    book_title = input("Book Title: ")
    quantity = int(input("Quantity: "))

    # Find the book in the inventory
    book = next((book for book in books if book.title.lower() == book_title.lower()), None)
    if not book:
        print("Error: Book not found.")
        return
    if book.quantity < quantity:
        print(f"Error: Only {book.quantity} copies available. Sale cannot be completed.")
        return

    # Process the sale
    book.quantity -= quantity
    sales.append(Transaction(customer_name, customer_email, customer_phone, book.title, quantity))
    print(f"Sale successful! Remaining quantity: {book.quantity}")

def view_sales():
    if not sales:
        print("No sales records available.")
        return
    for sale in sales:
        print(sale)
