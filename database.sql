

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


-- Insert data into Users table
INSERT INTO Users (name, mobile, latitude, longitude) VALUES
('John Doe', '+1234567890', 37.7749, -122.4194),
('Jane Doe', '+9876543210', 40.7128, -74.0060),
('Alice Johnson', '+1122334455', 34.0522, -118.2437);


-- Insert data into Products table
INSERT INTO Products (product_name, product_description, product_images, product_price, compressed_product_images) VALUES
('Laptop', 'High-performance laptop', '["laptop_image1.jpg", "laptop_image2.jpg"]', 999.99, '["compressed_laptop_image1.jpg", "compressed_laptop_image2.jpg"]'),
('Smartphone', 'Latest smartphone model', '["phone_image1.jpg", "phone_image2.jpg"]', 599.99, '["compressed_phone_image1.jpg", "compressed_phone_image2.jpg"]'),
('Coffee Maker', 'Automatic coffee maker', '["coffeemaker_image1.jpg", "coffeemaker_image2.jpg"]', 149.99, '["compressed_coffeemaker_image1.jpg", "compressed_coffeemaker_image2.jpg"]');
