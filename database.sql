-- Create a database
CREATE DATABASE IF NOT EXISTS your_database_name;
USE your_database_name;

CREATE TABLE Users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    mobile VARCHAR(15),
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE Products (
    product_id INT PRIMARY KEY AUTO_INCREMENT,
    product_name VARCHAR(255),
    product_description TEXT,
    product_images JSON,
    product_price DECIMAL(10, 2),
    compressed_product_images JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
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
