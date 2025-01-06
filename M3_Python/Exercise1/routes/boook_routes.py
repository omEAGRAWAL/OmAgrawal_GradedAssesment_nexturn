from flask import Blueprint, request, jsonify
from models.book_model import db, Book

book_bp = Blueprint('books', __name__)

@book_bp.route('/', methods=['POST'])
def add_book():
    data = request.json
    try:
        new_book = Book(
            title=data['title'],
            author=data['author'],
            published_year=data['published_year'],
            genre=data['genre']
        )
        db.session.add(new_book)
        db.session.commit()
        return jsonify({'message': 'Book added successfully', 'book_id': new_book.id}), 201
    except Exception as e:
        return jsonify({'error': 'Invalid data', 'message': str(e)}), 400

@book_bp.route('/', methods=['GET'])
def get_all_books():
    books = Book.query.all()
    return jsonify([book.to_dict() for book in books]), 200

@book_bp.route('/<int:book_id>', methods=['GET'])
def get_book_by_id(book_id):
    book = Book.query.get(book_id)
    if book:
        return jsonify(book.to_dict()), 200
    return jsonify({'error': 'Book not found', 'message': 'No book exists with the provided ID'}), 404

@book_bp.route('/<int:book_id>', methods=['PUT'])
def update_book(book_id):
    book = Book.query.get(book_id)
    if not book:
        return jsonify({'error': 'Book not found', 'message': 'No book exists with the provided ID'}), 404

    data = request.json
    try:
        book.title = data['title']
        book.author = data['author']
        book.published_year = data['published_year']
        book.genre = data['genre']
        db.session.commit()
        return jsonify({'message': 'Book updated successfully'}), 200
    except Exception as e:
        return jsonify({'error': 'Invalid data', 'message': str(e)}), 400

@book_bp.route('/<int:book_id>', methods=['DELETE'])
def delete_book(book_id):
    book = Book.query.get(book_id)
    if book:
        db.session.delete(book)
        db.session.commit()
        return jsonify({'message': 'Book deleted successfully'}), 200
    return jsonify({'error': 'Book not found', 'message': 'No book exists with the provided ID'}), 404
