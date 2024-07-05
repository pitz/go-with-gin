package models

type WorkoutExecution struct {
	ID        string `json:"id"`
	WorkoutID string `json:"workoutId"`
	Weight    string `json:"weight"`
}
