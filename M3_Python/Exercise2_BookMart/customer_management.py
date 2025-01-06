from models import Customer

customers = []  # List to store customer objects

def add_customer():
    name = input("Name: ")
    email = input("Email: ")
    phone = input("Phone: ")
    customers.append(Customer(name, email, phone))
    print("Customer added successfully!")

def view_customers():
    if not customers:
        print("No customers available.")
        return
    for customer in customers:
        print(customer)
