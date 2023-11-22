// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"product-management-app/api"
	"product-management-app/database"
	"product-management-app/image_analysis"
)

func main() {
	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Set up API routes
	router := api.NewRouter(db)

	// Start API server
	go func() {
		log.Fatal(http.ListenAndServe(":8080", router))
	}()

	// Start image analysis worker
	go image_analysis.StartWorker()

	fmt.Println("Server is running on :8080")
	select {}
}
