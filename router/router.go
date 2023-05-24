package router

import (
	"goserver2/services"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the Gin router and registers the route handlers
func SetupRouter() *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// Register routes
	r.POST("/users", services.RegisterUser)
	r.GET("/users", services.ListUsers)
	r.POST("/movies", services.AddMovie)
	r.DELETE("/movies/:id", services.DeleteMovie)
	r.GET("/movies/:user_id", services.ListMoviesForUser)
	r.GET("/movies", services.ListMovies)

	return r
}
