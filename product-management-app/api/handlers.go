// handlers.go

package api

import (
	"encoding/json"
	"net/http"
	"product-management-app/database"
	"time"
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

	// Create product in the database
	productID, err := ph.repo.CreateProduct(&product)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	// Respond with the created product ID
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"product_id": productID})
}

// GetUserHandler handles requests to retrieve user information.
func (ph *ProductHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from the request, for example, from the URL path or query parameters
	userID := extractUserID(r)

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
