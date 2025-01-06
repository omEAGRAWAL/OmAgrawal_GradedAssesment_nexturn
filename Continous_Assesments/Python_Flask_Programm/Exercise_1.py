# Exercise 1: Invoice Generator
# Program to take user inputs and generate an invoice

def generate_invoice():
    user_name = input("Enter your name: ")
    num_products = int(input("Enter the number of products bought: "))
    products = []

    for i in range(num_products):
        product_name = input(f"Enter the name of product {i + 1}: ")
        product_price = float(input(f"Enter the price of {product_name}: "))
        products.append((product_name, product_price))

    total_cost = sum(price for _, price in products)

    print(f"\nInvoice for {user_name}")
    print("=" * 20)
    for product, price in products:
        print(f"{product}: ${price:.2f}")
    print("=" * 20)
    print(f"Total Cost: ${total_cost:.2f}")

# Exercise 2: Temperature Converter with Units
# Program to convert temperature between Celsius and Fahrenheit

def temperature_converter():
    temperature = float(input("Enter the temperature: "))
    unit = input("Enter the unit (C/F): ").strip().upper()

    if unit == "C":
        converted = (temperature * 9/5) + 32
        print(f"{temperature}C is {converted:.2f}F")
    elif unit == "F":
        converted = (temperature - 32) * 5/9
        print(f"{temperature}F is {converted:.2f}C")
    else:
        print("Invalid unit. Please enter 'C' or 'F'.")

# Exercise 3: Employee Age Validator
# Program to validate integer input for employee age

def validate_age():
    while True:
        try:
            age = int(input("Enter the employee's age: "))
            print(f"Age {age} recorded.")
            break
        except ValueError:
            print("Invalid input. Please enter a valid integer.")

# Exercise 4: Data Type Validator
# Program to validate inputs for name, salary, and dependents

def validate_inputs():
    while True:
        name = input("Enter your name: ")
        if not name.isalpha():
            print("Invalid name. Enter alphabetic characters only.")
            continue

        try:
            salary = float(input("Enter your salary: "))
            dependents = int(input("Enter number of dependents: "))
            print("Inputs recorded successfully.")
            break
        except ValueError:
            print("Invalid input for salary or dependents. Please try again.")

# Exercise 5: Personal Budget Tracker
# Program to calculate budget surplus and expense percentages

def budget_tracker():
    salary = float(input("Enter your monthly salary: "))
    rent = float(input("Enter your rent expense: "))
    groceries = float(input("Enter your groceries expense: "))
    utilities = float(input("Enter your utilities expense: "))

    total_expenses = rent + groceries + utilities
    surplus = salary - total_expenses

    print(f"Surplus Money: ${surplus:.2f}")
    print(f"Rent: {rent / salary * 100:.2f}%")
    print(f"Groceries: {groceries / salary * 100:.2f}%")
    print(f"Utilities: {utilities / salary * 100:.2f}%")

# Exercise 6: Employee Salary Breakdown
# Program to calculate net salary

def salary_breakdown():
    basic_salary = float(input("Enter your basic salary: "))
    tax_rate = float(input("Enter your tax rate (%): "))
    bonus = float(input("Enter your bonus: "))

    tax = (tax_rate / 100) * basic_salary
    net_salary = basic_salary - tax + bonus

    print(f"Total Deductions: ${tax:.2f}")
    print(f"Net Salary: ${net_salary:.2f}")

# Exercise 7: Movie Ticket Price Calculator
# Program to calculate movie ticket price based on age

def ticket_price_calculator():
    age = int(input("Enter your age: "))

    if age < 5:
        price = 0
    elif 5 <= age <= 12:
        price = 5
    elif 13 <= age <= 60:
        price = 12
    else:
        price = 8

    print(f"Your ticket price is: ${price}")

# Exercise 8: Admission Checker
# Program to check admission eligibility

def admission_checker():
    age = int(input("Enter your age: "))
    score = float(input("Enter your entrance exam score: "))

    if 18 <= age <= 25 and score >= 70:
        print("Admission Granted")
    else:
        print("Admission Denied")

# Exercise 9: Product Stock Manager
# Program to manage product inventory

def stock_manager():
    inventory = {}

    while True:
        product = input("Enter product name (or 'done' to finish): ").strip()
        if product.lower() == 'done':
            break
        quantity = int(input(f"Enter the stock quantity for {product}: "))
        inventory[product] = quantity

    print("\nInventory List:")
    for product, quantity in inventory.items():
        print(f"{product}: {quantity}")

# Exercise 10: Even Numbers in a Range
# Program to print even numbers in a range

def even_numbers():
    start = int(input("Enter the start of the range: "))
    end = int(input("Enter the end of the range: "))

    print("Even numbers in the range:")
    for num in range(start, end + 1):
        if num % 2 == 0:
            print(num, end=" ")
    print()

# Exercise 11: Password Retry with Break
# Program to validate password with retry attempts

def password_retry():
    correct_password = "secret"
    attempts = 3

    while attempts > 0:
        entered_password = input("Enter the password: ")
        if entered_password == correct_password:
            print("Access Granted")
            break
        attempts -= 1
        print(f"Incorrect password. Attempts remaining: {attempts}")

    if attempts == 0:
        print("Access Denied")

# Exercise 12: Prime Number Skipping
# Program to print prime numbers between 1 and 50

def prime_numbers():
    print("Prime numbers between 1 and 50:")
    for num in range(2, 51):
        for i in range(2, int(num ** 0.5) + 1):
            if num % i == 0:
                break
        else:
            print(num, end=" ")
    print()

# Exercise 13: Student Marks Analysis
# Program to analyze marks

def marks_analysis():
    marks = list(map(int, input("Enter student marks separated by spaces: ").split()))
    print(f"Highest Mark: {max(marks)}")
    print(f"Lowest Mark: {min(marks)}")
    print(f"Average Mark: {sum(marks) / len(marks):.2f}")

# Exercise 14: Sorting Student Names
# Program to sort student names

def sort_names():
    names = input("Enter student names separated by commas: ").split(',')
    names = [name.strip() for name in names]
    names.sort()
    longest_name = max(names, key=len)

    print("Sorted Names:", names)
    print("Longest Name:", longest_name)

# Exercise 15: Grading System
# Program to calculate grades

def calculate_grade(marks):
    if marks >= 90:
        return 'A'
    elif marks >= 75:
        return 'B'
    elif marks >= 50:
        return 'C'
    else:
        return 'F'

# Exercise 16: Compound Interest Calculator
# Program to calculate compound interest

def calculate_compound_interest(principal, rate, time):
    return principal * (1 + rate / 100) ** time

# Exercise 17: Sorting Students by Marks
# Program to sort students by marks

def sort_students():
    students = [
        ("Alice", 85),
        ("Bob", 72),
        ("Charlie", 90),
        ("Diana", 68)
    ]
    students.sort(key=lambda x: x[1], reverse=True)
    print(students)

# Exercise 18: Filtering Passing Students
# Program to filter students passing based on a threshold

def filter_passing_students():
    students = [
        ("Alice", 85),
        ("Bob", 72),
        ("Charlie", 90),
        ("Diana", 68)
    ]
    passing = list(filter(lambda x: x[1] >= 70, students))
    print(passing)
