// tests/integration/integration_test.go

package integration

import (
	"Product_Management_App/api/handlers"
	"Product_Management_App/api/models"
	"Product_Management_App/database"
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db *sql.DB
var productHandler *handlers.ProductHandler

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	var err error

	// Connect to the database
	db, err = database.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Initialize the ProductHandler
	productHandler = handlers.NewProductHandler(db)

	// Run migrations (you may need to adjust this based on your migration mechanism)
	err = database.RunMigrations(db)
	if err != nil {
		log.Fatal("Error running migrations:", err)
	}
}

func tearDown() {
	// Close the database connection
	err := db.Close()
	if err != nil {
		log.Fatal("Error closing the database connection:", err)
	}
}

func TestProductIntegration(t *testing.T) {
	// Create a new product
	newProduct := models.Product{
		ProductName:        "Test Product",
		ProductDescription: "Description for Test Product",
		ProductImages:      []string{"image1.jpg", "image2.jpg"},
		ProductPrice:       19.99,
	}

	// Convert the new product to JSON
	jsonData, err := json.Marshal(newProduct)
	assert.Nil(t, err)

	// Create a request to simulate an HTTP POST request with the product data
	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonData))
	assert.Nil(t, err)

	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Call the CreateProduct function directly (assuming the function signature in handlers.ProductHandler)
	productHandler.CreateProduct(w, req)

	// Verify the HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Decode the response body to get the created product
	var createdProduct models.Product
	err = json.Unmarshal(w.Body.Bytes(), &createdProduct)
	assert.Nil(t, err)

	// Verify the created product's attributes
	assert.NotNil(t, createdProduct.ProductID)
	assert.Equal(t, newProduct.ProductName, createdProduct.ProductName)
	assert.Equal(t, newProduct.ProductDescription, createdProduct.ProductDescription)
	assert.Equal(t, newProduct.ProductImages, createdProduct.ProductImages)
	assert.Equal(t, newProduct.ProductPrice, createdProduct.ProductPrice)

	// Additional assertions or verifications as needed
}
