// tests/unit/unit_test.go

package unit

import (
	"Product_Management_App/image"
	"Product_Management_App/image/analysis"
	"testing"
)

func TestDownloadAndCompressImageFailure(t *testing.T) {
	// Create an instance of ImageAnalyzer
	ia := &analysis.ImageAnalyzer{}

	// Simulate downloading and compressing a single image with a failure
	compressedImagePath, err := ia.DownloadAndCompressImage("invalid-url")

	// Add assertions based on your actual implementation
	if err == nil {
		t.Error("Expected an error during image analysis, got nil.")
	}

	if compressedImagePath != "" {
		t.Error("Expected an empty compressed image path due to failure, got a non-empty string.")
	}
}

func TestDownloadAndCompressImage(t *testing.T) {
	// Create an instance of ImageAnalyzer
	ia := &analysis.ImageAnalyzer{}

	// Replace the URL with an actual image URL from your implementation
	imageURL := "https://example.com/image.jpg"
	compressedImagePath, err := ia.DownloadAndCompressImage(imageURL)

	// Add assertions based on your actual implementation
	if err != nil {
		t.Errorf("Error during image analysis: %v", err)
	}

	if compressedImagePath == "" {
		t.Error("Expected a non-empty compressed image path, got an empty string.")
	}
}

func TestCompressAndStoreImages(t *testing.T) {
	// Create an instance of ImageAnalyzer
	ia := &analysis.ImageAnalyzer{}

	// Simulate original images
	originalImages := []string{"https://example.com/image1.jpg", "https://example.com/image2.jpg"}

	// Simulate product ID
	productID := 1

	// Create a channel for receiving the results
	resultChan := make(chan struct {
		CompressedPaths []string
		Err             error
	})

	// Trigger image analysis asynchronously
	ia.TriggerImageAnalysis(productID, originalImages, resultChan)

	// Receive the result from the channel
	result := <-resultChan

	// Check if there was an error
	if result.Err != nil {
		t.Errorf("Error during image analysis: %v", result.Err)
	}

	// For example, check if compressedPaths is not empty, etc.
	if len(result.CompressedPaths) == 0 {
		t.Error("Expected non-empty compressed images, got an empty slice.")
	}
}

func TestPerformImageAnalysis(t *testing.T) {
	// Create an instance of ImageAnalyzer
	ia := &analysis.ImageAnalyzer{}

	// Create an instance of ImageService
	is := image.NewImageService(ia)

	// Simulate original images
	originalImages := []string{"https://example.com/image1.jpg", "https://example.com/image2.jpg"}

	// Simulate product ID
	productID := 1

	// Perform image analysis
	err := is.PerformImageAnalysis(productID, originalImages)
	if err != nil {
		t.Errorf("Error performing image analysis: %v", err)
	}
}

func TestCompressAndStoreImagesFailure(t *testing.T) {
	// Create an instance of ImageAnalyzer with a mocked failure
	ia := &analysis.ImageAnalyzer{MockFailure: true}

	// Simulate original images
	originalImages := []string{"https://example.com/image1.jpg", "https://example.com/image2.jpg"}

	// Simulate product ID
	productID := 1

	// Compress and store images with a failure
	err := ia.CompressAndStoreImages(productID, originalImages)
	if err == nil {
		t.Error("Expected an error during image analysis, got nil.")
	}
}

func TestPerformImageAnalysisFailure(t *testing.T) {
	// Create an instance of ImageAnalyzer with a mocked failure
	ia := &analysis.ImageAnalyzer{MockFailure: true}

	// Create an instance of ImageService
	is := image.NewImageService(ia)

	// Simulate original images
	originalImages := []string{"https://example.com/image1.jpg", "https://example.com/image2.jpg"}

	// Simulate product ID
	productID := 1

	// Perform image analysis with a failure
	err := is.PerformImageAnalysis(productID, originalImages)
	if err == nil {
		t.Error("Expected an error during image analysis, got nil.")
	}
}
