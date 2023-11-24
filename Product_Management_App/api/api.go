// api/api.go

package api

import (
	"Product_Management_App/api/handlers"
	"Product_Management_App/api/routes"
	"Product_Management_App/database"
	"Product_Management_App/database/repositories"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StartAPI initializes and starts the API server.
func StartAPI() {
	// Initialize database connection
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Initialize repositories
	productRepo := repositories.NewProductRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Initialize handlers
	productHandler := &handlers.ProductHandler{ProductRepo: productRepo}
	userHandler := &handlers.UserHandler{UserRepo: userRepo}

	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Set up routes for products and users
	routes.SetProductRoutes(router, productHandler)
	routes.SetUserRoutes(router, userHandler)

	// Define a default route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the Product Management API!")
	})

	// Start the HTTP server
	port := ":8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
