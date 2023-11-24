// api/models/product.go

package models

// Product represents the data model for a product.
type Product struct {
	ProductID               int      `json:"product_id"`
	ProductName             string   `json:"product_name"`
	ProductDescription      string   `json:"product_description"`
	ProductImages           []string `json:"product_images"`
	ProductPrice            float64  `json:"product_price"`
	CompressedProductImages []string `json:"compressed_product_images,omitempty"`
	CreatedAt               string   `json:"created_at,omitempty"`
	UpdatedAt               string   `json:"updated_at,omitempty"`
}
