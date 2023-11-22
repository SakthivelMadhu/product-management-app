package database

import "time"

type User struct {
	ID        int
	Name      string
	Mobile    string
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ProductID               int
	UserID                  int
	ProductName             string
	ProductDescription      string
	ProductImages           []string
	ProductPrice            float64
	CompressedProductImages []string
	CreatedAt               time.Time
	UpdatedAt               time.Time
}
