-- Create a database
CREATE DATABASE IF NOT EXISTS your_database_name;
USE your_database_name;

-- Create a table for users
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    mobile VARCHAR(15),
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create a table for products
CREATE TABLE IF NOT EXISTS products (
    product_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    product_name VARCHAR(255) NOT NULL,
    product_description TEXT,
    product_images JSON, -- assuming you store image paths in JSON format
    product_price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    compressed_product_images JSON -- assuming you store compressed image paths in JSON format
);

-- Create an index on the user_id column in the products table for faster retrieval
CREATE INDEX idx_user_id ON products (user_id);

-- Add foreign key constraint between products and users
ALTER TABLE products
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id) REFERENCES users(id);
