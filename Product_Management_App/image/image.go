// image/image.go

package image

import (
	"Product_Management_App/image/analysis"
	"fmt"
)

// ImageService represents the service for handling image-related operations.
type ImageService struct {
	Analyzer *analysis.ImageAnalyzer
}

// NewImageService creates a new instance of ImageService.
func NewImageService(analyzer *analysis.ImageAnalyzer) *ImageService {
	return &ImageService{Analyzer: analyzer}
}

// PerformImageAnalysis initiates the image analysis process for a product.
func (is *ImageService) PerformImageAnalysis(productID int, originalImages []string) error {
	fmt.Printf("Initiating image analysis for product ID %d...\n", productID)

	// Call the image analysis component to compress and store images
	err := is.Analyzer.CompressAndStoreImages(productID, originalImages)
	if err != nil {
		fmt.Printf("Error performing image analysis for product ID %d: %v\n", productID, err)
		return err
	}

	fmt.Printf("Image analysis completed for product ID %d.\n", productID)
	return nil
}
