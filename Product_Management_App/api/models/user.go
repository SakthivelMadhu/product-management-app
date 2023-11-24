// api/models/user.go

package models

// User represents the data model for a user.
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
