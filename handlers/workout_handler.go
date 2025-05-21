package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cxocodehub/golang-backend/models"
	"github.com/gofr-dev/gofr"
)

// GetWorkouts handles the GET /workouts request
func GetWorkouts(ctx *gofr.Context) (interface{}, error) {
	// Check if user_id query parameter is provided
	userIDStr := ctx.QueryParam("user_id")
	if userIDStr != "" {
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return nil, gofr.NewError(http.StatusBadRequest, "Invalid user ID")
		}
		
		// Get workouts for specific user
		workouts, err := models.GetUserWorkouts(ctx.DB(), userID)
		if err != nil {
			return nil, gofr.NewError(http.StatusInternalServerError, "Failed to fetch workouts: "+err.Error())
		}
		return workouts, nil
	}
	
	// Get all workouts
	workouts, err := models.GetWorkouts(ctx.DB())
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to fetch workouts: "+err.Error())
	}
	return workouts, nil
}

// GetWorkout handles the GET /workouts/{id} request
func GetWorkout(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid workout ID")
	}

	workout, err := models.GetWorkout(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusNotFound, "Workout not found")
	}
	
	// Get exercises for this workout
	exercises, err := models.GetWorkoutExercises(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to fetch workout exercises: "+err.Error())
	}
	
	// Return workout with exercises
	return map[string]interface{}{
		"workout": workout,
		"exercises": exercises,
	}, nil
}

// CreateWorkout handles the POST /workouts request
func CreateWorkout(ctx *gofr.Context) (interface{}, error) {
	var workout models.Workout
	if err := json.NewDecoder(ctx.Request().Body).Decode(&workout); err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid request body")
	}

	// Validate required fields
	if workout.Name == "" || workout.UserID == 0 {
		return nil, gofr.NewError(http.StatusBadRequest, "Workout name and user ID are required")
	}

	// Create the workout
	id, err := models.CreateWorkout(ctx.DB(), workout)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to create workout: "+err.Error())
	}

	// Return the created workout
	createdWorkout, err := models.GetWorkout(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Workout created but failed to retrieve")
	}

	return createdWorkout, nil
}

// UpdateWorkout handles the PUT /workouts/{id} request
func UpdateWorkout(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid workout ID")
	}

	// Check if workout exists
	existingWorkout, err := models.GetWorkout(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusNotFound, "Workout not found")
	}

	var workout models.Workout
	if err := json.NewDecoder(ctx.Request().Body).Decode(&workout); err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid request body")
	}

	workout.ID = id
	workout.UserID = existingWorkout.UserID // Preserve the original user ID

	// Update the workout
	if err := models.UpdateWorkout(ctx.DB(), workout); err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to update workout: "+err.Error())
	}

	// Return the updated workout
	updatedWorkout, err := models.GetWorkout(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Workout updated but failed to retrieve")
	}

	return updatedWorkout, nil
}

// DeleteWorkout handles the DELETE /workouts/{id} request
func DeleteWorkout(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid workout ID")
	}

	// Check if workout exists
	workout, err := models.GetWorkout(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusNotFound, "Workout not found")
	}

	// Delete the workout
	if err := models.DeleteWorkout(ctx.DB(), id, workout.UserID); err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to delete workout: "+err.Error())
	}

	return map[string]string{"message": "Workout deleted successfully"}, nil
}