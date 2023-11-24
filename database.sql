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


INSERT INTO users (name, mobile, latitude, longitude)
VALUES
    ('John Doe', '1234567890', 40.7128, -74.0060),
    ('Jane Smith', '9876543210', 37.7749, -122.4194);


INSERT INTO products (user_id, product_name, product_description, product_images, product_price)
VALUES
    (1, 'Product A', 'Description A', '["image1.jpg", "image2.jpg"]', 19.99),
    (2, 'Product B', 'Description B', '["image3.jpg", "image4.jpg"]', 29.99);
