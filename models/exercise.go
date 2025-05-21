package models

import (
	"database/sql"
	"time"
)

// Exercise represents a physical exercise
type Exercise struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateExerciseTable creates the exercises table if it doesn't exist
func CreateExerciseTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS exercises (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		description TEXT,
		category VARCHAR(50) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	return err
}

// GetExercises retrieves all exercises from the database
func GetExercises(db *sql.DB) ([]Exercise, error) {
	query := "SELECT id, name, description, category, created_at, updated_at FROM exercises"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []Exercise
	for rows.Next() {
		var exercise Exercise
		if err := rows.Scan(&exercise.ID, &exercise.Name, &exercise.Description, &exercise.Category, &exercise.CreatedAt, &exercise.UpdatedAt); err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}

	return exercises, nil
}

// GetExercise retrieves an exercise by ID
func GetExercise(db *sql.DB, id int) (Exercise, error) {
	query := "SELECT id, name, description, category, created_at, updated_at FROM exercises WHERE id = ?"
	var exercise Exercise
	err := db.QueryRow(query, id).Scan(&exercise.ID, &exercise.Name, &exercise.Description, &exercise.Category, &exercise.CreatedAt, &exercise.UpdatedAt)
	return exercise, err
}

// CreateExercise creates a new exercise in the database
func CreateExercise(db *sql.DB, exercise Exercise) (int, error) {
	query := "INSERT INTO exercises (name, description, category) VALUES (?, ?, ?)"
	result, err := db.Exec(query, exercise.Name, exercise.Description, exercise.Category)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

// UpdateExercise updates an existing exercise
func UpdateExercise(db *sql.DB, exercise Exercise) error {
	query := "UPDATE exercises SET name = ?, description = ?, category = ? WHERE id = ?"
	_, err := db.Exec(query, exercise.Name, exercise.Description, exercise.Category, exercise.ID)
	return err
}

// DeleteExercise deletes an exercise by ID
func DeleteExercise(db *sql.DB, id int) error {
	query := "DELETE FROM exercises WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}