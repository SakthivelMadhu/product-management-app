// database_test.go

package database

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Implement your test logic for CreateUser
	// ...

	// Example:
	db, err := sql.Open("postgres", "your-database-connection-string")
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	user := &User{
		Name:      "John Doe",
		Mobile:    "1234567890",
		Latitude:  40.7128,
		Longitude: -74.0060,
	}

	userID, err := repo.CreateUser(user)

	// Assertions
	assert.NoError(t, err)
	assert.NotZero(t, userID)
	// Add more assertions as needed
}

func TestGetUser(t *testing.T) {
	// Implement your test logic for GetUser
	// ...

	// Example:
	db, err := sql.Open("postgres", "your-database-connection-string")
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	userID := 1
	user, err := repo.GetUser(userID)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, user)
	// Add more assertions as needed
}

// Implement other tests as needed.
