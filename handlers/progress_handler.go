package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cxocodehub/go-backend-workout/models"
	"github.com/gofr-dev/gofr"
)

// GetUserProgress handles the GET /users/{userId}/progress request
func GetUserProgress(ctx *gofr.Context) (interface{}, error) {
	userIDStr := ctx.PathParam("userId")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid user ID")
	}

	// Check if exercise_id query parameter is provided
	exerciseIDStr := ctx.QueryParam("exercise_id")
	if exerciseIDStr != "" {
		exerciseID, err := strconv.Atoi(exerciseIDStr)
		if err != nil {
			return nil, gofr.NewError(http.StatusBadRequest, "Invalid exercise ID")
		}
		
		// Get progress for specific exercise
		progress, err := models.GetExerciseProgress(ctx.DB(), userID, exerciseID)
		if err != nil {
			return nil, gofr.NewError(http.StatusInternalServerError, "Failed to fetch progress: "+err.Error())
		}
		return progress, nil
	}
	
	// Get all progress for user
	progress, err := models.GetUserProgress(ctx.DB(), userID)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to fetch progress: "+err.Error())
	}
	
	// If no progress is found, return an empty array instead of null
	if progress == nil {
		progress = []models.Progress{}
	}
	
	return progress, nil
}

// RecordUserProgress handles the POST /users/{userId}/progress request
func RecordUserProgress(ctx *gofr.Context) (interface{}, error) {
	userIDStr := ctx.PathParam("userId")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid user ID")
	}

	var progress models.Progress
	if err := json.NewDecoder(ctx.Request().Body).Decode(&progress); err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid request body")
	}

	// Set user ID from path parameter
	progress.UserID = userID

	// Validate required fields
	if progress.WorkoutID == 0 || progress.ExerciseID == 0 || progress.Sets == 0 || progress.Reps == 0 {
		return nil, gofr.NewError(http.StatusBadRequest, "Workout ID, exercise ID, sets, and reps are required")
	}

	// If date is not provided, use current date
	if progress.Date.IsZero() {
		progress.Date = time.Now()
	}

	// Record progress
	id, err := models.RecordProgress(ctx.DB(), progress)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to record progress: "+err.Error())
	}

	return map[string]interface{}{
		"id": id,
		"message": "Progress recorded successfully",
	}, nil
}

// DeleteUserProgress handles the DELETE /users/{userId}/progress/{progressId} request
func DeleteUserProgress(ctx *gofr.Context) (interface{}, error) {
	userIDStr := ctx.PathParam("userId")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid user ID")
	}

	progressIDStr := ctx.PathParam("progressId")
	progressID, err := strconv.Atoi(progressIDStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid progress ID")
	}

	// Delete progress
	if err := models.DeleteProgress(ctx.DB(), progressID, userID); err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to delete progress: "+err.Error())
	}

	return map[string]string{"message": "Progress deleted successfully"}, nil
}