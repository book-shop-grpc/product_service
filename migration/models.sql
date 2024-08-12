CREATE DATABASE book_shop;

CREATE TABLE admin(
    id UUID PRIMARY KEY,
    name VARCHAR(60) NOT NULL
);

CREATE TABLE authorth(    
    author_id UUID PRIMARY KEY NOT NULL,
    full_name VARCHAR(128) NOT NULL,
    data_of_birth TIMESTAMP,
    data_of_death TIMESTAMP
);

CREATE TABLE books(
    book_id UUID PRIMARY KEY NOT NULL,
    book_name VARCHAR(64) NOT NULL,
    price DECIMAL(6,2),
    created_at TIMESTAMP DEFAULT current_timestamp,
    quantity INT DEFAULT 0,
    is_sail BOOLEAN DEFAULT false,
    author_id UUID REFERENCES authors(author_id)
);

CREATE TABLE users(
    user_id UUID PRIMARY KEY NOT NULL,
    username VARCHAR(16) NOT NULL,
    pone_number VARCHAR(10),
    password VARCHAR(60)
);

CREATE UNIQUE INDEX ON users(LOWER(username));

CREATE TABLE baskets(
    basket_id UUID PRIMARY KEY NOT NULL,
    user_id UUID REFERENCES users(user_id),
    quantity INT DEFAULT 0
);

CREATE TABLE selected_books (
    seletion_id UUID PRIMARY KEY NOT NULL,
    quantity INT DEFAULT 0,
    book_id UUID REFERENCES books(book_id),
    basket_id UUID REFERENCES baskets(basket_id),
    user_id UUID REFERENCES users(user_id)
);