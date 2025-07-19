package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Simple GET route
	router.GET("/hello", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Start the server on localhost:8080
	router.Run("localhost:8080")
}
