package models

import (
	"database/sql"
)

// WorkoutExercise represents the association between workouts and exercises
type WorkoutExercise struct {
	WorkoutID  int `json:"workout_id"`
	ExerciseID int `json:"exercise_id"`
	Sets       int `json:"sets"`
	Reps       int `json:"reps"`
	Weight     int `json:"weight"`
	Order      int `json:"order"`
}

// CreateWorkoutExerciseTable creates the workout_exercises table if it doesn't exist
func CreateWorkoutExerciseTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS workout_exercises (
		workout_id INT NOT NULL,
		exercise_id INT NOT NULL,
		sets INT NOT NULL DEFAULT 3,
		reps INT NOT NULL DEFAULT 10,
		weight INT NOT NULL DEFAULT 0,
		exercise_order INT NOT NULL,
		PRIMARY KEY (workout_id, exercise_id),
		FOREIGN KEY (workout_id) REFERENCES workouts(id) ON DELETE CASCADE,
		FOREIGN KEY (exercise_id) REFERENCES exercises(id) ON DELETE CASCADE
	);`

	_, err := db.Exec(query)
	return err
}

// GetWorkoutExercises retrieves all exercises for a specific workout
func GetWorkoutExercises(db *sql.DB, workoutID int) ([]WorkoutExercise, error) {
	query := `
	SELECT we.workout_id, we.exercise_id, we.sets, we.reps, we.weight, we.exercise_order
	FROM workout_exercises we
	WHERE we.workout_id = ?
	ORDER BY we.exercise_order ASC`

	rows, err := db.Query(query, workoutID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workoutExercises []WorkoutExercise
	for rows.Next() {
		var we WorkoutExercise
		if err := rows.Scan(&we.WorkoutID, &we.ExerciseID, &we.Sets, &we.Reps, &we.Weight, &we.Order); err != nil {
			return nil, err
		}
		workoutExercises = append(workoutExercises, we)
	}

	return workoutExercises, nil
}

// AddExerciseToWorkout adds an exercise to a workout
func AddExerciseToWorkout(db *sql.DB, we WorkoutExercise) error {
	// Get the highest order value for the workout
	var maxOrder int
	err := db.QueryRow("SELECT COALESCE(MAX(exercise_order), 0) FROM workout_exercises WHERE workout_id = ?", we.WorkoutID).Scan(&maxOrder)
	if err != nil {
		return err
	}

	// Set the order to be one more than the highest current order
	we.Order = maxOrder + 1

	query := "INSERT INTO workout_exercises (workout_id, exercise_id, sets, reps, weight, exercise_order) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, we.WorkoutID, we.ExerciseID, we.Sets, we.Reps, we.Weight, we.Order)
	return err
}

// UpdateWorkoutExercise updates the details of an exercise in a workout
func UpdateWorkoutExercise(db *sql.DB, we WorkoutExercise) error {
	query := "UPDATE workout_exercises SET sets = ?, reps = ?, weight = ? WHERE workout_id = ? AND exercise_id = ?"
	_, err := db.Exec(query, we.Sets, we.Reps, we.Weight, we.WorkoutID, we.ExerciseID)
	return err
}

// RemoveExerciseFromWorkout removes an exercise from a workout
func RemoveExerciseFromWorkout(db *sql.DB, workoutID, exerciseID int) error {
	query := "DELETE FROM workout_exercises WHERE workout_id = ? AND exercise_id = ?"
	_, err := db.Exec(query, workoutID, exerciseID)
	return err
}

// ReorderWorkoutExercises updates the order of exercises in a workout
func ReorderWorkoutExercises(db *sql.DB, workoutID int, exerciseIDs []int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for i, exerciseID := range exerciseIDs {
		_, err := tx.Exec("UPDATE workout_exercises SET exercise_order = ? WHERE workout_id = ? AND exercise_id = ?", i+1, workoutID, exerciseID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}