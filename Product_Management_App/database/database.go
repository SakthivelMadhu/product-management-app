// database/database.go

package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB establishes a connection to the database and returns the *sql.DB object.
func ConnectDB() (*sql.DB, error) {
	// Update these values with your actual MySQL database credentials
	dbUser := "root"
	dbPassword := "Sakthivel1402!"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "product_management_app"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
		return nil, err
	}

	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return nil, err
	}

	fmt.Println("Connected to the database!")

	return db, nil
}

// RunMigrations performs database migrations.
func RunMigrations(db *sql.DB) error {
	// Your migration logic here
	fmt.Println("Running migrations...")
	return nil
}
