package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cxocodehub/golang-backend/models"
	"github.com/gofr-dev/gofr"
	"golang.org/x/crypto/bcrypt"
)

// GetUsers handles the GET /users request
func GetUsers(ctx *gofr.Context) (interface{}, error) {
	users, err := models.GetUsers(ctx.DB())
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to fetch users: "+err.Error())
	}
	return users, nil
}

// GetUser handles the GET /users/{id} request
func GetUser(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := models.GetUser(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusNotFound, "User not found")
	}
	return user, nil
}

// CreateUser handles the POST /users request
func CreateUser(ctx *gofr.Context) (interface{}, error) {
	var user models.User
	if err := json.NewDecoder(ctx.Request().Body).Decode(&user); err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid request body")
	}

	// Validate required fields
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return nil, gofr.NewError(http.StatusBadRequest, "Username, email, and password are required")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to hash password")
	}
	user.Password = string(hashedPassword)

	// Create the user
	id, err := models.CreateUser(ctx.DB(), user)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to create user: "+err.Error())
	}

	// Return the created user
	createdUser, err := models.GetUser(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "User created but failed to retrieve")
	}

	return createdUser, nil
}

// UpdateUser handles the PUT /users/{id} request
func UpdateUser(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid user ID")
	}

	// Check if user exists
	_, err = models.GetUser(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusNotFound, "User not found")
	}

	var user models.User
	if err := json.NewDecoder(ctx.Request().Body).Decode(&user); err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid request body")
	}

	user.ID = id

	// If password is provided, hash it
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, gofr.NewError(http.StatusInternalServerError, "Failed to hash password")
		}
		user.Password = string(hashedPassword)
	}

	// Update the user
	if err := models.UpdateUser(ctx.DB(), user); err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to update user: "+err.Error())
	}

	// Return the updated user
	updatedUser, err := models.GetUser(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "User updated but failed to retrieve")
	}

	return updatedUser, nil
}

// DeleteUser handles the DELETE /users/{id} request
func DeleteUser(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, gofr.NewError(http.StatusBadRequest, "Invalid user ID")
	}

	// Check if user exists
	_, err = models.GetUser(ctx.DB(), id)
	if err != nil {
		return nil, gofr.NewError(http.StatusNotFound, "User not found")
	}

	// Delete the user
	if err := models.DeleteUser(ctx.DB(), id); err != nil {
		return nil, gofr.NewError(http.StatusInternalServerError, "Failed to delete user: "+err.Error())
	}

	return map[string]string{"message": "User deleted successfully"}, nil
}