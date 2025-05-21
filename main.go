package main

import (
	"github.com/cxocodehub/go-backend-workout/handlers"
	"github.com/cxocodehub/go-backend-workout/models"
	"github.com/gofr-dev/gofr"
)

func main() {
	// Initialize the Gofr app
	app := gofr.New()

	// Initialize database
	db := app.DB()
	
	// Create tables if they don't exist
	models.InitTables(db)

	// Register routes
	registerRoutes(app)

	// Start the server
	app.Start()
}

func registerRoutes(app *gofr.Gofr) {
	// User routes
	app.GET("/users", handlers.GetUsers)
	app.GET("/users/{id}", handlers.GetUser)
	app.POST("/users", handlers.CreateUser)
	app.PUT("/users/{id}", handlers.UpdateUser)
	app.DELETE("/users/{id}", handlers.DeleteUser)

	// Workout routes
	app.GET("/workouts", handlers.GetWorkouts)
	app.GET("/workouts/{id}", handlers.GetWorkout)
	app.POST("/workouts", handlers.CreateWorkout)
	app.PUT("/workouts/{id}", handlers.UpdateWorkout)
	app.DELETE("/workouts/{id}", handlers.DeleteWorkout)

	// Exercise routes
	app.GET("/exercises", handlers.GetExercises)
	app.GET("/exercises/{id}", handlers.GetExercise)
	app.POST("/exercises", handlers.CreateExercise)
	app.PUT("/exercises/{id}", handlers.UpdateExercise)
	app.DELETE("/exercises/{id}", handlers.DeleteExercise)

	// Workout-Exercise association routes
	app.GET("/workouts/{workoutId}/exercises", handlers.GetWorkoutExercises)
	app.POST("/workouts/{workoutId}/exercises/{exerciseId}", handlers.AddExerciseToWorkout)
	app.DELETE("/workouts/{workoutId}/exercises/{exerciseId}", handlers.RemoveExerciseFromWorkout)

	// User progress routes
	app.GET("/users/{userId}/progress", handlers.GetUserProgress)
	app.POST("/users/{userId}/progress", handlers.RecordUserProgress)
}