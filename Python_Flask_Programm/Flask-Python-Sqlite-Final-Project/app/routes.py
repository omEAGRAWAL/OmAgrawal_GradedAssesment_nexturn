from flask import Blueprint, render_template, request, flash, redirect, url_for
from app.models import User
from . import db

main = Blueprint('main', __name__)

@main.route('/')
def index():
    users = User.query.all()
    return render_template('index.html', users=users)

@main.route('/create', methods=['GET', 'POST'])
def create():
    if request.method == 'POST':
        name = request.form.get('name')
        email = request.form.get('email')
        new_user = User(name=name, email=email)
        db.session.add(new_user)
        db.session.commit()
        flash('User created successfully')
        return redirect(url_for('main.index'))
    return render_template('create.html')

@main.route('/edit/<int:id>', methods=['GET', 'POST'])
def edit(id):
    user = User.query.get_or_404(id)
    if request.method == 'POST':
        user.name = request.form.get('name')
        user.email = request.form.get('email')
        db.session.commit()
        flash('User updated successfully')
        return redirect(url_for('main.index'))
    return render_template('edit.html', user=user)

@main.route('/delete/<int:id>')
def delete(id):
    user = User.query.get(id)
    if not user:
        flash('User not found')
        return redirect(url_for('main.index'))
    db.session.delete(user)
    db.session.commit()
    flash('User deleted successfully')
    return redirect(url_for('main.index'))