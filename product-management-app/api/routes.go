// routes.go

package api

import (
	"database/sql"
	"product-management-app/database"

	"github.com/gorilla/mux"
)

// Set up your API routes using the Gorilla Mux router.
func SetupRoutes(db *sql.DB) *mux.Router {
	// Create a new instance of the repository
	repo := database.NewRepository(db)

	// Create a new instance of the product handler
	productHandler := NewProductHandler(repo)

	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Set up your API routes
	router.HandleFunc("/api/products", productHandler.CreateProductHandler).Methods("POST")
	router.HandleFunc("/api/users/{userID}", productHandler.GetUserHandler).Methods("GET")

	// Add more routes as needed

	return router
}

// NewRouter sets up the API routes.
func NewRouter(repo *database.Repository) *mux.Router {
	router := mux.NewRouter()

	// Create an instance of ProductHandler
	productHandler := NewProductHandler(repo)

	// Define your API routes using Gorilla Mux
	router.HandleFunc("/products", productHandler.CreateProductHandler).Methods("POST")
	router.HandleFunc("/users/{userID}", productHandler.GetUserHandler).Methods("GET")

	// Add more routes as needed

	return router
}
