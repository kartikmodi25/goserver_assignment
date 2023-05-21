package services

import (
	"database/sql"
	"goserver2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

// var db = utils.GetConnection()

// defer db.Close()
// User struct to represent a registered user

// RegisterUser registers a new user
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert user into the database
	query := `
		INSERT INTO users (email, name, dob)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`
	err := db.QueryRow(query, user.Email, user.Name, user.DOB).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// ListUsers retrieves a list of all registered users
func ListUsers(c *gin.Context) {
	rows, err := db.Query("SELECT id, email, name, dob, created_at FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.DOB, &user.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

// AddMovie adds a new movie for a user
func AddMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert movie into the database
	query := `
		INSERT INTO movies (user_id, title)
		VALUES ($1, $2)
		RETURNING id, created_at
	`
	err := db.QueryRow(query, movie.UserID, movie.Title).Scan(&movie.ID, &movie.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

// DeleteMovie deletes a movie by its ID
func DeleteMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	// Delete movie from the database
	result, err := db.Exec("DELETE FROM movies WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}

// ListMoviesForUser retrieves a list of movies watched by a specific user
func ListMoviesForUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	rows, err := db.Query("SELECT id, user_id, title, created_at FROM movies WHERE user_id = $1", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	movies := []models.Movie{}
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.ID, &movie.UserID, &movie.Title, &movie.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		movies = append(movies, movie)
	}

	c.JSON(http.StatusOK, movies)
}

func SetDB(dbconn *sql.DB) {
	db = dbconn
	// fmt.Println(db)
}
