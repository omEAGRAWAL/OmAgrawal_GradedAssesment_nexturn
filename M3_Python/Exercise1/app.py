from flask import Flask
from routes.book_routes import book_bp
from config import DATABASE

app = Flask(__name__)
app.config['DATABASE'] = DATABASE

# Register the Blueprint for book routes
app.register_blueprint(book_bp, url_prefix='/books')

if __name__ == '__main__':
    app.run(debug=True)
