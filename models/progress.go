package models

import (
	"database/sql"
	"time"
)

// Progress represents a user's workout progress
type Progress struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	WorkoutID  int       `json:"workout_id"`
	ExerciseID int       `json:"exercise_id"`
	Sets       int       `json:"sets"`
	Reps       int       `json:"reps"`
	Weight     int       `json:"weight"`
	Notes      string    `json:"notes"`
	Date       time.Time `json:"date"`
	CreatedAt  time.Time `json:"created_at"`
}

// CreateProgressTable creates the progress table if it doesn't exist
func CreateProgressTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS progress (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		workout_id INT NOT NULL,
		exercise_id INT NOT NULL,
		sets INT NOT NULL,
		reps INT NOT NULL,
		weight INT NOT NULL,
		notes TEXT,
		date DATE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (workout_id) REFERENCES workouts(id) ON DELETE CASCADE,
		FOREIGN KEY (exercise_id) REFERENCES exercises(id) ON DELETE CASCADE
	);`

	_, err := db.Exec(query)
	return err
}

// GetUserProgress retrieves progress records for a specific user
func GetUserProgress(db *sql.DB, userID int) ([]Progress, error) {
	query := `
	SELECT id, user_id, workout_id, exercise_id, sets, reps, weight, notes, date, created_at
	FROM progress
	WHERE user_id = ?
	ORDER BY date DESC, created_at DESC`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var progressRecords []Progress
	for rows.Next() {
		var progress Progress
		if err := rows.Scan(&progress.ID, &progress.UserID, &progress.WorkoutID, &progress.ExerciseID, 
			&progress.Sets, &progress.Reps, &progress.Weight, &progress.Notes, &progress.Date, &progress.CreatedAt); err != nil {
			return nil, err
		}
		progressRecords = append(progressRecords, progress)
	}

	return progressRecords, nil
}

// GetExerciseProgress retrieves progress records for a specific exercise by a user
func GetExerciseProgress(db *sql.DB, userID, exerciseID int) ([]Progress, error) {
	query := `
	SELECT id, user_id, workout_id, exercise_id, sets, reps, weight, notes, date, created_at
	FROM progress
	WHERE user_id = ? AND exercise_id = ?
	ORDER BY date DESC, created_at DESC`

	rows, err := db.Query(query, userID, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var progressRecords []Progress
	for rows.Next() {
		var progress Progress
		if err := rows.Scan(&progress.ID, &progress.UserID, &progress.WorkoutID, &progress.ExerciseID, 
			&progress.Sets, &progress.Reps, &progress.Weight, &progress.Notes, &progress.Date, &progress.CreatedAt); err != nil {
			return nil, err
		}
		progressRecords = append(progressRecords, progress)
	}

	return progressRecords, nil
}

// RecordProgress adds a new progress record
func RecordProgress(db *sql.DB, progress Progress) (int, error) {
	query := `
	INSERT INTO progress (user_id, workout_id, exercise_id, sets, reps, weight, notes, date)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(query, progress.UserID, progress.WorkoutID, progress.ExerciseID, 
		progress.Sets, progress.Reps, progress.Weight, progress.Notes, progress.Date)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

// DeleteProgress deletes a progress record
func DeleteProgress(db *sql.DB, id, userID int) error {
	query := "DELETE FROM progress WHERE id = ? AND user_id = ?"
	_, err := db.Exec(query, id, userID)
	return err
}