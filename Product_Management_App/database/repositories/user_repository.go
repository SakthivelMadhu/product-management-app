// database/repositories/user_repository.go

package repositories

import (
	"Product_Management_App/api/models"
	"database/sql"
	"fmt"
	"log"
)

// UserRepository handles database operations related to users.
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser creates a new user in the database.
func (ur *UserRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (name, mobile, latitude, longitude, created_at, updated_at)
		VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id
	`

	err := ur.DB.QueryRow(
		query,
		user.Name,
		user.Mobile,
		user.Latitude,
		user.Longitude,
	).Scan(&user.ID)

	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		return err
	}

	return nil
}

// GetUserByID retrieves a user by ID from the database.
func (ur *UserRepository) GetUserByID(userID int) (*models.User, error) {
	query := `
		SELECT * FROM users
		WHERE id = $1
	`

	var user models.User
	err := ur.DB.QueryRow(query, userID).Scan(
		&user.ID,
		&user.Name,
		&user.Mobile,
		&user.Latitude,
		&user.Longitude,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error retrieving user with ID %d: %v\n", userID, err)
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates an existing user in the database.
func (ur *UserRepository) UpdateUser(userID int, updatedUser *models.User) error {
	query := `
		UPDATE users
		SET name = $1, mobile = $2, latitude = $3, longitude = $4, updated_at = CURRENT_TIMESTAMP
		WHERE id = $5
	`

	_, err := ur.DB.Exec(
		query,
		updatedUser.Name,
		updatedUser.Mobile,
		updatedUser.Latitude,
		updatedUser.Longitude,
		userID,
	)

	if err != nil {
		log.Printf("Error updating user with ID %d: %v\n", userID, err)
		return err
	}

	return nil
}

// DeleteUser deletes a user by ID from the database.
func (ur *UserRepository) DeleteUser(userID int) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	result, err := ur.DB.Exec(query, userID)
	if err != nil {
		log.Printf("Error deleting user with ID %d: %v\n", userID, err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected after deleting user with ID %d: %v\n", userID, err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No user found with ID %d", userID)
	}

	return nil
}
