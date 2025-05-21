package models

import (
	"database/sql"
)

// InitTables initializes all database tables
func InitTables(db *sql.DB) error {
	// Create tables in the correct order to respect foreign key constraints
	if err := CreateUserTable(db); err != nil {
		return err
	}

	if err := CreateWorkoutTable(db); err != nil {
		return err
	}

	if err := CreateExerciseTable(db); err != nil {
		return err
	}

	if err := CreateWorkoutExerciseTable(db); err != nil {
		return err
	}

	if err := CreateProgressTable(db); err != nil {
		return err
	}

	return nil
}