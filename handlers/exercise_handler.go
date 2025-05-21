package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cxocodehub/go-backend-workout/models"
	"github.com/gofr-dev/gofr"
)

// GetExercises handles the GET /exercises request
func GetExercises(ctx *gofr.Context) (interface{}, error) {
	// Check if category query parameter is provided
	category := ctx.QueryParam("category")
	if category != "" {
		// TODO: Implement filtering by category
		// This would require adding a method to models package
	}
	
	exercises, err := models.GetExercises(ctx.DB())
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to fetch exercises: "+err.Error())
	}
	return exercises, nil
}

// GetExercise handles the GET /exercises/{id} request
func GetExercise(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid exercise ID")
	}

	exercise, err := models.GetExercise(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusNotFound, "Exercise not found")
	}
	return exercise, nil
}

// CreateExercise handles the POST /exercises request
func CreateExercise(ctx *gofr.Context) (interface{}, error) {
	var exercise models.Exercise
	if err := json.NewDecoder(ctx.Request().Body).Decode(&exercise); err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid request body")
	}

	// Validate required fields
	if exercise.Name == "" || exercise.Category == "" {
		return nil, gofr.NewError(http.StatusBadRequest, "Exercise name and category are required")
	}

	// Create the exercise
	id, err := models.CreateExercise(ctx.DB(), exercise)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to create exercise: "+err.Error())
	}

	// Return the created exercise
	createdExercise, err := models.GetExercise(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Exercise created but failed to retrieve")
	}

	return createdExercise, nil
}

// UpdateExercise handles the PUT /exercises/{id} request
func UpdateExercise(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid exercise ID")
	}

	// Check if exercise exists
	_, err = models.GetExercise(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusNotFound, "Exercise not found")
	}

	var exercise models.Exercise
	if err := json.NewDecoder(ctx.Request().Body).Decode(&exercise); err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid request body")
	}

	exercise.ID = id

	// Update the exercise
	if err := models.UpdateExercise(ctx.DB(), exercise); err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to update exercise: "+err.Error())
	}

	// Return the updated exercise
	updatedExercise, err := models.GetExercise(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Exercise updated but failed to retrieve")
	}

	return updatedExercise, nil
}

// DeleteExercise handles the DELETE /exercises/{id} request
func DeleteExercise(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid exercise ID")
	}

	// Check if exercise exists
	_, err = models.GetExercise(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusNotFound, "Exercise not found")
	}

	// Delete the exercise
	if err := models.DeleteExercise(ctx.DB(), id); err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to delete exercise: "+err.Error())
	}

	return map[string]string{"message": "Exercise deleted successfully"}, nil
}