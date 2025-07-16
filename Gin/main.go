package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.String(200, "hello from test")
}

func log() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request)
	}
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("auth middleware")
	}
}

func main() {
	fmt.Println("Gin")

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	//global middleware
	router.Use(log())

	router.GET("/test", test)

	authGrp := router.Group("/api")

	authGrp.Use(auth())
	{
		authGrp.GET("/user", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "success"})
		})
		router.Run(":8080")
	}
}
