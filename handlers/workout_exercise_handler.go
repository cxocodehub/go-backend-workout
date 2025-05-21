package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cxocodehub/golang-backend/models"
	"github.com/gofr-dev/gofr"
)

// GetWorkoutExercises handles the GET /workouts/{workoutId}/exercises request
func GetWorkoutExercises(ctx *gofr.Context) (interface{}, error) {
	workoutIDStr := ctx.PathParam("workoutId")
	workoutID, err := strconv.Atoi(workoutIDStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid workout ID")
	}

	// Check if workout exists
	_, err = models.GetWorkout(ctx.DB(), workoutID)
	if err != nil {
		return nil, gofr.NewError(http.StatusNotFound, "Workout not found")
	}

	// Get exercises for this workout
	workoutExercises, err := models.GetWorkoutExercises(ctx.DB(), workoutID)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to fetch workout exercises: "+err.Error())
	}

	// If no exercises are found, return an empty array instead of null
	if workoutExercises == nil {
		workoutExercises = []models.WorkoutExercise{}
	}

	return workoutExercises, nil
}

// AddExerciseToWorkout handles the POST /workouts/{workoutId}/exercises/{exerciseId} request
func AddExerciseToWorkout(ctx *gofr.Context) (interface{}, error) {
	workoutIDStr := ctx.PathParam("workoutId")
	workoutID, err := strconv.Atoi(workoutIDStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid workout ID")
	}

	exerciseIDStr := ctx.PathParam("exerciseId")
	exerciseID, err := strconv.Atoi(exerciseIDStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid exercise ID")
	}

	// Check if workout exists
	_, err = models.GetWorkout(ctx.DB(), workoutID)
	if err != nil {
		return nil, gofr.NewError(http.StatusNotFound, "Workout not found")
	}

	// Check if exercise exists
	_, err = models.GetExercise(ctx.DB(), exerciseID)
	if err != nil {
		return nil, gofr.NewError(http.StatusNotFound, "Exercise not found")
	}

	// Parse request body for sets, reps, and weight
	var workoutExercise models.WorkoutExercise
	if err := json.NewDecoder(ctx.Request().Body).Decode(&workoutExercise); err != nil {
		// If no body is provided, use default values
		workoutExercise = models.WorkoutExercise{
			Sets:  3,
			Reps:  10,
			Weight: 0,
		}
	}

	workoutExercise.WorkoutID = workoutID
	workoutExercise.ExerciseID = exerciseID

	// Add exercise to workout
	if err := models.AddExerciseToWorkout(ctx.DB(), workoutExercise); err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to add exercise to workout: "+err.Error())
	}

	return map[string]string{"message": "Exercise added to workout successfully"}, nil
}

// UpdateWorkoutExercise handles the PUT /workouts/{workoutId}/exercises/{exerciseId} request
func UpdateWorkoutExercise(ctx *gofr.Context) (interface{}, error) {
	workoutIDStr := ctx.PathParam("workoutId")
	workoutID, err := strconv.Atoi(workoutIDStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid workout ID")
	}

	exerciseIDStr := ctx.PathParam("exerciseId")
	exerciseID, err := strconv.Atoi(exerciseIDStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid exercise ID")
	}

	var workoutExercise models.WorkoutExercise
	if err := json.NewDecoder(ctx.Request().Body).Decode(&workoutExercise); err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid request body")
	}

	workoutExercise.WorkoutID = workoutID
	workoutExercise.ExerciseID = exerciseID

	// Update workout exercise
	if err := models.UpdateWorkoutExercise(ctx.DB(), workoutExercise); err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to update workout exercise: "+err.Error())
	}

	return map[string]string{"message": "Workout exercise updated successfully"}, nil
}

// RemoveExerciseFromWorkout handles the DELETE /workouts/{workoutId}/exercises/{exerciseId} request
func RemoveExerciseFromWorkout(ctx *gofr.Context) (interface{}, error) {
	workoutIDStr := ctx.PathParam("workoutId")
	workoutID, err := strconv.Atoi(workoutIDStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid workout ID")
	}

	exerciseIDStr := ctx.PathParam("exerciseId")
	exerciseID, err := strconv.Atoi(exerciseIDStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid exercise ID")
	}

	// Remove exercise from workout
	if err := models.RemoveExerciseFromWorkout(ctx.DB(), workoutID, exerciseID); err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to remove exercise from workout: "+err.Error())
	}

	return map[string]string{"message": "Exercise removed from workout successfully"}, nil
}

// ReorderWorkoutExercises handles the PUT /workouts/{workoutId}/exercises/reorder request
func ReorderWorkoutExercises(ctx *gofr.Context) (interface{}, error) {
	workoutIDStr := ctx.PathParam("workoutId")
	workoutID, err := strconv.Atoi(workoutIDStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid workout ID")
	}

	// Parse request body for exercise IDs in new order
	var requestBody struct {
		ExerciseIDs []int `json:"exercise_ids"`
	}
	if err := json.NewDecoder(ctx.Request().Body).Decode(&requestBody); err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid request body")
	}

	if len(requestBody.ExerciseIDs) == 0 {
		return nil, gofr.NewError(http.StatusBadRequest, "Exercise IDs are required")
	}

	// Reorder exercises
	if err := models.ReorderWorkoutExercises(ctx.DB(), workoutID, requestBody.ExerciseIDs); err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to reorder exercises: "+err.Error())
	}

	return map[string]string{"message": "Exercises reordered successfully"}, nil
}