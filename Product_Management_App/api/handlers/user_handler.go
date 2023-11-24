// api/handlers/user_handler.go

package handlers

import (
	"Product_Management_App/api/models"
	"Product_Management_App/database/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	UserRepo *repositories.UserRepository
}

// CreateUser handles the creation of a new user.
func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Create the user in the database
	err = uh.UserRepo.CreateUser(&user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetUser retrieves a user by ID.
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := uh.UserRepo.GetUserByID(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving user with ID %d", userID), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser updates an existing user.
func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update the user in the database
	err = uh.UserRepo.UpdateUser(userID, &updatedUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating user with ID %d", userID), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser deletes a user by ID.
func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Delete the user from the database
	err = uh.UserRepo.DeleteUser(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting user with ID %d", userID), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
