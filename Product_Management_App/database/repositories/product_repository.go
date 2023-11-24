// database/repositories/product_repository.go

package repositories

import (
	"Product_Management_App/api/models"
	"database/sql"
	"fmt"
	"log"
)

// ProductRepository handles database operations related to products.
type ProductRepository struct {
	DB *sql.DB
}

// NewProductRepository creates a new instance of ProductRepository.
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// CreateProduct creates a new product in the database.
func (pr *ProductRepository) CreateProduct(product *models.Product) error {
	query := `
		INSERT INTO products (product_name, product_description, product_images, product_price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING product_id
	`

	err := pr.DB.QueryRow(
		query,
		product.ProductName,
		product.ProductDescription,
		product.ProductImages,
		product.ProductPrice,
	).Scan(&product.ProductID)

	if err != nil {
		log.Printf("Error creating product: %v\n", err)
		return err
	}

	return nil
}

// GetProductByID retrieves a product by ID from the database.
func (pr *ProductRepository) GetProductByID(productID int) (*models.Product, error) {
	query := `
		SELECT * FROM products
		WHERE product_id = $1
	`

	var product models.Product
	err := pr.DB.QueryRow(query, productID).Scan(
		&product.ProductID,
		&product.ProductName,
		&product.ProductDescription,
		&product.ProductImages,
		&product.ProductPrice,
		&product.CompressedProductImages,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error retrieving product with ID %d: %v\n", productID, err)
		return nil, err
	}

	return &product, nil
}

// UpdateProduct updates an existing product in the database.
func (pr *ProductRepository) UpdateProduct(productID int, updatedProduct *models.Product) error {
	query := `
		UPDATE products
		SET product_name = $1, product_description = $2, product_images = $3, product_price = $4, updated_at = CURRENT_TIMESTAMP
		WHERE product_id = $5
	`

	_, err := pr.DB.Exec(
		query,
		updatedProduct.ProductName,
		updatedProduct.ProductDescription,
		updatedProduct.ProductImages,
		updatedProduct.ProductPrice,
		productID,
	)

	if err != nil {
		log.Printf("Error updating product with ID %d: %v\n", productID, err)
		return err
	}

	return nil
}

// DeleteProduct deletes a product by ID from the database.
func (pr *ProductRepository) DeleteProduct(productID int) error {
	query := `
		DELETE FROM products
		WHERE product_id = $1
	`

	result, err := pr.DB.Exec(query, productID)
	if err != nil {
		log.Printf("Error deleting product with ID %d: %v\n", productID, err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected after deleting product with ID %d: %v\n", productID, err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No product found with ID %d", productID)
	}

	return nil
}

// UpdateProductImages updates the product images in the database.
func (pr *ProductRepository) UpdateProductImages(productID int, compressedImages []string) error {
	// Your logic to update product images in the database
	// ...

	// Example: Update product images in the database
	_, err := pr.DB.Exec("UPDATE products SET compressed_product_images = ? WHERE product_id = ?", compressedImages, productID)
	if err != nil {
		return err
	}

	return nil
}
