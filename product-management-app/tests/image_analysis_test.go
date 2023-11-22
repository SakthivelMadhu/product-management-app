// image_analysis_test.go

package image_analysis

import (
	"product-management-app/database"
	"testing"
	"time"
)

func TestStartWorker(t *testing.T) {
	// Implement your test logic for StartWorker
	// ...

	// Example:
	repo := &mockRepository{} // Implement a mock repository for testing
	worker := NewWorker(repo)

	// Use a shorter interval for testing
	interval := time.Millisecond * 10
	go worker.StartWorker(interval)

	// Let the worker run for a short duration for testing
	time.Sleep(time.Millisecond * 50)

	// Assertions
	// Add assertions based on the behavior of your worker
	// ...

	// Stop the worker (if needed)
	// ...

	// Add more assertions as needed
}

// Implement other tests as needed.

// Mock Repository for testing
type mockRepository struct {
}

func (m *mockRepository) GetProductsWithoutCompressedImages() ([]database.Product, error) {
	// Implement mock logic
	return nil, nil
}

func (m *mockRepository) UpdateProduct(product *database.Product) error {
	// Implement mock logic
	return nil
}
