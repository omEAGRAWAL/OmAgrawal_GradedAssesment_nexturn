from models import Book

books = []  # List to store book objects

def add_book():
    try:
        title = input("Title: ")
        author = input("Author: ")
        price = float(input("Price: "))
        if price <= 0:
            raise ValueError("Price must be a positive number.")
        quantity = int(input("Quantity: "))
        if quantity <= 0:
            raise ValueError("Quantity must be a positive number.")
        books.append(Book(title, author, price, quantity))
        print("Book added successfully!")
    except ValueError as e:
        print(f"Invalid input! {e}")

def view_books():
    if not books:
        print("No books available.")
        return
    for book in books:
        print(book)

def search_book():
    query = input("Enter title or author to search: ").lower()
    found_books = [book for book in books if query in book.title.lower() or query in book.author.lower()]
    if not found_books:
        print("No books found.")
    else:
        for book in found_books:
            print(book)
