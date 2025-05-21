package models

import (
	"database/sql"
	"time"
)

// Workout represents a workout plan
type Workout struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateWorkoutTable creates the workouts table if it doesn't exist
func CreateWorkoutTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS workouts (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		description TEXT,
		user_id INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	_, err := db.Exec(query)
	return err
}

// GetWorkouts retrieves all workouts from the database
func GetWorkouts(db *sql.DB) ([]Workout, error) {
	query := "SELECT id, name, description, user_id, created_at, updated_at FROM workouts"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workouts []Workout
	for rows.Next() {
		var workout Workout
		if err := rows.Scan(&workout.ID, &workout.Name, &workout.Description, &workout.UserID, &workout.CreatedAt, &workout.UpdatedAt); err != nil {
			return nil, err
		}
		workouts = append(workouts, workout)
	}

	return workouts, nil
}

// GetWorkout retrieves a workout by ID
func GetWorkout(db *sql.DB, id int) (Workout, error) {
	query := "SELECT id, name, description, user_id, created_at, updated_at FROM workouts WHERE id = ?"
	var workout Workout
	err := db.QueryRow(query, id).Scan(&workout.ID, &workout.Name, &workout.Description, &workout.UserID, &workout.CreatedAt, &workout.UpdatedAt)
	return workout, err
}

// GetUserWorkouts retrieves all workouts for a specific user
func GetUserWorkouts(db *sql.DB, userID int) ([]Workout, error) {
	query := "SELECT id, name, description, user_id, created_at, updated_at FROM workouts WHERE user_id = ?"
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workouts []Workout
	for rows.Next() {
		var workout Workout
		if err := rows.Scan(&workout.ID, &workout.Name, &workout.Description, &workout.UserID, &workout.CreatedAt, &workout.UpdatedAt); err != nil {
			return nil, err
		}
		workouts = append(workouts, workout)
	}

	return workouts, nil
}

// CreateWorkout creates a new workout in the database
func CreateWorkout(db *sql.DB, workout Workout) (int, error) {
	query := "INSERT INTO workouts (name, description, user_id) VALUES (?, ?, ?)"
	result, err := db.Exec(query, workout.Name, workout.Description, workout.UserID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

// UpdateWorkout updates an existing workout
func UpdateWorkout(db *sql.DB, workout Workout) error {
	query := "UPDATE workouts SET name = ?, description = ? WHERE id = ? AND user_id = ?"
	_, err := db.Exec(query, workout.Name, workout.Description, workout.ID, workout.UserID)
	return err
}

// DeleteWorkout deletes a workout by ID
func DeleteWorkout(db *sql.DB, id int, userID int) error {
	query := "DELETE FROM workouts WHERE id = ? AND user_id = ?"
	_, err := db.Exec(query, id, userID)
	return err
}