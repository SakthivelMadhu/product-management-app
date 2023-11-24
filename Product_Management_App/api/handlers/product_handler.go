// api/handlers/product_handler.go

package handlers

import (
	"Product_Management_App/api/models"
	"Product_Management_App/database/repositories"
	"Product_Management_App/image/analysis"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ProductHandler handles product-related HTTP requests.
type ProductHandler struct {
	ProductRepo   *repositories.ProductRepository
	DB            *sql.DB
	ImageAnalyzer *analysis.ImageAnalyzer
}

// NewProductHandler creates a new instance of ProductHandler.
func NewProductHandler(db *sql.DB) *ProductHandler {
	// Your constructor logic here
	return &ProductHandler{DB: db}
}

// CreateProduct handles the creation of a new product.
func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Create the product in the database
	err = ph.ProductRepo.CreateProduct(&product)
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	// Perform image analysis and update the database with compressed image paths
	compressedPaths, err := analysis.AnalyzeImages(product.ProductID, product.ProductImages)
	if err != nil {
		fmt.Printf("Error analyzing images for Product ID %d: %v\n", product.ProductID, err)
		// Handle error accordingly
	}

	// Update the product in the database with compressed image paths
	err = ph.ProductRepo.UpdateProductImages(product.ProductID, compressedPaths)
	if err != nil {
		fmt.Printf("Error updating product with compressed images for Product ID %d: %v\n", product.ProductID, err)
		// Handle error accordingly
	}

	// Trigger image analysis asynchronously
	// ph.ImageAnalyzer.TriggerImageAnalysis(product.ProductID, product.ProductImages)

	// Create an instance of ProductHandler
	ph = InitializeProductHandler()
	// Correct usage with a new variable
	// var ph *handlers.ProductHandler
	// ph = handlers.InitializeProductHandler()

	// Simulate original images
	originalImages := []string{"https://stat.overdrive.in/wp-content/odgallery/2022/07/63248_2022_Suzuki_Brezza_1.jpg", "https://stat.overdrive.in/wp-content/odgallery/2018/08/46307_Maruti_Suzuki_Vitara_Breeza_003.JPG"}

	// Simulate product ID
	productID := 1

	// Create a channel for receiving the results
	resultChan := make(chan struct {
		CompressedPaths []string
		Err             error
	})

	// Trigger image analysis asynchronously
	ph.ImageAnalyzer.TriggerImageAnalysis(productID, originalImages, resultChan)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// GetProduct retrieves a product by ID.
func (ph *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID, err := strconv.Atoi(params["product_id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := ph.ProductRepo.GetProductByID(productID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving product with ID %d", productID), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// GetProductByID retrieves a product by ID.
func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the request parameters
	vars := mux.Vars(r)
	productID := vars["id"]

	// Your logic to fetch the product from the database by ID
	// ...

	// For example, let's send a JSON response with a dummy product
	dummyProduct := map[string]interface{}{
		"productID":     productID,
		"productName":   "Dummy Product",
		"description":   "This is a dummy product.",
		"productImages": []string{"image1.jpg", "image2.jpg"},
		"productPrice":  29.99,
	}

	// Convert the product to JSON
	responseJSON, err := json.Marshal(dummyProduct)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	// Set the content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(responseJSON)
}

// UpdateProduct updates an existing product.
func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID, err := strconv.Atoi(params["product_id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update the product in the database
	err = ph.ProductRepo.UpdateProduct(productID, &updatedProduct)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating product with ID %d", productID), http.StatusInternalServerError)
		return
	}

	// Perform image analysis and update the database with compressed image paths
	compressedPaths, err := analysis.AnalyzeImages(productID, updatedProduct.ProductImages)
	if err != nil {
		fmt.Printf("Error analyzing images for updated Product ID %d: %v\n", productID, err)
		// Handle error accordingly
	}

	// Update the product in the database with compressed image paths
	err = ph.ProductRepo.UpdateProductImages(productID, compressedPaths)
	if err != nil {
		fmt.Printf("Error updating product with compressed images for updated Product ID %d: %v\n", productID, err)
		// Handle error accordingly
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProduct)
}

// DeleteProduct deletes a product by ID.
func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID, err := strconv.Atoi(params["product_id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Delete the product from the database
	err = ph.ProductRepo.DeleteProduct(productID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting product with ID %d", productID), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// InitializeProductHandler initializes a new ProductHandler.
func InitializeProductHandler() *ProductHandler {
	// Initialize other fields...

	// Initialize ImageAnalyzer
	ia := &analysis.ImageAnalyzer{}
	return &ProductHandler{
		// Other field assignments...
		ImageAnalyzer: ia,
	}
}
