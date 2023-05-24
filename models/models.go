package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	DOB       time.Time `json:"dob"`
	CreatedAt time.Time `json:"created_at"`
}

// Movie struct to represent a movie
type Movie struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}
