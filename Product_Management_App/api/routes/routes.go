// api/routes/routes.go

package routes

import (
	"Product_Management_App/api/handlers"

	"github.com/gorilla/mux"
)

// SetProductRoutes sets up product-related routes.
func SetProductRoutes(router *mux.Router, productHandler *handlers.ProductHandler) {
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{product_id}", productHandler.GetProduct).Methods("GET")
	router.HandleFunc("/products/{product_id}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{product_id}", productHandler.DeleteProduct).Methods("DELETE")
}

// SetUserRoutes sets up user-related routes.
func SetUserRoutes(router *mux.Router, userHandler *handlers.UserHandler) {
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{user_id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users/{user_id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{user_id}", userHandler.DeleteUser).Methods("DELETE")
}

// SetupRoutes configures the API routes.
func SetupRoutes(router *mux.Router, productHandler *handlers.ProductHandler) {
	router.HandleFunc("/api/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", productHandler.GetProductByID).Methods("GET")
	router.HandleFunc("/api/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", productHandler.DeleteProduct).Methods("DELETE")
}
