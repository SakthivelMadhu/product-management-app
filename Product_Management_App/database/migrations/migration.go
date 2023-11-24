// database/migrations/migration.go

package migrations

import (
	"database/sql"
	"fmt"
	"log"
)

// RunMigrations runs database migrations.
func RunMigrations(db *sql.DB) error {
	// Add your migration logic here
	// Example: Create tables, modify schema, etc.
	// You can use a migration library like https://github.com/golang-migrate/migrate

	fmt.Println("Running database migrations...")

	// Example: Create a users table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(255),
			mobile VARCHAR(15),
			latitude VARCHAR(20),
			longitude VARCHAR(20),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal("Error creating users table:", err)
		return err
	}

	// Example: Create a products table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			product_id INT PRIMARY KEY AUTO_INCREMENT,
			product_name VARCHAR(255),
			product_description TEXT,
			product_images JSON,
			product_price DECIMAL(10, 2),
			compressed_product_images JSON,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal("Error creating products table:", err)
		return err
	}

	fmt.Println("Database migrations completed!")

	return nil
}
