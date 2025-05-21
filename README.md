# Workout App Backend

A RESTful API for a workout tracking application built with Go and the Gofr framework.

## Features

- User management (registration, authentication)
- Workout plans management
- Exercise library
- Workout-exercise associations
- Progress tracking

## API Endpoints

### Users

- `GET /users` - Get all users
- `GET /users/{id}` - Get a specific user
- `POST /users` - Create a new user
- `PUT /users/{id}` - Update a user
- `DELETE /users/{id}` - Delete a user

### Workouts

- `GET /workouts` - Get all workouts
- `GET /workouts?user_id={userId}` - Get workouts for a specific user
- `GET /workouts/{id}` - Get a specific workout with its exercises
- `POST /workouts` - Create a new workout
- `PUT /workouts/{id}` - Update a workout
- `DELETE /workouts/{id}` - Delete a workout

### Exercises

- `GET /exercises` - Get all exercises
- `GET /exercises/{id}` - Get a specific exercise
- `POST /exercises` - Create a new exercise
- `PUT /exercises/{id}` - Update an exercise
- `DELETE /exercises/{id}` - Delete an exercise

### Workout-Exercise Associations

- `GET /workouts/{workoutId}/exercises` - Get all exercises for a workout
- `POST /workouts/{workoutId}/exercises/{exerciseId}` - Add an exercise to a workout
- `PUT /workouts/{workoutId}/exercises/{exerciseId}` - Update exercise details in a workout
- `DELETE /workouts/{workoutId}/exercises/{exerciseId}` - Remove an exercise from a workout

### Progress Tracking

- `GET /users/{userId}/progress` - Get all progress records for a user
- `GET /users/{userId}/progress?exercise_id={exerciseId}` - Get progress for a specific exercise
- `POST /users/{userId}/progress` - Record new progress
- `DELETE /users/{userId}/progress/{progressId}` - Delete a progress record

## Setup and Installation

1. Clone the repository
2. Set up environment variables (database connection, etc.)
3. Run the application:
   ```
   go run main.go
   ```

## Database Schema

The application uses the following database tables:

- `users` - User information
- `workouts` - Workout plans
- `exercises` - Exercise library
- `workout_exercises` - Association between workouts and exercises
- `progress` - User progress records

## Development

This project uses the Gofr framework, which provides a simple and efficient way to build RESTful APIs in Go.