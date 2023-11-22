// repository.go

package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import your SQL driver, replace with your database driver
)

// Repository handles database operations.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new instance of the Repository.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

// CreateUser creates a new user in the database.
func (r *Repository) CreateUser(user *User) (int, error) {
	// Example SQL statement, replace with your actual schema
	query := `
		INSERT INTO users (name, mobile, latitude, longitude, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var userID int
	err := r.db.QueryRow(query, user.Name, user.Mobile, user.Latitude, user.Longitude, user.CreatedAt, user.UpdatedAt).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %v", err)
	}

	return userID, nil
}

// GetUser retrieves a user by ID from the database.
func (r *Repository) GetUser(userID int) (*User, error) {
	// Example SQL statement, replace with your actual schema
	query := `
		SELECT id, name, mobile, latitude, longitude, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &User{}
	err := r.db.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Mobile, &user.Latitude, &user.Longitude, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	return user, nil
}

// CreateProduct creates a new product in the database.
func (r *Repository) CreateProduct(product *Product) (int, error) {
	// Example SQL statement, replace with your actual schema
	query := `
		INSERT INTO products (user_id, product_name, product_description, product_images, product_price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING product_id
	`

	var productID int
	err := r.db.QueryRow(query, product.UserID, product.ProductName, product.ProductDescription, product.ProductImages, product.ProductPrice, product.CreatedAt, product.UpdatedAt).Scan(&productID)
	if err != nil {
		return 0, fmt.Errorf("error creating product: %v", err)
	}

	return productID, nil
}

// GetProduct retrieves a product by ID from the database.
func (r *Repository) GetProduct(productID int) (*Product, error) {
	// Example SQL statement, replace with your actual schema
	query := `
		SELECT product_id, user_id, product_name, product_description, product_images, product_price, created_at, updated_at
		FROM products
		WHERE product_id = $1
	`

	product := &Product{}
	err := r.db.QueryRow(query, productID).Scan(&product.ProductID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductImages, &product.ProductPrice, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error getting product: %v", err)
	}

	return product, nil
}
