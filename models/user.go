package models

import (
	"database/sql"
	"time"
)

// User represents a user in the system
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Password is not included in JSON responses
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserTable creates the users table if it doesn't exist
func CreateUserTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	return err
}

// GetUsers retrieves all users from the database
func GetUsers(db *sql.DB) ([]User, error) {
	query := "SELECT id, username, email, created_at, updated_at FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUser retrieves a user by ID
func GetUser(db *sql.DB, id int) (User, error) {
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?"
	var user User
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

// CreateUser creates a new user in the database
func CreateUser(db *sql.DB, user User) (int, error) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	result, err := db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

// UpdateUser updates an existing user
func UpdateUser(db *sql.DB, user User) error {
	query := "UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?"
	_, err := db.Exec(query, user.Username, user.Email, user.Password, user.ID)
	return err
}

// DeleteUser deletes a user by ID
func DeleteUser(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}