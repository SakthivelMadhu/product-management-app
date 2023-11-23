// handlers.go

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"product-management-app/database"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql" // Uncomment this line for MySQL
	"github.com/gorilla/mux"
)

// ProductHandler handles product-related API requests.
type ProductHandler struct {
	repo *database.Repository
}

// NewProductHandler creates a new instance of ProductHandler.
func NewProductHandler(repo *database.Repository) *ProductHandler {
	return &ProductHandler{repo}
}

// CreateProductHandler handles the creation of a new product.
func (ph *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var product database.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Set timestamps
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	// Perform validation if needed
	if err := validateProduct(&product); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %v", err), http.StatusBadRequest)
		return
	}

	// Create product in the database
	productID, err := ph.repo.CreateProduct(&product)
	if err != nil {
		// Log the error for debugging
		log.Printf("Error creating product: %v", err)
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	// Respond with the created product ID
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"product_id": productID})
}

// validateProduct performs input validation on the product.
func validateProduct(product *database.Product) error {
	// Example validation logic
	if product.ProductName == "" {
		return fmt.Errorf("product name is required")
	}
	if product.ProductPrice < 0 {
		return fmt.Errorf("product price cannot be negative")
	}
	// Add more validation rules as needed
	return nil
}

// GetUserHandler handles requests to retrieve user information.
func (ph *ProductHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from the request, for example, from the URL path or query parameters
	userID, err := extractUserID(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error extracting user ID: %v", err), http.StatusBadRequest)
		return
	}

	// Retrieve user from the database
	user, err := ph.repo.GetUser(userID)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	// Respond with the user data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := extractUserID(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error extracting user ID: %v", err), http.StatusBadRequest)
		return
	}

	// Now you can use the userID in your handler logic
	fmt.Fprintf(w, "User ID: %d", userID)
}

func extractUserID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	userIDStr, ok := vars["userID"]
	if !ok {
		return 0, fmt.Errorf("user ID not found in URL path")
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert user ID to integer: %v", err)
	}

	return userID, nil
}
