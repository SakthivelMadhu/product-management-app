// main.go

package main

import (
	"Product_Management_App/api/handlers"
	"Product_Management_App/api/routes"
	"Product_Management_App/database"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Run migrations (you may need to adjust this based on your migration mechanism)
	err = database.RunMigrations(db)
	if err != nil {
		log.Fatal("Error running migrations:", err)
	}

	// Create instances of handlers
	productHandler := handlers.NewProductHandler(db)

	// Create a new Gorilla mux router
	router := mux.NewRouter()

	// Define routes
	routes.SetupRoutes(router, productHandler)

	// Specify the port to run the server on
	port := 8080

	// Start the server
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
