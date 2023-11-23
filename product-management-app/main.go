// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"product-management-app/api"
	"product-management-app/database"
	"product-management-app/image_analysis"
	"time"

	_ "github.com/go-sql-driver/mysql" // Uncomment this line for MySQL
)

func main() {
	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repository
	repo := database.NewRepository(db)

	// Set up API routes
	router := api.NewRouter(repo)

	// Start API server
	go func() {
		log.Fatal(http.ListenAndServe(":8080", router))
	}()

	// Start image analysis worker
	worker := image_analysis.NewWorker(repo)
	go worker.StartWorker(time.Second * 10) // Adjust the interval as needed

	fmt.Println("Server is running on :8080")
	select {}
}
