import sqlite3
from config import DATABASE

def initialize_db():
    connection = sqlite3.connect(DATABASE)
    cursor = connection.cursor()
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS books (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            author TEXT NOT NULL,
            published_year INTEGER NOT NULL,
            genre TEXT NOT NULL
        )
    """)
    connection.commit()
    connection.close()

if __name__ == '__main__':
    initialize_db()
    print("Database initialized successfully.")
